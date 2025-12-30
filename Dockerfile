FROM golang:1.22.3-alpine
LABEL description="image for the project ascii-art-web-dockerize"
LABEL authors="elmehdi rezoug, aymane bouzerda"
WORKDIR /app
COPY . .
RUN go build -o app
EXPOSE 8080
ENTRYPOINT ["./app"]