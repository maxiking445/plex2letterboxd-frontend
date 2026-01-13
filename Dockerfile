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


# ====== STAGE 3: Build Python Script Env ======
FROM python:3.9-slim AS python-builder

WORKDIR /script

COPY backend/script/ .

RUN python -m venv /script/env && \
    /script/env/bin/pip install --no-cache-dir --upgrade pip && \
    /script/env/bin/pip install --no-cache-dir .

# ====== STAGE 4: Runtime ======
FROM python:3.9-slim

WORKDIR /app
RUN apt-get update && apt-get install -y nginx ca-certificates && rm -rf /var/lib/apt/lists/*
RUN mkdir -p /run/nginx

RUN mkdir -p /app/data

COPY --from=builder /myapp /myapp
COPY --from=vue-builder /app/dist /usr/share/nginx/html
COPY frontend/nginx.conf /etc/nginx/nginx.conf
COPY --from=python-builder /script ./script

EXPOSE 80 8080


CMD sh -c "/myapp & nginx -g 'daemon off;'"