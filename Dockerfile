FROM golang:1.20.5 as builder
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -a -o go-codewars .

FROM alpine:3.15
COPY --from=builder /build/go-codewars .
COPY --from=builder /build/routes/assets ./routes/assets
EXPOSE 3000
ENTRYPOINT [ "./go-codewars" ]