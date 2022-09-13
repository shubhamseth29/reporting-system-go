#build stage here will be the build stage where we just build the binary file
#just add the as keyword as builder

FROM golang:1.19.0-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

#after binary file is produced we will build the second stage
#run stage

FROM alpine:3.16
WORKDIR /app

#copy the exe binary file from builder image to this run stage
COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "/app/main"]