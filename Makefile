deps:
	go mod vendor

key_with_secret:
	openssl genrsa -aes128 -passout pass:secret -out crypt/private.secret.key.pem  1024
	openssl rsa -in crypt/private.secret.key.pem -passin pass:secret -pubout -out crypt/public.secret.key.pem

key_without_secret:
	openssl genrsa -out crypt/private.key.pem  1024
	openssl rsa -in crypt/private.key.pem -pubout -out crypt/public.key.pem

.PHONY: key_with_secret
test_prepare: key_with_secret key_without_secret

test: test_prepare
	go test ./... -coverprofile=c.out -v
	go tool cover -html=c.out -o coverage.html