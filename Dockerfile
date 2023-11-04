
# Start from the latest golang base image
FROM golang:1.19-alpine3.16 as builder

# Add Maintainer Info
LABEL maintainer="Albin Antony"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux go build -a -o go-api .
RUN --mount=type=cache,target=/root/.cache/go-build  go build -o sports-day .

######## Start a new stage from scratch #######
FROM alpine:3.16

RUN apk --no-cache add ca-certificates

RUN apk add --no-cache tzdata

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/sports-day .


# Expose port 3333 to the outside world
EXPOSE 3333

# Build Args
ARG LOG_DIR=/app/logs

ARG STAGE

# Create Log Directory
RUN mkdir -p ${LOG_DIR}

# Environment Variables

ENV LOG_DIR=${LOG_DIR}/app.log

# Declare volumes to mount
# VOLUME [${LOG_DIR}]

ENV STAGE=${STAGE}


# Command to run the executable
CMD ["./sports-day"]