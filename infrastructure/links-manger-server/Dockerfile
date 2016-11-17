FROM golang:1.7

RUN go get github.com/viktor-br/links-manager-server/app/links-manager-server
RUN go build github.com/viktor-br/links-manager-server/app/links-manager-server

CMD ["/go/bin/links-manager-server"]

EXPOSE 8080