FROM golang:1.16-alpine

WORKDIR /pbly
COPY . .

RUN apk add git
RUN go build -v ./...
RUN go install -v ./...
RUN mkdir -p /etc/pbly/

EXPOSE 8080:8080

CMD ["pbly", "run"]