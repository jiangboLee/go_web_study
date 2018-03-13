FROM golang

ADD . $GOPATH/src/github.com/jiangboLee/go_web_study/firstwebapp 
WORKDIR $GOPATH/src/github.com/jiangboLee/go_web_study/firstwebapp 
RUN go install github.com/jiangboLee/go_web_study/firstwebapp 

ENTRYPOINT /go/bin/firstwebapp

EXPOSE 8080

# http://blog.csdn.net/razilfelix/article/details/70247766
# http://blog.csdn.net/boonya/article/details/74906927