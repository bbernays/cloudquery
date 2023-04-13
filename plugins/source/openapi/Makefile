.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/openapi
	go run main.go doc ../../../website/tables/openapi
	sed 's_(\(.*\))_(../../../../../website/tables/openapi/\1)_' ../../../website/tables/openapi/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/openapi/README.md
	mv ../../../website/tables/openapi/README.md ../../../website/pages/docs/plugins/sources/openapi/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/openapi/*.md
	rm -rf ../../../website/tables/openapi/*.bak

# All gen targets
.PHONY: gen
gen: gen-docs