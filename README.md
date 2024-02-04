# Start

```console
$ docker compose up --build
```

# Usage

## Create a task

```console
$ curl http://localhost:8080/tasks \
    -X POST \
    -H 'Content-Type: application/json' \
    --data '{ "title": "aaa" }'
```

## Get tasks

```console
$ curl http://localhost:8080/tasks
```

## Get a task

```console
$ curl http://localhost:8080/tasks/1
```

## Update a task

```console
$ curl http://localhost:8080/tasks/1 \
    -X PATCH \
    -H 'Content-Type: application/json' \
    --data '{ "title": "updated", "completed": true }'
```

## Delete a task

```console
$ curl http://localhost:8080/tasks/1 \
    -X DELETE
```

# LICENSE

[MIT](./LICENSE)
