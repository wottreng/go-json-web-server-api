
echo "[-->] sending json data to server"
curl -X POST -H 'Content-Type: application/json' -d '{"temp":"68","humidity%":"45"}' "0.0.0.0:8080?topic=rand"
echo ""
#
echo "[-->] retrieve data from server"
curl "0.0.0.0:8080/?topic=rand&alldata=true" # | jq
echo ""
#
echo "[-->] try to retrieve data for a topic that does not exist"
curl -i -H "Accept: application/json" "0.0.0.0:8080/?topic=doesntExist"
echo ""
