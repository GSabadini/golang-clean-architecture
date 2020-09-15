FROM golang:1.14-stretch

WORKDIR /app

COPY . .

RUN go mod download
RUN go get github.com/cespare/reflex

COPY /reflex.conf /

ENTRYPOINT ["reflex", "-c", "./reflex.conf"]