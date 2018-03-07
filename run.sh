#!/bin/sh
docker build -t sumoproto:0.1 .
docker-compose build && docker-compose up