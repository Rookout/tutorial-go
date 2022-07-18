FROM golang:1.18-alpine3.15 as builder

ARG ARTIFACTORY_CREDS

RUN go env -w GONOSUMDB="github.com/Rookout/GoSDK"
RUN go env -w GOPROXY="https://proxy.golang.org,https://${ARTIFACTORY_CREDS}@rookout.jfrog.io/artifactory/api/go/rookout-go,direct"

RUN apk --update --no-cache add git gcc musl-dev protobuf-dev openssl-libs-static openssl-dev build-base zlib-static

WORKDIR /app
ADD . .

# Because of the GoRook deployment architecture, the go.mod is pointing to the stub version and the CI is responsible for replacing it with the real version
# We download the GoRook explicitly so it would register it in the go.sum
RUN go get github.com/Rookout/GoSDK@v0.1.15

RUN go mod download
RUN go mod tidy

RUN go build -tags=alpine314,rookout_static -gcflags='all=-N -l' cmd/main.go

FROM alpine:3.15 as release
COPY --from=builder /app/main ./
COPY --from=builder /app/web /web

ENV PORT 1994
EXPOSE 1994
CMD ["./main"]
