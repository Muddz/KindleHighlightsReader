FROM golang:alpine
ENV GO111MODULE=on
WORKDIR /betterman
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
EXPOSE 8080
RUN go build -o betterman-server ./main
CMD ["./betterman-server"]
