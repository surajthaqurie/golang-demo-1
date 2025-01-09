# FROM golang:1.12.0-alpine3.9
FROM golang:1.12.0 as builder


RUN mkdir /app
ADD . /app
WORKDIR /app

# RUN go build -o main .
RUN CGO_ENABLE=0 GOOS=linux go build -o main ./...

FROM alpine:latest as production
COPY --from=builder ./app .
CMD [ "./main" ]

