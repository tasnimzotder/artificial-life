FROM golang:1.21-alpine
LABEL authors="Tasnim Zotder <hello@tasnimzotder.com>"
LABEL version="0.0.1"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

#RUN go build -o main .
#
#CMD ["./main"]

CMD ["go", "run", "main.go"]