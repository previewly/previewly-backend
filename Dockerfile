FROM golang:1.22
WORKDIR /app
COPY go.mod go.sum ./source/
RUN cd source && go mod download
COPY . ./source

RUN cd ./source/ && CGO_ENABLED=0 GOOS=linux go build -o /app/wsw-backend && rm -rf /app/source
EXPOSE 8000

#CMD ["/app/wsw-backend"]
