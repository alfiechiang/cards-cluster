FROM golang:1.19

WORKDIR /app
COPY . .
RUN go build -o /api-run

EXPOSE 3000

CMD [ "/api-run" ]