#!/bin/sh

set -e

DATA_FILE="/data/0_0.tigerbeetle"

if [ ! -f "$DATA_FILE" ]; then
    echo "Formatting TigerBeetle data file..."
    /usr/local/bin/tigerbeetle format --cluster=0 --replica=0 --replica-count=1 "$DATA_FILE"
    echo "TigerBeetle data file formatted successfully."
else
    echo "TigerBeetle data file already exists."
fi

echo "Starting TigerBeetle..."
exec /usr/local/bin/tigerbeetle start --addresses=0.0.0.0:3000 "$DATA_FILE" 