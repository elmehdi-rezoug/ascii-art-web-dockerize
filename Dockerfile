FROM golang:1.22.3-alpine
LABEL description="ascii-art-web-dockerize"
WORKDIR /app
COPY . .
EXPOSE 8080
ENTRYPOINT ["go", "run", "."]