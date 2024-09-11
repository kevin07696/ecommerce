# Stage 1: Build and test
FROM alpine:latest AS build-stage

# Install Go using apk
RUN apk add --no-cache go

# Install npm using apk
RUN apk add --no-cache npm

# Define labels to provide metadata about the image
LABEL maintainer="kevinlam07696+project@gmail.com"
LABEL version="1.0"
LABEL description="Default Dockerfile for golang projects with node modules"

# Set the working directory in the container
WORKDIR /app

# Copy all files to into the container
COPY . .

# Install the Go dependencies
RUN go mod download

# Install the npm dependencies
RUN npm install

# Build the user package
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint /app/cmd/main.go

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian12 AS build-release-stage

# Set the working directory inside the container
WORKDIR /

COPY --from=build-stage /entrypoint /entrypoint
COPY --from=build-stage /app/static /static
COPY --from=build-stage /app/.env /.env

# Expose port 8080 to the outside world
EXPOSE ${PORT}
USER nonroot:nonroot

# Command to run the executable
ENTRYPOINT ["/entrypoint"]
