# Our builder image used to build the GO binary
FROM golang:1.15.0-alpine as builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go mod download

RUN CGO_ENABLE=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Our production image used to run our app
FROM alpine:latest as production
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache git make musl-dev go
COPY --from=builder ./app/main .

#Configure GO
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

# More here

CMD [ "./main" ]

