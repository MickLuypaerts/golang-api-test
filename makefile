run:
	go run main.go

getproducts:
	curl localhost:8080 -v | jq

addproduct:
	curl -v localhost:8080 -d "{\"name\": \"test\", \"description\": \"test desc remove this\"}" | jq

updateproduct:
	curl -v localhost:8080/$(id) -d "{\"name\": \"updated name\", \"description\": \"test desc remove this\"}" -XPUT | jq
