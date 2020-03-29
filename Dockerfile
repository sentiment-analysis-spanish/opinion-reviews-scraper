
# build it:
# docker build -t opinion-reviews-scraper .
# run it:
# docker run --restart always --network my-network -d opinion-reviews-scraper
# docker run -d opinion-reviews-scraper
# docker run --network my-network -d opinion-reviews-scraper
# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang v1.12 base image
FROM golang:1.12

# Add Maintainer Info
LABEL maintainer="Hugo J. Bello <hjbello.wk@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR $home/test

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]