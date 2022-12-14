#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux
RUN go build -ldflags="-s -w" -o /go/bin/app -v .

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=algorithms-in-go Version=0.1.1
# EXPOSE 3000
