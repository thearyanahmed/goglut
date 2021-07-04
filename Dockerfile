FROM golang:latest

RUN mkdir -p /go/src/goglut

WORKDIR /go/src/goglut

#COPY . /go/src/goglut

ENV PORT 8080

ENTRYPOINT ["tail", "-f", "/dev/null"]
