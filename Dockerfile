FROM golang:1.21 as build

RUN curl -sL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs
RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN wails dev

# Stage 2: Create a lightweight image to run the application
FROM debian:bullseye-slim

# Install any necessary dependencies (adjust as needed)
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Copy the compiled Wails application from the build stage
COPY --from=build /app/build/bin /app

# Set the working directory
WORKDIR /app

# Set the entrypoint to the compiled binary
ENTRYPOINT ["./myproject"]