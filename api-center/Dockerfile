FROM golang:1.19

WORKDIR /app
COPY . .
RUN go build -o /api-run

ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait
EXPOSE 3000

CMD [ "/api-run" ]