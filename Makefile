air:
	air -c .air.toml

run:
	go run ./cmd/app/main.go

setup-clone-only:
	./scripts/setup-clone-only.sh