# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Hammad Tahir <hammad.tahir.ch@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

RUN apt-get update
RUN apt-get install nano -y

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy everything from the current directory to the Working Directory inside the container
COPY ./ /app

# install golang driver for mysql
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/Altoros/gorm-goose/cmd/gorm-goose"
RUN go get "github.com/githubnemo/CompileDaemon"

# Expose port 3001 to the outside world
EXPOSE 3001


ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
