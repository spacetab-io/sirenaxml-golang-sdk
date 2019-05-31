T = '*'

.PHONY:test
test:
	go test ./... -v -run=$T -mod=vendor -count=1