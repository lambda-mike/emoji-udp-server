package service

import (
	"errors"
	"github.com/emoji-udp-server/contracts"
	"testing"
)

type mockTranformer struct {
	counter int
}

func (t *mockTranformer) Transform(c contracts.Cmd) (contracts.Cmd, error) {
	t.counter += 1
	return c, nil
}

type failingMockTranformer struct{}

func (t *failingMockTranformer) Transform(c contracts.Cmd) (contracts.Cmd, error) {
	return c, errors.New("FAIL from failingMockTranformer")
}

func TestBuild(t *testing.T) {
	t.Log("Build")
	{
		t.Log("Given EmojiConcatenator")
		{
			t.Log("When passed proper cmd with n=0")
			{
				var n uint = 0
				emoji := "👌"
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
				var n uint = 1
				emoji := "👌"
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
				var n uint = 2
				emoji := "👌"
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
				var n uint = 3
				emoji := "👌"
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
				var n uint = 3
				emoji := ""
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
				var n uint = 3
				emoji := ":ok:"
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

func TestTransformCmd(t *testing.T) {
	t.Log("transformCmd")
	{
		t.Log("Given EmojiService with no transformers")
		{
			t.Log("When cmd and nil err are passed")
			{
				sut := &EmojiService{}
				cmd := contracts.Cmd{5, ":ok:"}
				res, err := sut.transformCmd(cmd, nil)
				if err != nil {
					t.Fatalf("It should not return error, got: %v", err)
				}
				t.Log("It should not return error")
				if res != cmd || res.N != cmd.N || res.Emoji != cmd.Emoji {
					t.Fatalf("It should return exact same cmd: %v, got: %v", cmd, res)
				}
				t.Log("It should return exact same cmd")
			}
			t.Log("When cmd and not nil err are passed")
			{
				sut := &EmojiService{}
				cmd := contracts.Cmd{5, ":ok:"}
				err := errors.New("FAIL")
				_, resErr := sut.transformCmd(cmd, err)
				if resErr == nil {
					t.Fatalf("It should return error: %v, got: %v", err, resErr)
				}
				t.Log("It should return error")
			}
		}
		t.Log("Given EmojiService with one transformer")
		{
			t.Log("When cmd and nil err are passed")
			{
				trans := mockTranformer{}
				sut := &EmojiService{}
				sut.transformers = []contracts.CmdTransformer{&trans}
				cmd := contracts.Cmd{5, ":ok:"}
				res, err := sut.transformCmd(cmd, nil)
				if err != nil {
					t.Fatalf("It should not return error, got: %v", err)
				}
				t.Log("It should not return error")
				if res != cmd || res.N != cmd.N || res.Emoji != cmd.Emoji {
					t.Fatalf("It should return exact same cmd: %v, got: %v", cmd, res)
				}
				t.Log("It should return exact same cmd")
				if trans.counter != 1 {
					t.Fatalf("It should execute transformer exactly once, got: %d", trans.counter)
				}
				t.Log("It should execute transformer exactly once")
			}
			t.Log("When cmd and not nil err are passed")
			{
				trans := mockTranformer{}
				sut := &EmojiService{}
				sut.transformers = []contracts.CmdTransformer{&trans}
				cmd := contracts.Cmd{5, ":ok:"}
				err := errors.New("FAIL")
				_, resErr := sut.transformCmd(cmd, err)
				if resErr == nil {
					t.Fatalf("It should return error: %v, got: %v", err, resErr)
				}
				t.Log("It should return error")
				if trans.counter != 0 {
					t.Fatalf("It should not execute transformer, executions number: %d", trans.counter)
				}
				t.Log("It should not execute transformer at all")
			}
		}
		t.Log("Given EmojiService with two transformers")
		{
			t.Log("When cmd and nil err are passed")
			{
				trans1 := mockTranformer{}
				trans2 := mockTranformer{}
				sut := &EmojiService{}
				sut.transformers = []contracts.CmdTransformer{&trans1, &trans2}
				cmd := contracts.Cmd{5, ":ok:"}
				res, err := sut.transformCmd(cmd, nil)
				if err != nil {
					t.Fatalf("It should not return error, got: %v", err)
				}
				t.Log("It should not return error")
				if res != cmd || res.N != cmd.N || res.Emoji != cmd.Emoji {
					t.Fatalf("It should return exact same cmd: %v, got: %v", cmd, res)
				}
				t.Log("It should return exact same cmd")
				if trans1.counter != 1 {
					t.Fatalf("It should execute transformer1 exactly once, got: %d", trans1.counter)
				}
				t.Log("It should execute transformer1 exactly once")
				if trans2.counter != 1 {
					t.Fatalf("It should execute transformer2 exactly once, got: %d", trans2.counter)
				}
				t.Log("It should execute transformer2 exactly once")
			}
		}
		t.Log("Given EmojiService with two transformers: one failing")
		{
			t.Log("When cmd and nil err are passed")
			{
				trans1 := failingMockTranformer{}
				trans2 := mockTranformer{}
				sut := &EmojiService{}
				sut.transformers = []contracts.CmdTransformer{&trans1, &trans2}
				cmd := contracts.Cmd{5, ":ok:"}
				_, err := sut.transformCmd(cmd, nil)
				if err == nil {
					t.Fatalf("It should return error, got: nil")
				}
				t.Log("It should return error", err)
				if trans2.counter != 0 {
					t.Fatalf("It should not execute transformer2, got number of executions: %d", trans2.counter)
				}
				t.Log("It should not execute transformer2 at all")
			}
		}
	}
}
