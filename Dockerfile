FROM golang:1.23 as build
WORKDIR /app
COPY go.mod go.sum ./source/
RUN cd source && go mod download
COPY . ./source

RUN cd ./source/ && CGO_ENABLED=0 GOOS=linux go build -o /app/previewly-backend && rm -rf /app/source

#CMD ["/app/wsw-backend"]

FROM ghcr.io/go-rod/rod

COPY --from=build /app/previewly-backend /app/previewly-backend

EXPOSE 8000
WORKDIR /app
