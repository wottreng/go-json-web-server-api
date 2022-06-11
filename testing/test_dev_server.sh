#!/bin/bash
echo "[-->] try url with no args"
curl -i "localhost:8080/"  # not an endpoint
echo ""
#
echo "[-->] try url with help arg"
curl -i "localhost:8080/?help"  # Usage: /?topic=<topic>
echo ""
#
echo "[-->] try to retrieve data for a topic that does not exist"
curl -i -H "Accept: application/json" "localhost:8080/?topic=doesntExist"
echo ""
#
echo "[-->] sending json data to server"
curl -X POST -H 'Content-Type: application/json' -d '{"temp":"68","humidity%":"45"}' "localhost:8080?topic=rand"
echo ""
#
echo "[-->] retrieve data from server"
curl "localhost:8080/?topic=rand&alldata=true" # | jq
echo ""
