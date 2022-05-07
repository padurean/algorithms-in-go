package maxnondupsubstr_test

import (
	"fmt"
	"testing"

	"github.com/padurean/algorithms-in-go/maxnondupsubstr"
	"github.com/stretchr/testify/require"
)

func TestMaxNonDupSubStr(t *testing.T) {
	testCases := []struct {
		name  string
		s     string
		wantS string
	}{
		{"case 0", "", ""},
		{"case 1", "  ", " "},
		{"case 2", "nndNfdfdf", "ndNf"},
		{"case 3", "界界", "界"},
		{"case 4", "世世界界", "世界"},
		{"case 5", "aabab", "ab"},
		{"case 6", "abcabc", "abc"},
		{"case 7", "界汉字界汉字", "界汉字"},
		{"case 8", "112341", "1234"},
		{"case 9", "2455lmno#%kk", "5lmno#%k"},
		{"case 10", "kkkkk", "k"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wantLen := len([]rune(tc.wantS))
			want := fmt.Sprintf(
				"max substring of\n  '%s'\nmust be\n  '%s'\nof length %d",
				tc.s, tc.wantS, wantLen)
			s, n := maxnondupsubstr.MaxNonDupSubStr(tc.s)
			require.Equal(t, tc.wantS, s, want)
			require.Equal(t, int64(wantLen), n, want)
			require.Equal(t, wantLen, len([]rune(s)), want)
		})
	}
}
