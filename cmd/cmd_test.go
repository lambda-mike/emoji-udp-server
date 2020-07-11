package cmd

import (
	"github.com/emoji-udp-server/contracts"
	"testing"
)

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
