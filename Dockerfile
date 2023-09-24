FROM golang:1.20

# Set destination for COPY
WORKDIR /app

# Copy the source code
COPY . .

# Build
RUN go build -v -o /docker-asciiweb

# Run
CMD ["/docker-asciiweb"]