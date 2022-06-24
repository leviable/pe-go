package main

import (
	"fmt"
	"testing"
)

func TestDecryptMsg(t *testing.T) {
	examples := []struct {
		msg  string
		key  []byte
		want string
	}{
		{msg: string(rune(65)), key: []byte{byte(rune(42))}, want: string(rune(107))},
		{msg: string(rune(107)), key: []byte{byte(rune(42))}, want: string(rune(65))},
	}

	for _, tt := range examples {
		t.Run(fmt.Sprintf("msg %s -> key %s", tt.msg, tt.key), func(t *testing.T) {
			got := decryptMsg(tt.msg, tt.key)
			want := tt.want

			if got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
