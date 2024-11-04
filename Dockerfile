From node:18 AS frontend_build

WORKDIR /app/frontend

COPY frontend/package*.json ./
RUN npm install
COPY frontend/ ./
RUN npm run build

#Install serve globally
RUN npm install -g serve

#Stage 2 - build the backend
FROM golang:1.23 AS backend_build

WORKDIR /app/backend

COPY backend/go.mod ./ backend/go.sum ./
RUN go mod download
COPY backend ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#Stage 3 - build the final image
FROM debian:bullseye-slim

#Install required packages