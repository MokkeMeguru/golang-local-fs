FROM golang:1.16.3
LABEL maintainer="MokkeMeguru <meguru.mokke@gmail.com>"
ENV APP_PATH /app
ENV GIN_MODE=release
ENV LOCAL_FILE_ROOT /local_fs
# This File Storage allow overwrite duplicate files?
ARG OW_FILE false

WORKDIR $APP_PATH
COPY . $APP_PATH
RUN make server
CMD [ "/app/bin/server" ]
