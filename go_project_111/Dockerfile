FROM golang:latest
WORKDIR /main_app
COPY . . 
RUN go env -w GO111MODULE=off
RUN go build -o main
CMD ["./main"]
