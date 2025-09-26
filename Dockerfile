FROM golang:1.25 AS builder
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -a -o go-codewars .

FROM golang:1.25-alpine
COPY --from=builder /build/go-codewars .
COPY --from=builder /build/routes/assets ./routes/assets
EXPOSE 3000
ENTRYPOINT [ "./go-codewars" ]