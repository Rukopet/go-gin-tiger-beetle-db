FROM ghcr.io/tigerbeetle/tigerbeetle:latest

COPY scripts/init-tigerbeetle.sh /usr/local/bin/init-tigerbeetle.sh
RUN chmod +x /usr/local/bin/init-tigerbeetle.sh

ENTRYPOINT ["/usr/local/bin/init-tigerbeetle.sh"] 