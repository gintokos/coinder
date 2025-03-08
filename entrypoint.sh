#!/bin/bash
# Не завершать скрипт при ошибках
set +e

# Environment variables
# Set these when running the container
# DOMAINS="example.com,www.example.com"
# EMAIL="admin@example.com"

# Check for required environment variables
if [ -z "$DOMAINS" ]; then
  echo "Error: DOMAINS environment variable is not set."
  echo "Example: DOMAINS=example.com,www.example.com"
  exit 1
fi

if [ -z "$EMAIL" ]; then
  echo "Error: EMAIL environment variable is not set."
  echo "Example: EMAIL=admin@example.com"
  exit 1
fi

# Function to update certificates
update_certs() {
  # Split domains by comma
  IFS=',' read -ra DOMAIN_LIST <<< "$DOMAINS"
  DOMAIN_ARGS=""
  
  # Form arguments for certbot
  for domain in "${DOMAIN_LIST[@]}"; do
    DOMAIN_ARGS="$DOMAIN_ARGS -d $domain"
  done
  
  # Start nginx to handle acme-challenge requests
  echo "Starting Nginx for ACME challenge handling..."
  nginx
  
  # First check if the webroot path is accessible
  echo "Testing ACME challenge path..."
  mkdir -p /var/www/certbot/.well-known/acme-challenge/
  echo "test" > /var/www/certbot/.well-known/acme-challenge/test
  curl -s http://localhost/.well-known/acme-challenge/test
  
  # Check if certificate already exists for the first domain
  FIRST_DOMAIN=${DOMAIN_LIST[0]}
  
  # Пропустим стадию тестового сертификата и сразу получим боевой
  echo "Getting production certificate for $DOMAINS"
  
  # Удалим существующие сертификаты, если они есть
  if [ -d "/etc/letsencrypt/live/$FIRST_DOMAIN" ]; then
    echo "Removing existing certificates for clean installation"
    rm -rf /etc/letsencrypt/live/$FIRST_DOMAIN
    rm -rf /etc/letsencrypt/archive/$FIRST_DOMAIN
    rm -f /etc/letsencrypt/renewal/$FIRST_DOMAIN.conf
  fi
  
  # Получаем сразу боевой сертификат
  certbot --nginx \
    $DOMAIN_ARGS \
    --email $EMAIL \
    --agree-tos \
    --no-eff-email \
    --non-interactive
  
  # Если не удалось получить боевой сертификат, генерируем самоподписанный как резервный вариант
  if [ $? -ne 0 ]; then
    echo "Production certificate failed, generating self-signed certificate as fallback"
    mkdir -p /etc/letsencrypt/live/$FIRST_DOMAIN
    openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
      -keyout /etc/letsencrypt/live/$FIRST_DOMAIN/privkey.pem \
      -out /etc/letsencrypt/live/$FIRST_DOMAIN/fullchain.pem \
      -subj "/CN=$FIRST_DOMAIN"
    # Create empty chain.pem
    touch /etc/letsencrypt/live/$FIRST_DOMAIN/chain.pem
    
    # Manually configure nginx for SSL with self-signed certificate
    if ! grep -q "ssl_certificate" /etc/nginx/nginx.conf; then
      # If nginx.conf doesn't exist, try site-specific configs
      if [ -f "/etc/nginx/conf.d/default.conf" ]; then
        # Try to modify site config if it exists
        sed -i "/listen 80;/a \    listen 443 ssl;\n    ssl_certificate /etc/letsencrypt/live/$FIRST_DOMAIN/fullchain.pem;\n    ssl_certificate_key /etc/letsencrypt/live/$FIRST_DOMAIN/privkey.pem;\n    ssl_trusted_certificate /etc/letsencrypt/live/$FIRST_DOMAIN/chain.pem;\n    include /etc/letsencrypt/options-ssl-nginx.conf;\n    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;" /etc/nginx/conf.d/default.conf
      else
        # Try to modify main config
        sed -i "/server {/a \    listen 443 ssl;\n    ssl_certificate /etc/letsencrypt/live/$FIRST_DOMAIN/fullchain.pem;\n    ssl_certificate_key /etc/letsencrypt/live/$FIRST_DOMAIN/privkey.pem;\n    ssl_trusted_certificate /etc/letsencrypt/live/$FIRST_DOMAIN/chain.pem;\n    include /etc/letsencrypt/options-ssl-nginx.conf;\n    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;" /etc/nginx/nginx.conf
      fi
    fi
  fi
  
  # Restart Nginx with the new configuration
  nginx -s stop
}

# Set up automatic certificate renewal
setup_cron() {
  echo "Setting up automatic certificate renewal"
  echo "0 12 * * * certbot renew --quiet --non-interactive && nginx -s reload" > /etc/crontabs/root
  crond
}

# Main process
main() {
  # Update certificates
  update_certs
  
  # Set up automatic renewal
  setup_cron
  
  # Check Nginx configuration
  echo "Checking Nginx configuration..."
  nginx -t
  
  echo "Starting Nginx..."
  # Run Nginx in foreground
  nginx -g "daemon off;"
}

# Run main process
main