run:
	go run main.go

getproducts:
	curl localhost:8080 -XGET -v | jq

addproduct:
	curl -v localhost:8080 -d "{\"name\": \"test\", \"description\": \"test desc remove this\"}" | jq