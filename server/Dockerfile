FROM golang:1.14.0 AS builder
WORKDIR /go/src/app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /server .

FROM scratch
COPY --from=builder ./server ./
EXPOSE 3000
ENTRYPOINT ["./server"]