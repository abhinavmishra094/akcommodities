FROM golang:1.20

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /exchange

EXPOSE 5000

CMD [ "/exchange" ]