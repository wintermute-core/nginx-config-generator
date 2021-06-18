FROM golang:1.16 as builder
RUN mkdir /build
ADD * /build/
WORKDIR /build
RUN go mod download
RUN go build -a -o nginx-config-generator .

FROM alpine:3.14
COPY --from=builder /build/nginx-config-generator .

ENTRYPOINT [ "./nginx-config-generator" ]
