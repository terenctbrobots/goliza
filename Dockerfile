FROM golang:1.16-alpine
WORKDIR /app
COPY eliza ./eliza
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY /*.go ./
RUN go build -o /server
EXPOSE 8000
CMD ["/server"]