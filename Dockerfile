FROM golang:1.18-alpine3.15

ARG ARTIFACTORY_CREDS

RUN go env -w GONOSUMDB="github.com/Rookout/GoSDK"
RUN go env -w GOPROXY="https://proxy.golang.org,https://${ARTIFACTORY_CREDS}@rookout.jfrog.io/artifactory/api/go/rookout-go,direct"

WORKDIR /app
ADD . .
RUN go mod download
RUN go mod tidy
RUN go build cmd/main.go

ENV PORT 1994
EXPOSE 1994
CMD ["./main"]
