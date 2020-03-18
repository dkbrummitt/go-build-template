# hack/

## Overview

This directory is created to ease/stream line more mundane/repeated tasks
during code maintenance and local testing.
## Data

### DGraph - Golang Graph datastore

More information for [gettng started](https://docs.dgraph.io/get-started) with DGraph

```sh
docker-compose up -d -f hack/dgraph-docker-compose.yml
```

### Postgres

```sh
# Make sure we have something to work with
mkdir -p $HOME/docker/volumes/

# Clean database
rm -rf $HOME/docker/volumes/postgres
docker run --rm   --name pg-docker -e POSTGRES_PASSWORD=docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data  postgres

# Connect to postgres via the shell or your preferred SQL tooling
psql -h localhost -U postgres -d postgres
```

### Mongo

```sh
# Make sure we have something to work with
mkdir -p $HOME/docker/volumes/

# Clean database
mkdir -p $HOME/docker/volumes/
rm -rf $HOME/docker/volumes/mongo
docker run -d -p 27017:27017 --name mongodb mongo -v $HOME/docker/volumes/mongo:/data/db

```
