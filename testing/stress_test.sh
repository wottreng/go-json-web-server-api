#!/bin/bash
# stress test file I/O for race conditions
# call in multiple terminals
#echo "[-->] retrieve latest row of data from server"

clear

for i in {1..10}
do
  echo $i
  curl -X POST -H 'Content-Type: application/json' -d '{"temp":"68","humidity%":"45"}' "localhost:8080?topic=rand" &
  curl "localhost:8080/?topic=rand" & # | jq
done
echo "[FINISH]"
