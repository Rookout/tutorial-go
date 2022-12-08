FROM golang:1.18-alpine3.15 as builder

RUN apk --update --no-cache add git gcc musl-dev protobuf-dev openssl-libs-static openssl-dev build-base zlib-static

WORKDIR /app
ADD . .

RUN go mod download
RUN go mod tidy

RUN go build -tags=alpine314,rookout_static -gcflags='all=-N -l' cmd/main.go

FROM alpine:3.15 as release
COPY --from=builder /app/main ./
COPY --from=builder /app/web /web

ENV PORT 1994
EXPOSE 1994
CMD ["./main"]
