# Build stage
FROM node:18-alpine as builder

# Set working directory
WORKDIR /app

# Copy only package files for dependency installation
COPY package*.json ./

# Use existing node_modules if they exist
COPY node_modules ./node_modules

# Install only missing dependencies
RUN npm ci --prefer-offline --no-audit

# Copy source code
COPY . .

# Build the application
RUN npm run build

# Production stage
FROM nginx:alpine

# Copy built files to nginx server directory    
COPY --from=builder /app/dist /usr/share/nginx/html

# Copy custom nginx configuration if needed
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Start nginx
CMD ["nginx", "-g", "daemon off;"]