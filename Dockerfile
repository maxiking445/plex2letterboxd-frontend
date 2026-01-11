# ====== STAGE 1: Build Go Binary ======
FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o /myapp .

# ====== STAGE 2: Build Vue Frontend ======
FROM node:22-alpine AS vue-builder
WORKDIR /app

COPY frontend/package*.json ./
RUN npm install

COPY frontend/ .
RUN npm run build


# ====== STAGE 3: Build Vue Frontend ======
FROM python:3.9-slim

WORKDIR /app/backend

COPY backend/script ./script
COPY backend/script/config.ini ./config.ini
COPY backend/script/pyproject.toml ./pyproject.toml
COPY backend/script/plex2letterboxd ./plex2letterboxd

RUN python -m venv env


RUN ./env/bin/pip install --no-cache-dir .

# ====== STAGE 4: Runtime ======
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache nginx ca-certificates
RUN mkdir -p /run/nginx

COPY --from=builder /myapp /myapp
COPY --from=vue-builder /app/dist /usr/share/nginx/html
COPY frontend/nginx.conf /etc/nginx/nginx.conf

EXPOSE 80 8080

CMD sh -c "/myapp & nginx -g 'daemon off;'"