package service

import (
	"github.com/emoji-udp-server/contracts"
	"testing"
)

func TestBuild(t *testing.T) {
	t.Log("Build")
	{
		t.Log("Given EmojiConcatenator")
		{
			t.Log("When passed proper cmd")
			{
				emoji := "👌"
				n := 3
				cmd := contracts.Cmd{n, emoji}
				sep := ","
				sut := CreateResponseBuilder(sep)
				res := sut.Build(cmd)
				expected := "👌,👌,👌"
				if res != expected {
					t.Fatalf("It should return correct response: %s, got: %s", expected, res)
				}
				t.Log("It should return correct response")
			}

		}
	}
}
