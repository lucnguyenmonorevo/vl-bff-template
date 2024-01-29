air:
	air -c .air.toml

generate:
	go run ./cmd/app/main.go

clone:
	./scripts/setup-clone-only.sh

run-bff:
	./scripts/run-bff.sh

setup-bff:
	./scripts/setup-bff.sh