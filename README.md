# Concordia Boats Web Site

This repo has files used to create [concordiaboats.com](https://www.concordiaboats.com/).

## Files & Directories

```sh
  Makefile              - Makefile with targets for running site locally and installing.
  README.md             - this file.
  cmd/                  - Go programs.
```

## Database (PostgreSQL)

```sh

apt-
apt-get install emacs

docker volume create concordia

docker run -v concordia:/opt/concordia --name 14.1-alpine -e POSTGRES_PASSWORD=dbpw -d postgres
--mount type=bind,source=$(shell pwd)/src/var/nyt/www/api/internal,target=/nyt/php/www-internal \

docker exec -it 14.1-alpine /bin/sh

psql -U postgres
```

https://pkg.go.dev/github.com/mattn/go-sqlite3
https://sqlitebrowser.org/

Boat images
Boat info
Boat links
Boat owner info
