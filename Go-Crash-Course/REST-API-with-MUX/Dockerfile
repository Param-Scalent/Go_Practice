FROM golang:latest

LABEL maintainer="Param <param@scalent.io>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT=8000

RUN go build

CMD [ "./REST-API-with-MUX" ]
