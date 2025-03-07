FROM nginx:alpine

# Install required packages
RUN apk add --no-cache certbot certbot-nginx openssl bash curl

# Remove default configuration
RUN rm /etc/nginx/conf.d/default.conf

# Copy our configuration
COPY default.conf /etc/nginx/conf.d/

# Create directories for certbot
RUN mkdir -p /var/www/certbot
RUN mkdir -p /etc/letsencrypt

# Copy and make executable the entrypoint script
COPY entrypoint.sh /
RUN chmod +x /entrypoint.sh


# Run entrypoint
ENTRYPOINT ["/entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]