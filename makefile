run:
	go run main.go

getproducts:
	curl localhost:8080 -v | jq

addproduct:
	curl -v localhost:8080 -d "{\"name\": \"test\", \"description\": \"test desc remove this\", \"price\": 1, \"sku\": \"abc-abc-1\"}" | jq

updateproduct:
	curl -v localhost:8080/$(id) -d "{\"name\": \"updated name\", \"description\": \"test desc remove this\"}" -XPUT | jq
