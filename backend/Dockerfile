FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go mod download

RUN apt update
RUN apt install ffmpeg -y

COPY . .

EXPOSE 3000

CMD ["air"]
