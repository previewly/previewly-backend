FROM golang:1.23 as GO
WORKDIR /app/
COPY . /app/

RUN apt-get update && \ 
  apt-get install --no-install-recommends -y libvips libvips-dev libvips-tools && \ 
  go mod download && \
  CGO_ENABLED=0 GOOS=linux go build -o /app/previewly-backend

FROM ghcr.io/go-rod/rod:v0.116.2
WORKDIR /app/
COPY --from=GO /app/previewly-backend /app/previewly-backend

RUN apt-get update && \
  apt-get install --no-install-recommends -y libvips && \
  apt-get remove -y automake curl build-essential && \
  apt-get autoremove -y && \
  apt-get autoclean && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

EXPOSE 8000
