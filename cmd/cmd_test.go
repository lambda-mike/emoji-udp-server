package cmd

import (
	"fmt"
	"github.com/emoji-udp-server/contracts"
	"testing"
)

func TestParse(t *testing.T) {
	t.Log("Parse")
	{
		t.Log("Given correct cmd string")
		{
			emoji := ":ok:"
			n := 3
			str := fmt.Sprintf("%d %s", n, emoji)
			sut := CreateParser()
			cmd, err := sut.Parse(str)
			if err != nil {
				t.Fatal("It should not return error, got: ", err)
			}
			t.Log("It should not return error")
			if cmd.N != n {
				t.Fatalf("It should return Cmd with correct N: %v, got: %v", n, cmd.N)
			}
			t.Log("It should return Cmd with correct N")
			if cmd.Emoji != emoji {
				t.Fatalf("It should return Cmd with correct Emoji: %s, got: %s", emoji, cmd.Emoji)
			}
			t.Log("It should return Cmd with correct Emoji")
		}
	}
}

func TestCreateTranslator(t *testing.T) {
	t.Log("CreateTranslator")
	{
		t.Log("Given boolean flag with false value")
		{
			raw := false
			sut := CreateTranslator(raw)
			if _, ok := sut.(*memoryTableTranslator); !ok {
				t.Fatalf("It should return correct type of translator: memoryTableTranslator, got sth else")
			}
			t.Log("It should return correct type of translator (memoryTableTranslator)")
		}
		t.Log("Given boolean flag with true value")
		{
			raw := true
			sut := CreateTranslator(raw)
			if _, ok := sut.(*IdentityTranslator); !ok {
				t.Fatalf("It should return correct type of translator: IdentityTranslator, got sth else")
			}
			t.Log("It should return correct type of translator (IdentityTranslator)")
		}
	}
}

func TestMultiplier(t *testing.T) {
	t.Log("Multiplier")
	{
		t.Log("Given multiplier Transformer and correct Cmd object")
		{
			n := 3
			sut := CreateMultiplier(n)
			cmd := contracts.Cmd{7, ":myemoji:"}
			exp := cmd.N * n
			res, err := sut.Transform(cmd)
			if err != nil {
				t.Fatalf("It should not return err, got: %v", err)
			}
			t.Log("It should not return error")
			if res.N != exp {
				t.Fatalf("It should return cmd with properly modified N: %d, got: %d", exp, res.N)
			}
			t.Log("It should return cmd with properly modified N")
			if res.Emoji != cmd.Emoji {
				t.Fatalf("It should not modify cmd Emoji: %v, to: %v", cmd.Emoji, res.Emoji)
			}
			t.Log("It should not modify cmd Emoji")
		}
	}
}

func TestIdentityTranslator(t *testing.T) {
	t.Log("IdentityTranslatort")
	{
		t.Log("Given identity Translator and correct Cmd object")
		{
			sut := IdentityTranslator{}
			cmd := contracts.Cmd{7, ":myemoji:"}
			res, err := sut.Transform(cmd)
			if err != nil {
				t.Fatalf("It should not return err, got: %v", err)
			}
			t.Log("It should not return error")
			if res.N != cmd.N {
				t.Fatalf("It should not modify cmd N: %d, got: %d", cmd.N, res.N)
			}
			t.Log("It should not modify cmd N")
			if res.Emoji != cmd.Emoji {
				t.Fatalf("It should not modify cmd Emoji: %v, to: %v", cmd.Emoji, res.Emoji)
			}
			t.Log("It should not modify cmd Emoji")
		}
	}
}

func TestMemoryTableTranslator(t *testing.T) {
	t.Log("MemoryTableTranslator")
	{
		t.Log("Given memoryTableTranslator and correct Cmd object")
		{
			sut := CreateMemoryTableTranslator()
			cmd := contracts.Cmd{7, ":ok:"}
			res, err := sut.Transform(cmd)
			if err != nil {
				t.Fatalf("It should not return err, got: %v", err)
			}
			t.Log("It should not return error")
			okh := "👌"
			if res.N != cmd.N {
				t.Fatalf("It should not modify cmd N: %d, got: %d", cmd.N, res.N)
			}
			t.Log("It should not modify cmd N")
			if res.Emoji != okh {
				t.Fatalf("It should return changed Emoji: %v, got: %v", okh, res.Emoji)
			}
			t.Log("It should return changed Emoji")
		}
		t.Log("Given memoryTableTranslator and Cmd object with uknown Emoji")
		{
			sut := CreateMemoryTableTranslator()
			cmd := contracts.Cmd{7, ":unknown:"}
			res, err := sut.Transform(cmd)
			if err == nil {
				t.Fatalf("It should return err, got: nil")
			}
			t.Log("It should return error", err)
			if res.N != cmd.N {
				t.Fatalf("It should not modify cmd N: %d, got: %d", cmd.N, res.N)
			}
			t.Log("It should not modify cmd N")
			if res.Emoji != cmd.Emoji {
				t.Fatalf("It should not modify cmd Emoji: %v, to: %v", cmd.Emoji, res.Emoji)
			}
			t.Log("It should not modify cmd Emoji")
		}
	}
}
