from golang:1.16.3-alpine3.13 

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY . ./

# remove the test files
RUN rm -rf ./test
RUN rm -rf ./config/MockDBSetup.go
RUN go mod tidy

# Download the Go module dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN go build -o myapp

# Set environment variables
ENV GO_ENV=${GO_ENV}
ENV DB_LOCAL_HOST=${DB_LOCAL_HOST}
ENV DB_LOCAL_PORT=${DB_LOCAL_PORT}
ENV DB_LOCAL_USER=${ENV_VAR2}
ENV DB_LOCAL_PASSWORD=${DB_LOCAL_PASSWORD}
ENV DB_LOCAL_NAME=${DB_LOCAL_NAME}
ENV JWT_SECRET=${JWT_SECRET}
ENV TOKEN_HOUR_LIFESPAN=${TOKEN_HOUR_LIFESPAN}
# Expose port 8080
EXPOSE 8080

# Run the Go application
CMD ["./myapp"]


# run mkdir -p /go/src/myFirstAPI

# ENV GO111MODULE=on
# ENV GOFLAGS=-mod=vendor

# workdir /go/src/myFirstAPI


# copy . /go/src/myFirstAPI

# run go mod download
# run go mod vendor

# run go install myFirstAPI

# expose 8080