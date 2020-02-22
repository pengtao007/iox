package crypto

import (
	"testing"
)

func bytesEq(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestExpand32(t *testing.T) {
	src36 := []byte{
		0, 1, 2, 3, 4, 5, 6, 7,
		8, 9, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25,
	}

	src16 := []byte{
		0, 1, 2, 3, 4, 5, 6, 7,
		8, 9, 0xA, 0xB, 0xC, 0xD, 0xE, 0xF,
	}

	src10 := []byte{
		0, 1, 2, 3, 4, 5, 6, 7,
		8, 9,
	}

	var key, iv []byte
	key, iv = expand32(src36)
	if !bytesEq(key, src16) || !bytesEq(iv, []byte{
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
	}) {
		t.Error("src36 error")
	}

	key, iv = expand32(src16)
	if !bytesEq(key, src16) || !bytesEq(iv, []byte{
		0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
		0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	}) {
		t.Error("src16 error")
	}

	key, iv = expand32(src10)
	if !bytesEq(key, append(src10, []byte{
		0x16, 0x16, 0x16, 0x16, 0x16, 0x16,
	}...)) || !bytesEq(iv, []byte{
		0x16, 0x16, 0x16, 0x16, 0x16, 0x16, 0x16, 0x16,
		0x16, 0x16, 0x16, 0x16, 0x16, 0x16, 0x16, 0x16,
	}) {
		t.Error("src10 error")
	}
}

func TestStreamXOR(t *testing.T) {
	cipherA, cipherB, _ := NewCipherPair([]byte("KEY"))
	plain := []byte("testing plain text...")
	output1 := make([]byte, len(plain))
	cipherA.StreamXOR(output1, plain)

	output2 := make([]byte, len(plain))
	cipherB.StreamXOR(output2, output1)

	if !bytesEq(output2, plain) || bytesEq(output1, plain) {
		t.Error("AES-CTR error")
	}
}
