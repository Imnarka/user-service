FROM golang:1.23

RUN apt-get update && apt-get install -y make git iputils-ping net-tools postgresql-client

WORKDIR /app

RUN go version

COPY go.mod go.sum ./

COPY .env .

COPY Makefile ./

RUN echo "Running make deps" && make deps

COPY . .

RUN make build

EXPOSE 50051

ENTRYPOINT ["./entrypoint.sh"]

CMD ["./bin/main"]