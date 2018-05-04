FROM golang:1.10.0 AS builder
WORKDIR /go/src/github.com/tophatsteve/saas/
COPY ./service/ ./service/
COPY ./vendor/ ./vendor/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./service/cmd/saas ./service/cmd/main.go

FROM scratch
COPY --from=builder /go/src/github.com/tophatsteve/saas/service/cmd/saas .
EXPOSE 9191
ENTRYPOINT  ["/saas"]
