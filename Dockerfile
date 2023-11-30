FROM golang:1.17-alpine
LABEL authors="Tasnim Zotder <hello@tasnimzotder.com>"
LABEL version="0.0.1"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

#RUN go build -o main .
#
#CMD ["./main"]

RUN go install "fyne.io/fyne/v2/cmd/fyne@latest"
RUN go install "github.com/gopherjs/gopherjs@latest"

EXPOSE 8080

CMD ["fyne", "serve"]

#CMD ["go", "run", "main.go"]