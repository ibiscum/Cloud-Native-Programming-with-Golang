## Install mongodb-sh on Debian

    wget -qO - https://www.mongodb.org/static/pgp/server-7.0.asc | sudo tee /etc/apt/trusted.gpg.d/server-7.0.asc

    echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/debian bookworm/mongodb-org/7.0 main" | sudo tee /etc/apt/sources.list.d/mongodb-org-7.0.list

    sudo apt update
    sudo apt-get install -y mongodb-mongosh

## Run mongod as Docker container

    docker pull mongo:7.0
    docker run --rm -d -p 127.0.0.1:27017:27017/tcp mongo:7.0

## Start service

Within chapter02 directory execute:

    go run eventsservice/main.go
