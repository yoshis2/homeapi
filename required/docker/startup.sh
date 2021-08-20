#!/bin/bash

if [ ! -e "required/logs/access.log" ]; then
    mkdir -p required/logs
    touch required/logs/access.log
    echo "created access log file"
fi

swag init
fresh -c required/docker/runner.conf

# dlv debug --headless --listen=:5050 --log --api-version=2
