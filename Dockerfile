FROM golang:1.16-alpine as build
WORKDIR /app
COPY eliza ./eliza
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY /*.go ./
RUN go build -o /server


FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /server /server
EXPOSE 8000
USER nonroot:nonroot
CMD ["/server"]