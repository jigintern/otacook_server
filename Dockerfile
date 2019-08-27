FROM golang:1.12.9
CMD /bin/bash
ENV GO111MODULE=on
WORKDIR /go/src/github.com/jigintern/otacook_server
RUN go get -u github.com/gin-gonic/gin \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql \
  && go get github.com/joho/godotenv
