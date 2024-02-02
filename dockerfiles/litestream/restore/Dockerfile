FROM litestream/litestream:0.3.13

COPY dockerfiles/litestream/litestream.yaml /etc/litestream.yaml

ENTRYPOINT ["/bin/sh"]
CMD ["-c", "rm -f /data/todo.db && /usr/local/bin/litestream restore -config /etc/litestream.yaml /data/todo.db"]
