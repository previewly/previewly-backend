FROM golang:1.23 as build
WORKDIR /app
COPY go.mod go.sum ./source/
RUN cd source && go mod download
COPY . ./source

RUN cd ./source/ && CGO_ENABLED=0 GOOS=linux go build -o /app/previewly-backend && rm -rf /app/source

FROM ghcr.io/go-rod/rod:v0.116.2

WORKDIR /app
COPY --from=build /app/previewly-backend /app/previewly-backend

EXPOSE 8000
