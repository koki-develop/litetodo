#!/bin/sh

set -e

rm -f /app/todo.db
litestream restore -if-replica-exists -config /etc/litestream.yml /app/todo.db
litestream replicate -exec /app/app -config /etc/litestream.yml
