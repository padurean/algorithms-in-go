package findwordfromlinkedletters_test

import (
	"testing"

	findwordfromlinkedletters "github.com/padurean/algorithms-in-go/find_word_from_linked_letters"
	"github.com/stretchr/testify/require"
)

func TestFindWord(t *testing.T) {
	testCases := []struct {
		rules []string
		want  string
	}{
		{
			[]string{"P>E", "E>R", "R>U"},
			"PERU",
		},
		{
			[]string{"I>N", "A>I", "P>A", "S>P"},
			"SPAIN",
		},
		{
			[]string{"U>N", "G>A", "R>Y", "H>U", "N>G", "A>R"},
			"HUNGARY",
		},
		{
			[]string{"I>F", "W>I", "S>W", "F>T"},
			"SWIFT",
		},
		{
			[]string{"R>T", "A>L", "P>O", "O>R", "G>A", "T>U", "U>G"},
			"PORTUGAL",
		},
		{
			[]string{"W>I", "R>L", "T>Z", "Z>E", "S>W", "E>R", "L>A", "A>N", "N>D", "I>T"},
			"SWITZERLAND",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.want, func(t *testing.T) {
			require.Equal(t, tc.want, findwordfromlinkedletters.FindWord(tc.rules))
		})
	}
}
