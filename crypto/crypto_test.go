package crypto

import (
	"testing"
)

func TestHash(t *testing.T) {
	plain := "alfa"
	hash, err := HashPassword(plain)
	if err != nil {
		t.Log(err)
		return
	}

	if CheckPasswordHash(plain, hash) {
		t.Log("ok")
	} else {
		t.Log("invalid")
	}

}
