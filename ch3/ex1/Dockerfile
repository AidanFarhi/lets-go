# syntax=docker/dockerfile:1
FROM golang:1.22.0 as build
WORKDIR /src
COPY main.go .
RUN go build -o /bin/main ./main.go

FROM scratch
COPY --from=build /bin/main /bin/main
CMD [ "/bin/main" ]