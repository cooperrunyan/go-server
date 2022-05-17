FROM golang:1.18.2-alpine as dev

RUN apk update; apk add curl

ARG port
EXPOSE ${port}
ENV PORT=${port}

WORKDIR /app

ADD . .
RUN go get -d -v .

FROM dev as prod
RUN go install -v .
