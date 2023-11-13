FROM golang:1.21
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /golang-clean-architecture
EXPOSE 3001
CMD ["/golang-clean-architecture"]