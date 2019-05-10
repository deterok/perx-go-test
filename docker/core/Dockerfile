FROM golang:1

WORKDIR /go/src/perx-go-test
ADD . .
COPY ./docker-local.json ./local.json

# RUN apt-get -qq update
# RUN apt-get -qq  install --no-install-recommends -y libvips-dev

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

CMD go run .
