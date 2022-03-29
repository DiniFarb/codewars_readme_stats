FROM golang:1.18.0-alpine3.15 as builder
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -a -o go-codewars .

FROM alpine:3.15
COPY --from=builder /build/go-codewars .
COPY --from=builder /build/codewars/templates ./codewars/templates
EXPOSE 3000
ENTRYPOINT [ "./go-codewars" ]