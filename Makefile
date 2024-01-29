air:
	air -c .air.toml

clone:
	./scripts/setup-clone-only.sh

generate:
	go run ./cmd/app/main.go

setup-bff:
	./scripts/setup-bff.sh

run-bff:
	./scripts/run-bff.sh
