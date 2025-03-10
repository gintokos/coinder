#!/bin/bash
set +e

# Environment variables
# Set these when running the container
# DOMAINS="example.com,www.example.com"
# EMAIL="admin@example.com"

# Флаг для пропуска обновления сертификата (установите 1, чтобы пропустить)
SKIP_CERT_RENEWAL=0

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
  
  # First domain is the primary one
  FIRST_DOMAIN=${DOMAIN_LIST[0]}
  
  # Create required directories
  mkdir -p /var/www/certbot/.well-known/acme-challenge/
  
  # Проверяем, не запущен ли уже Nginx
  if ! pgrep -x "nginx" > /dev/null; then
    echo "Starting Nginx temporarily for ACME challenge..."
    nginx
    
    # Даем Nginx время на запуск
    sleep 2
  else
    echo "Nginx already running, using it for ACME challenge..."
  fi
  
  # Запускаем certbot только если не указано пропустить обновление
  if [ $SKIP_CERT_RENEWAL -ne 1 ]; then
    echo "Requesting/renewing certificate for $FIRST_DOMAIN (if needed)"
    
    # certbot сам решит, нужно ли обновлять сертификат
    certbot --nginx \
      $DOMAIN_ARGS \
      --email $EMAIL \
      --agree-tos \
      --no-eff-email \
      --non-interactive \
      --keep-until-expiring \
      --expand
    
    CERTBOT_EXIT_CODE=$?
    if [ $CERTBOT_EXIT_CODE -ne 0 ]; then
      echo "Error requesting certificate (exit code: $CERTBOT_EXIT_CODE)"
      echo "Continuing with existing configuration if available"
    else
      echo "Certificate successfully obtained/renewed and Nginx configured"
    fi
  else
    echo "Certificate renewal skipped by user configuration"
  fi
  
  # Перезагружаем конфигурацию Nginx
  echo "Reloading Nginx configuration..."
  nginx -s reload
}

# Set up automatic certificate renewal
setup_cron() {
  echo "Setting up automatic certificate renewal"
  echo "0 12 * * * certbot renew --quiet --non-interactive && nginx -s reload" > /etc/crontabs/root
  crond
}

# Main process
main() {
  # Create required directories
  mkdir -p /etc/letsencrypt/live
  mkdir -p /var/www/certbot
  
  # Проверяем, нет ли уже запущенных экземпляров Nginx
  if pgrep -x "nginx" > /dev/null; then
    echo "Found running Nginx instances. Stopping them before starting..."
    nginx -s stop
    sleep 3  # Даем Nginx время на полную остановку
    
    # Дополнительная проверка, если nginx все еще запущен
    if pgrep -x "nginx" > /dev/null; then
      echo "Warning: Nginx is still running. Trying to kill processes..."
      killall -9 nginx
      sleep 1
    fi
  fi
  
  # Запускаем Nginx в фоне
  echo "Starting Nginx..."
  nginx
  
  # Update certificates
  update_certs
  
  # Set up automatic renewal
  setup_cron
  
  # Останавливаем текущий Nginx и запускаем в основном режиме
  echo "Restarting Nginx in foreground mode..."
  nginx -s stop
  sleep 2
  
  # Run Nginx in foreground
  echo "Starting Nginx in foreground mode..."
  exec nginx -g "daemon off;"
}

# Run main process
main