FROM golang:latest

# Set the current working directory inside the container
WORKDIR /mediateq

# Copy the go.mod and go.sum files to the WORKDIR
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Install libvips-dev
RUN apt-get update && \
    apt-get install -y libvips-dev


# Copy the source code to the WORKDIR
COPY . .

# Build mediateq
RUN chmod +x build.sh
RUN ./build.sh 

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./bin/mediateq"]
