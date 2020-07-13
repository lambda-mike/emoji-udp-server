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
			var n uint = 3
			emoji := ":ok:"
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

func TestMultiplier(t *testing.T) {
	t.Log("Multiplier")
	{
		t.Log("Given multiplier Transformer and correct Cmd object")
		{
			var n uint = 3
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

func TestMemoryTableTranslator(t *testing.T) {
	t.Log("MemoryTableTranslator")
	{
		t.Log("Given memoryTableTranslator with translation enabled")
		{
			isRaw := false
			t.Log("When correct Cmd object is passed to Transform")
			{
				sut := CreateTranslator(isRaw)
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
			t.Log("Given memoryTableTranslator with translation enabled and Cmd object with uknown Emoji")
			{
				sut := CreateTranslator(isRaw)
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
		t.Log("Given memoryTableTranslator with translation disabled")
		{
			isRaw := true
			t.Log("When correct Cmd object is passed to Transform")
			{
				sut := CreateTranslator(isRaw)
				cmd := contracts.Cmd{7, ":ok:"}
				res, err := sut.Transform(cmd)
				if err != nil {
					t.Fatalf("It should not return err, got: %v", err)
				}
				t.Log("It should not return error")
				okh := ":ok:"
				if res.N != cmd.N {
					t.Fatalf("It should not modify cmd N: %d, got: %d", cmd.N, res.N)
				}
				t.Log("It should not modify cmd N")
				if res.Emoji != okh {
					t.Fatalf("It should return changed Emoji: %v, got: %v", okh, res.Emoji)
				}
				t.Log("It should return changed Emoji")
			}
			t.Log("Given memoryTableTranslator with translation enabled and Cmd object with uknown Emoji")
			{
				sut := CreateTranslator(isRaw)
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
}
