FROM golang

ADD . go/src/firstwebapp 
WORKDIR /go/src/firstwebapp
RUN go install firstwebapp

ENTRYPOINT /go/bin/firstwebapp

EXPOSE 8080