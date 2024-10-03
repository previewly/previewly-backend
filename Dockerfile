FROM golang:1.23 as GO
WORKDIR /app/
COPY . /app/
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o /app/previewly-backend

FROM ghcr.io/go-rod/rod:v0.116.2
WORKDIR /app/
COPY --from=GO /app/previewly-backend /app/previewly-backend

EXPOSE 8000
