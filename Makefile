##
# Golang Local FileSever
#
# communicate local folder's file
# using http request
#
# this repo for the dev/local development file server
#   which will be replaced by google cloud storage
#
# @file
# @version 0.1

server:
	go build -o bin/server cmd/server.go

.phony: clean
clean:
	rm -rf bin/*

# end
