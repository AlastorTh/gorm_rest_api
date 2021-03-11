FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="Aleksandr Kasatkin <alexkasatkin90@gmail.com>"


RUN go version
ENV GOPATH=/

COPY ./ ./


# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 
RUN go build -o gorm_rest_api main.go
# Copy the source from the current directory to the working Directory inside the container 
CMD ["./gorm_rest_api"]