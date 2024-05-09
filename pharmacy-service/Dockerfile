FROM golang:1.22.3-alpine AS build
WORKDIR /app
ADD ./app .

RUN go mod download && go build -o /bin/main ./cmd/main.go

FROM alpine:3.9
RUN apk --no-cache add ca-certificates
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip
COPY --from=build /bin/main /bin/main
ENTRYPOINT ["/bin/main"]
