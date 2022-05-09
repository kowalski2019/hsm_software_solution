#!/bin/env bash 

sudo docker run -d \
	-p 8008:8008 \
	-v /crypto/config/:/crypto/config/ \
	crypto:latest
