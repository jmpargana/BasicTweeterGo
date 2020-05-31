FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go install

EXPOSE 8080

CMD ["tweeter"]
