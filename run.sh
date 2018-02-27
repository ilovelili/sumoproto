#!/bin/sh

sudo docker build -t sumoproto:0.1 .

sudo docker-compose build && sudo docker-compose up
