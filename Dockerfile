FROM golang:1.14-buster

# Add Maintainer info
LABEL maintainer="Aleksandr Kasatkin <alexkasatkin90@gmail.com>"


RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 
RUN go build -o gorm_rest_api main.go
# Copy the source from the current directory to the working Directory inside the container 
CMD ["./gorm_rest_api"]