FROM golang:1.15.2-buster

COPY ./vendor /app/vendor
COPY ./go.mod ./go.sum /app/
COPY ./main.go /app
WORKDIR /app
RUN GOOS=linux go build -a -mod vendor -o /bin/app

EXPOSE 8080
ENTRYPOINT ["/bin/app"]
