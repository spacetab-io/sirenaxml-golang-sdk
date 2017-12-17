package des_test

import (
	"fmt"
	"github.com/tmconsulting/sirenaxml-golang-sdk/des"
	"sirena-agent-go/random"
	"testing"
)

func TestDesEncrypt(t *testing.T) {
	key := []byte(random.String(8))
	origtext := []byte("hello world123563332")

	erytext, err := des.Encrypt(origtext, key)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Encrypted: %v\n", erytext)
	destext, err2 := des.Decrypt(erytext, key)
	if err2 != nil {
		t.Fatal(err2)
	}
	fmt.Println(string(destext))
	fmt.Println(len(origtext), len(string(destext)))
	fmt.Println(string(origtext) == string(destext))
}
