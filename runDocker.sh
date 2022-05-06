#!/bin/env bash 

sudo docker run -d \
	--name go_test \
	-p 8008:8008 \
	-v /crypto/config/:/crypto/config/ \
	crypto:latest
