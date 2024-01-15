#!/bin/bash

if [ "$(docker ps -q -f name=apniuniversity-mongodb-1)" ]; then
  docker rm -f apniuniversity-mongodb-1
fi