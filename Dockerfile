FROM node:18 AS frontend_build

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
RUN apt-get update && apt-get install -y \
netcat \
postgresql-client \
curl \
&& apt-get clean \
&& rm -rf /var/lib/apt/lists/*

#Install nodeJS and npm to use serve
RUN curl -fsSL https://deb.nodesource.com/setup_16.x | bash -
RUN apt-get install -y nodejs
RUN npm install -g serve

#Copy the built frontend to the image
COPY --from=frontend_build /app/frontend/dist /usr/share/nginx/html

#Copy the backend binary and .env to the image
COPY --from=backend_build /app/backend/main /usr/local/bin/main

#Copy start.sh to the image and make it executable
COPY start.sh /usr/local/bin/start.sh
RUN chmod +x /usr/local/bin/start.sh

#Expose the port
EXPOSE 2000
EXPOSE 8000

#Start the backend and frontend
ENTRYPOINT [ "/usr/local/bin/start.sh" ]