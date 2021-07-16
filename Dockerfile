FROM golang:latest as builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/brewery-api
COPY go.* .
RUN go mod download
COPY . .
RUN go build -ldflags="-w -s" -o /go/bin/brewery-api


FROM scratch
EXPOSE 8080
WORKDIR /app/
COPY --from=builder /go/bin/brewery-api .
ENTRYPOINT ["./brewery-api" ]