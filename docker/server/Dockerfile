FROM golang:1.11.1

WORKDIR /
COPY src/server.go .
COPY main.html .
COPY head.html .
COPY body.html .
COPY makefile .

EXPOSE 8080

RUN  go get -d github.com/tarantool/go-tarantool
RUN ["make", "server"]

CMD [ "./server" ]