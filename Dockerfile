# Use OpenResty image, which includes Nginx with Lua support
FROM openresty/openresty:alpine

# Copy your custom Nginx configuration into the container
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf

# Copy your script for updating InfluxDB into the container
COPY ./influxdb/update_influxdb.sh /usr/local/bin/update_influxdb.sh

# Make sure the script is executable
RUN chmod +x /usr/local/bin/update_influxdb.sh

# Expose port 80 for HTTP traffic
EXPOSE 80

# Start Nginx when the container starts
CMD ["nginx", "-g", "daemon off;"]
