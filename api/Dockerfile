FROM golang:1.20-alpine AS builder


WORKDIR /app

COPY go.mod go.sum ./
COPY *.go ./
COPY app/ app/
COPY handlers/ handlers/
COPY utils/ utils/
COPY model_trained/ model_trained/
COPY docs/ /app/docs/


RUN go mod download
RUN go build -o main .



FROM python:3.9-slim


WORKDIR /app


COPY --from=builder /app/main /app/
COPY app/text_generator.py /app/
COPY app/requirements.txt /app/
COPY model_trained/ /app/api/model_trained


RUN pip install -r requirements.txt
RUN pip install torch


EXPOSE 8080


CMD ["./main"]
