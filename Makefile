clean: 
	rm -rf api

gen: clean
	java -jar swagger-codegen-cli.jar generate -i ./openapi.json -l dynamic-html -o api/
