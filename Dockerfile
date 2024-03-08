# syntax=docker/dockerfile:experimental
FROM golang:1.19.2-alpine3.16 AS buildImg
RUN apk add git openssh

ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8

WORKDIR /go/src/chaincode
ENV GO111MODULE=on

COPY src/go.mod src/go.sum ./
RUN go mod download

RUN rm -rf ~/.ssh

COPY src .

# create test package
RUN CGO_ENABLED=0 go test -c . -o /out/chaincode.test -cover

# Build application
RUN go build -o /out/chaincode

EXPOSE 9999
CMD ./chaincode

FROM alpine:3.11

WORKDIR /go/src/chaincode

COPY --from=buildImg /out/ .

EXPOSE 9999

CMD ./chaincode
