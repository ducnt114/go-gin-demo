package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	hashedPass := HashPassword("pass123", "randomSalt")
	if hashedPass != "be2b5f4315cfa0167809ce2637a492c9c5411369a4b700dfb7acb7277fc4b9f9018ad093f0d23e43b0002ce8cf64479ed9bc0d4031e13efd96638b73fc8dd04f" {
		t.Fail()
	}
}
