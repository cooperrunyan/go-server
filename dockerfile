FROM golang:1.18.2-alpine

ARG port=3000

WORKDIR /app

COPY . .

RUN go get -d -v .

RUN go install -v .

EXPOSE ${port}
ENV PORT=${port}

CMD ["http"]
