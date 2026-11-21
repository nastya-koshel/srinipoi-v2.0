FROM golang:1.17-alpine AS build
WORKDIR /src/
COPY app.go /src/
RUN CGO_ENABLED=0 go build -o /bin/demo app.go

FROM scratch
COPY --from=build /bin/demo /bin/demo
EXPOSE 8888
ENTRYPOINT ["/bin/demo"]
