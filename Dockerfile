FROM golang:alpine

#WORKDIR $GOPATH/src/one/go-misc
#COPY . $GOPATH/src/one/go-misc
#RUN go build .
ADD ./go-misc /app/
# cd到app文件夹下
WORKDIR /app

ENTRYPOINT ["./go-misc"]