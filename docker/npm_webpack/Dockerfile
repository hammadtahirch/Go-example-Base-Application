# Start from node:8.4 base image
FROM node:8.4

# Add Maintainer Info
LABEL maintainer="Hammad Tahir <hammad.tahir.ch@gmail.com>"

# gloable install npm webpack 
RUN npm install -g webpack webpack-dev-server yarn

# directory
VOLUME /app

# Expose port 2000 to the outside world
EXPOSE 2000

#working directory
WORKDIR /app
