FROM golang:1.14-stretch

WORKDIR /app

COPY . .

RUN go mod download && go get github.com/cespare/reflex

COPY /reflex.conf /

EXPOSE 3001

ENTRYPOINT ["reflex", "-c", "./reflex.conf"]