FROM golang:1.12.4-alpine3.9 as build
RUN apk add --no-cache make
WORKDIR /go/src/github.com/abhisheknsit/go-pinger
ADD . .
RUN make build

FROM alpine:3.9
COPY --from=build /go/src/github.com/abhisheknsit/go-pinger/bin/go-pinger /bin/go-pinger
CMD ["go-pinger"]
