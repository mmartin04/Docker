# Start from a Go image.
FROM golang:1.18

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
