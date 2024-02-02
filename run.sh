#!/bin/sh

set -e

rm -f /data/todo.db
litestream restore -if-replica-exists -config /etc/litestream.yml /data/todo.db
litestream replicate -exec /app/app -config /etc/litestream.yml
