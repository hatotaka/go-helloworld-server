FROM busybox

EXPOSE 8080

ADD ./go-helloworld-server /go-helloworld-server

ENTRYPOINT /go-helloworld-server
