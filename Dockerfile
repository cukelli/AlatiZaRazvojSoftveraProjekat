FROM golang:alpine AS build

WORKDIR /app

# Copy the source code from the current directory to the working directory inside the container

COPY go.mod /app
COPY go.sum .
RUN go mod download

# Build the Go app
COPY . .
RUN go build -o main .

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /app/main .
COPY ./swagger.yaml .

EXPOSE 8000

# Expose port 8000 to the outside world

# Command to run the executable
CMD ["./main"]
