FROM docker.io/golang:latest as build

COPY . /
WORKDIR /
RUN CGO_ENABLED=0 GOOS=linux go build -a -o receiver

FROM scratch
COPY --from=build /receiver .

CMD ["/receiver"]