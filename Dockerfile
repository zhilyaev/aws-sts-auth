ARG GOLANG_VERSION=1.17
ARG ALPINE_VERSION=3.14

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS builder

LABEL maintainer="zhilyaev.dmitriy+aws-sts-auth@gmail.com"
LABEL name="aws-sts-auth"

# enable Go modules support
ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR aws-sts-auth

COPY go.mod go.sum ./
RUN go mod download

# Copy src code from the host and compile it
COPY cmd cmd
COPY pkg pkg
RUN go build -a -o /aws-sts-auth ./main.go

###
FROM alpine:${ALPINE_VERSION} as base-release
RUN apk --no-cache add ca-certificates
ENTRYPOINT ["/bin/aws-sts-auth"]

###
FROM base-release as goreleaser
COPY aws-sts-auth /bin/

###
FROM base-release
COPY --from=builder /aws-sts-auth /bin/
