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
			t.Log("When passed proper cmd with n=0")
			{
				emoji := "👌"
				n := 0
				cmd := contracts.Cmd{n, emoji}
				sep := ","
				sut := CreateResponseBuilder(sep)
				res := sut.Build(cmd)
				expected := ""
				if res != expected {
					t.Fatalf("It should return correct response: %s, got: %s", expected, res)
				}
				t.Log("It should return correct response")
			}
			t.Log("When passed proper cmd with n=1")
			{
				emoji := "👌"
				n := 1
				cmd := contracts.Cmd{n, emoji}
				sep := ","
				sut := CreateResponseBuilder(sep)
				res := sut.Build(cmd)
				expected := "👌"
				if res != expected {
					t.Fatalf("It should return correct response: %s, got: %s", expected, res)
				}
				t.Log("It should return correct response")
			}
			t.Log("When passed proper cmd")
			{
				emoji := "👌"
				n := 2
				cmd := contracts.Cmd{n, emoji}
				sep := ","
				sut := CreateResponseBuilder(sep)
				res := sut.Build(cmd)
				expected := "👌,👌"
				if res != expected {
					t.Fatalf("It should return correct response: %s, got: %s", expected, res)
				}
				t.Log("It should return correct response")
			}
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
			t.Log("When passed proper cmd with empty emoji")
			{
				emoji := ""
				n := 3
				cmd := contracts.Cmd{n, emoji}
				sep := ","
				sut := CreateResponseBuilder(sep)
				res := sut.Build(cmd)
				expected := ""
				if res != expected {
					t.Fatalf("It should return correct response: %s, got: %s", expected, res)
				}
				t.Log("It should return correct response")
			}
			t.Log("When passed incorrect cmd with negative N")
			{
				emoji := ":ok:"
				n := -3
				cmd := contracts.Cmd{n, emoji}
				sep := ","
				sut := CreateResponseBuilder(sep)
				res := sut.Build(cmd)
				expected := ""
				if res != expected {
					t.Fatalf("It should return correct response: %s, got: %s", expected, res)
				}
				t.Log("It should return correct response")
			}
			t.Log("When passed correct cmd without separator")
			{
				emoji := ":ok:"
				n := 3
				cmd := contracts.Cmd{n, emoji}
				sep := ""
				sut := CreateResponseBuilder(sep)
				res := sut.Build(cmd)
				expected := ":ok::ok::ok:"
				if res != expected {
					t.Fatalf("It should return correct response: %s, got: %s", expected, res)
				}
				t.Log("It should return correct response")
			}
		}
	}
}
