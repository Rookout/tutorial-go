FROM golang:1.18-alpine3.15

ARG ARTIFACTORY_CREDS

RUN go env -w GONOSUMDB="github.com/Rookout/GoSDK"
RUN go env -w GOPROXY="https://proxy.golang.org,https://${ARTIFACTORY_CREDS}@rookout.jfrog.io/artifactory/api/go/rookout-go,direct"

WORKDIR /app
ADD . .

# Because of the GoRook deployment architecture, the go.mod is pointing to the stub version and the CI is responsible for replacing it with the real version
# We download the GoRook explicitly so it would register it in the go.sum
RUN go get -d github.com/Rookout/GoSDK@v0.1.8

RUN go mod download
RUN go mod tidy
RUN go build cmd/main.go

ENV PORT 1994
EXPOSE 1994
CMD ["./main"]
