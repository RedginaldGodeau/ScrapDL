FROM golang:1.20

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/build/scrapDL ./internal/cmd
RUN CGO_ENABLED=0 GOOS=windows go build -o bin/build/scrapDL.exe ./internal/cmd
