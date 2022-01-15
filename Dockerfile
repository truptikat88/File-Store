FROM golang:1.12.0-alpine3.9


RUN mkdir /build
ADD . /build
WORKDIR /build

# Build the application
RUN go build -o main .

CMD ["build/main"]

# Command to run
ENTRYPOINT ["/main"]


