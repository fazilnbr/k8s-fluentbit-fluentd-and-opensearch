FROM golang:1.22.3 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY *.go ./

# RUN CGO_ENABLED=0 GOOS=linux go build -o fluentd_testing . no goarc defined please update your go arc if needed as per the requirment 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fluentd_testing .

EXPOSE 8080

CMD ["./fluentd_testing"]