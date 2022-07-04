#!/bin/bash
# stress test file I/O for race conditions
# call in multiple terminals
#echo "[-->] retrieve latest row of data from server"

clear

for i in {1..5}
do
#  echo $i
  curl "localhost:8080/?topic=rand" &
  str="{\"temp\":\"33\",\"count\":\"$i\"}"
  curl -X POST -H 'Content-Type: application/json' -d $str "localhost:8080?topic=rand" &

done
echo "[FINISH]"
