FROM golang:1.14.3-alpine AS build-env

WORKDIR /src
ENV CGO_ENABLED=0
COPY . .
RUN GOOS=linux go build -o /out/testwebservice .

FROM alpine:latest

COPY --from=build-env /out/testwebservice /
ENTRYPOINT ["/testwebservice"]
EXPOSE 3000