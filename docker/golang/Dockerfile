FROM golang:1.22.3-bullseye


RUN mkdir -p /usr/local/src/app/
WORKDIR /usr/local/src/app/

# env
ENV CONFIG_PATH=./config/config.yaml

# dependecies
COPY ../../app/go.mod ../../app/go.sum ./
RUN go mod download

# build
COPY ../../app ./
RUN go build -o ./bin/app ./cmd/app/main.go

# run
CMD ["./bin/app"]
