FROM golang:1.9.4-alpine3.7 as build

WORKDIR /go/src/github.com/puppetlabs/kreamlet/bootstrap
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/bootstrap github.com/puppetlabs/kreamlet/bootstrap

FROM scratch

WORKDIR /usr/bin
COPY --from=build /go/src/github.com/puppetlabs/kreamlet/bootstrap/bin /usr/bin
CMD ["bootstrap"]

