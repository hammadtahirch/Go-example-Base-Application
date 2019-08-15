# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Rajeev Singh <rajeevhub@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

RUN apt-get update
RUN apt-get install vim -y

# Install gin
RUN go get github.com/codegangsta/gin
# install golang driver for mysql
RUN go get "github.com/go-sql-driver/mysql"

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
# RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 3001

# Command to run the executable
# CMD ["./main"]
CMD ["gin","-i","-a","3002","run","main.go"]
