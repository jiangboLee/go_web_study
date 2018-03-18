FROM golang

ADD . $GOPATH/src/github.com/jiangboLee/go_web_study/leechitchat 
WORKDIR $GOPATH/src/github.com/jiangboLee/go_web_study/leechitchat 

RUN go get github.com/go-sql-driver/mysql

RUN go install github.com/jiangboLee/go_web_study/leechitchat 

ENTRYPOINT /go/bin/leechitchat

EXPOSE 8080

# http://blog.csdn.net/razilfelix/article/details/70247766
# http://blog.csdn.net/boonya/article/details/74906927