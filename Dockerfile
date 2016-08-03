FROM golang:1.7-alpine
COPY rest-api.go /usr/local/go/src/go-rest-api/
RUN apk add --update git
RUN mkdir -p /usr/local/go/src/vendor/github.com/gorilla && cd /usr/local/go/src/vendor/github.com/gorilla && git clone https://github.com/gorilla/mux.git
RUN go install go-rest-api
EXPOSE 8080
CMD ["go-rest-api"]
