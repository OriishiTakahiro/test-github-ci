FROM golang:1.18 as app1
WORKDIR /app
COPY ./app1 ./
RUN go build

CMD ["./app1"]

FROM golang:1.18 as app2
WORKDIR /app
COPY ./app2 ./
RUN go build

CMD ["./app2"]

FROM golang:1.18 as app3
WORKDIR /app
COPY ./app3 ./
RUN go build
CMD ["./app3"]
