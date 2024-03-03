# Start from the latest Golang base image
FROM golang:latest

# Install Air for live reloading
RUN go install github.com/cosmtrek/air@latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all dependencies
RUN go mod download

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable using Air
CMD ["air"]