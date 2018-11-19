FROM golang:1.11.2-stretch as builder

LABEL maintainer "vegard.skaansar@gmail.com"

RUN go get github.com/VegardSkaansar/Announcement-Service

WORKDIR /go/src/github.com/VegardSkaansar/goprojects/announce

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o announcement


FROM scratch

LABEL maintainer "vegard.skaansar@gmail.com"

WORKDIR /

COPY --from=builder /go/src/github.com/VegardSkaansar/goprojects/announce/announcement .

CMD ["/announcement"]