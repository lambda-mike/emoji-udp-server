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
				emoji := ":ok:"
				n := 3
				cmd := contracts.Cmd{n, emoji}
				sep := ","
				sut := CreateResponseBuilder(sep)
				res, err := sut.Build(cmd)
				if err != nil {
					t.Fatal("It should not return error, got: ", err)
				}
				t.Log("It should not return error")
				expected := "👌,👌,👌"
				if res != expected {
					t.Fatalf("It should return correct response: %s, got: %s", expected, res)
				}
				t.Log("It should return correct response")
			}
		}
	}
}
