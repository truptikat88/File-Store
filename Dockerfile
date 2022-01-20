# builder image
FROM golang:1.17.6-alpine3.15 as builder
RUN mkdir /File-Store
COPY . /File-Store/
WORKDIR /File-Store
RUN ls .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o file-service .


# generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /File-Store/file-service .
COPY FileStore /FileStore 

# executable
ENTRYPOINT [ "./file-service" ]


