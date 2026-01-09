clean:
	rm -rf api

gen: clean
	java -jar openapi-generator-cli.jar generate -g markdown -i ./openapi.json -o api/
