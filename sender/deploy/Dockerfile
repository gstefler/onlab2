FROM docker.io/golang:latest as build

COPY . /
WORKDIR /
RUN CGO_ENABLED=0 GOOS=linux go build -a -o sender

FROM scratch
COPY --from=build /sender .

CMD ["/sender"]