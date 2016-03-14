FROM ubuntu

EXPOSE 8080

ADD go-helloworld-server /bin/go-helloworld-server

ENTRYPOINT /bin/go-helloworld-server
