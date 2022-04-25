ARG APP_NAME="app"
ARG CONF_NAME="apiserver.toml"
ARG BUILD_DIR="/usr/local/go/src/app"

# builder
FROM golang:1.18-alpine3.15 as builder
LABEL stage=builder
ARG APP_NAME
ARG BUILD_DIR
WORKDIR ${BUILD_DIR}
COPY go.mod go.sum ./
RUN apk add git && \
    apk add make && \
    go mod download && \
    go mod verify
COPY . .
RUN  go mod tidy && make -e APP_PATH=${APP_NAME}

# finish
FROM alpine:3.15
ARG APP_NAME
ARG CONF_NAME
ARG BUILD_DIR
ENV APP_PATH="/usr/local/bin/$APP_NAME"
ENV CONF_PATH="/etc/${CONF_NAME}"
COPY --from=builder ${BUILD_DIR}/${APP_NAME} ${APP_PATH}
COPY --from=builder ${BUILD_DIR}/configs/apiserver.toml ${CONF_PATH}
#RUN echo -e "#!/bin/sh\n${APP_PATH} -c ${CONF_PATH}" > /entrypoint.sh && \
#    chmod +x /entrypoint.sh
#CMD [ "/entrypoint.sh" ]
CMD ${APP_PATH} -c ${CONF_PATH}