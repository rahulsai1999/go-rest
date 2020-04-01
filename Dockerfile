FROM golang:alpine as builder
WORKDIR /go/src/app
COPY . .
RUN go get -u
RUN go build ./main.go

FROM alpine
COPY --from=builder /go/src/app/main .
COPY --from=builder /go/src/app/.env .
CMD ["./main"]