.PHONY: buildlocal runlocal

buildlocal:
	CGO_ENABLED=1 go build -tags="nomsgpack sqlite_foreign_keys" .

runlocal:
	CGO_ENABLED=1 go run -tags="nomsgpack sqlite_foreign_keys" .
