FROM golang:1.23-alpine AS builder

ARG TARGETOS
ARG TARGETARCH

RUN apk add --no-cache --update

COPY . /json-convert

RUN go install  github.com/parsyl/parquet/cmd/parquetgen@latest

WORKDIR /json-convert

RUN go mod download \
    && go mod verify \
    && go generate \
    && CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -a -installsuffix cgo -ldflags="-w -s" -o convert

FROM scratch

COPY --from=builder /json-convert/convert /convert

ENTRYPOINT ["/convert"]