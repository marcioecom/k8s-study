FROM golang:1.18 as builder
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go build -o server

FROM scratch as runner
COPY --from=builder /go/src/app/server /server
ENTRYPOINT [ "./server" ]
