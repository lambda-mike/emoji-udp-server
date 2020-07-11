package cmd

import (
	"github.com/emoji-udp-server/contracts"
	"testing"
)

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
			res := sut.Transform(cmd)
			exp := cmd.N * n
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
			res := sut.Transform(cmd)
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
			res := sut.Transform(cmd)
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
			res := sut.Transform(cmd)
			if res.N != cmd.N {
				t.Fatalf("It should not modify cmd N: %d, got: %d", cmd.N, res.N)
			}
			t.Log("It should not modify cmd N")
			t.Fatalf("TODO It should return err instead of cmd")
		}
	}
}
