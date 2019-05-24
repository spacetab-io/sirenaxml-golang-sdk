T = '*'

.PHONY:test
test:
	KEYS_PATH=$$(pwd)/keys go test ./service -v -run=$T -mod=vendor -count=1