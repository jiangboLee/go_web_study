FROM golang

ADD . go/src/github.com/jiangboLee/go_web_study/firstwebapp 
WORKDIR /go/src/github.com/jiangboLee/go_web_study/firstwebapp 
RUN go install github.com/jiangboLee/go_web_study/firstwebapp 

ENTRYPOINT /go/bin/firstwebapp

EXPOSE 8080