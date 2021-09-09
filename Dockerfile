FROM golang:1.17-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /opt/gopath/src/github.com/evanesco/

ADD ./miner-linux.zip /opt/gopath/src/github.com/evanesco/
ADD ./QmNpJg4jDFE4LMNvZUzysZ2Ghvo4UJFcsjguYcx4dTfwKx /opt/gopath/src/github.com/evanesco/
RUN unzip miner-linux.zip && mv ./miner-linux miner && rm miner-linux.zip && mv ./QmNpJg4jDFE4LMNvZUzysZ2Ghvo4UJFcsjguYcx4dTfwKx ./miner

CMD ["/bin/bash"]
