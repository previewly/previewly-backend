FROM golang:1.23 AS golang
RUN apt-get update && \ 
  apt-get install --no-install-recommends -y libvips libvips-dev libvips-tools  

WORKDIR /app/
COPY . /app/

RUN go mod download && GOOS=linux go build -o /app/previewly-backend

FROM ghcr.io/go-rod/rod:v0.116.2
RUN apt-get update && \
  apt-get install --no-install-recommends -y libvips libvips-tools

WORKDIR /app/
COPY --from=golang /app/previewly-backend /app/previewly-backend

EXPOSE 8000
