FROM golang:latest

RUN mkdir -p /go/src/rpayApi

WORKDIR /go/src/rpayApi

COPY . /go/src/rpayApi
COPY ./cmd/main/* /go/src/rpayApi/
COPY main.go /main.go
COPY entrypoint.sh /entrypoint.sh

#RUN go run main.go

CMD go run main.go
#ENTRYPOINT [ "sh", "/entrypoint.sh" ]


EXPOSE 8080
