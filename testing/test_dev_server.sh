
echo "sending json data to server"
curl -X POST -H 'Content-Type: application/json' -d '{"user":"mark","temp":"68"}' "0.0.0.0:8080?topic=rand"

echo "retrieve data from server"
curl "0.0.0.0:8080/?topic=rand&alldata=true" # | jq

printf "\ntry to retrieve data for a topic that does not exist \n"
curl "0.0.0.0:8080/?topic=doesntExist"
