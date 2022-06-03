
echo "sending json data to server"
curl -X POST -H 'Content-Type: application/json' -d '{"user":"mark","temp":"68"}' "0.0.0.0:8080?topic=rand"

echo "retreive data from server"
curl "0.0.0.0:8080/?topic=rand&alldata=false"
