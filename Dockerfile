FROM golang:1.14

WORKDIR /app

COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go build -o main .


EXPOSE 8080
CMD ["/app/main"]


