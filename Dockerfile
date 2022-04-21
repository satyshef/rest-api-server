ARG APP_NAME="api-server"

FROM golang:1.15-alpine3.13 as builder
WORKDIR /go/src
COPY . .
RUN go get -d -v ./... && \
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${APP_NAME} .

FROM alpine:3.13
COPY --from=builder ./${APP_NAME} ./${APP_NAME}
ENV APP_PATH=APP_NAME \
	CONFIG_PATH=/configs/apiserver.toml

CMD ["/${APP_NAME}", "-c", "{CONFIG_PATH}"]