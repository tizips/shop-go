# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image

FROM golang:1.23 as builder

# Add Maintainer Info
LABEL maintainer="tizips <tizips@163.com>"

# Set the Current Working Directory inside the container
WORKDIR /build

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

WORKDIR /build/web

# Build the Go app
RUN CGO_ENABLED=0 go build -o application .

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy the Pre-build binary file from the previous stage
COPY --from=builder /build/web/application .
COPY --from=builder /build/web/conf ./conf

# Expose port 8080 to the outside world
EXPOSE 9600

# Command to run the executable
CMD ["/app/application", "server"]