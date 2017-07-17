FROM golang:1.8-alpine
RUN apk update && apk upgrade && apk add --no-cache bash git openssh
RUN mkdir -p /go/src/github.com/gost/server/
ADD . /go/src/github.com/gost/server
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/yaml.v2
RUN go get github.com/lib/pq
RUN go get github.com/gost/now
RUN go get github.com/gost/godata
RUN go get github.com/eclipse/paho.mqtt.golang
RUN go build -o /go/bin/gost/gost github.com/gost/server
RUN cp /go/src/github.com/gost/server/config.yaml /go/bin/gost/config.yaml
WORKDIR /go/bin/gost
ENTRYPOINT ["/go/bin/gost/gost"]
EXPOSE 8080
