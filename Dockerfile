FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-asciiweb
# Expose port 8080
EXPOSE 8080

# Run
CMD ["/docker-asciiweb"]