# > docker build -t cosmos-cash .
#
# Server:
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.poaapp:/root/.poaapp poaapp init test-chain
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.poaapp:/root/.poaapp poaapp start
#
# Client: (Note the poaapp binary always looks at ~/.poaapp we can bind to different local storage)
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.poaappcli:/root/.poaapp poaapp keys add foo
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.poaappcli:/root/.poaapp poaapp keys list
FROM golang:alpine AS build-env

RUN apk update
RUN apk add --no-cache curl make git libc-dev bash gcc linux-headers eudev-dev python3

WORKDIR /go/src/github.com/PaddyMc/authority

COPY . .

RUN make install

## Final image
#FROM gcr.io/distroless/base:latest
FROM alpine:edge

## Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

COPY --from=build-env /go/bin/poad /usr/bin/poad
COPY --from=build-env /go/bin/poacli /usr/bin/poacli

EXPOSE 26656 26657 1317 9090

CMD ["poad"]

