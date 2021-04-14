FROM golang:1.16.3
LABEL maintainer="MokkeMeguru <meguru.mokke@gmail.com>"
ENV APP_PATH "/app"
ENV LOCAL_FILE_ROOT /local_fs
ARG OW_FILE false
WORKDIR $LOCAL_FILE_ROOT
RUN git clone https://github.com/MokkeMeguru/golang-local-fs $APP_PATH
WORKDIR APP_PATH
RUN make server
CMD [ "sh" "$APP_PATH/bin/sever" ]
