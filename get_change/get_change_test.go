package getchange_test

import (
	"fmt"
	"testing"

	getchange "github.com/padurean/algorithms-in-go/get_change"
	"github.com/stretchr/testify/require"
)

func TestGetChange(t *testing.T) {
	billsAndCoinsTypes := []float64{1, .5, .25, .10, .05, .01}

	testCases := []struct {
		payed      float64
		price      float64
		wantChange []int
	}{
		{1.00, 1.00, []int{0, 0, 0, 0, 0, 0}},
		{0.99, 0.98, []int{1, 0, 0, 0, 0, 0}},
		{1.00, 0.95, []int{0, 1, 0, 0, 0, 0}},
		{2.00, 1.90, []int{0, 0, 1, 0, 0, 0}},
		{1.00, 0.80, []int{0, 0, 2, 0, 0, 0}},
		{5.00, 4.75, []int{0, 0, 0, 1, 0, 0}},
		{0.75, 0.25, []int{0, 0, 0, 0, 1, 0}},
		{1.75, 0.75, []int{0, 0, 0, 0, 0, 1}},
		{5.00, 0.99, []int{1, 0, 0, 0, 0, 4}},
		{3.14, 1.99, []int{0, 1, 1, 0, 0, 1}},
		{3.00, 0.01, []int{4, 0, 2, 1, 1, 2}},
		{4.00, 3.14, []int{1, 0, 1, 1, 1, 0}},
		{0.45, 0.34, []int{1, 0, 1, 0, 0, 0}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%.2f-%.2f", tc.payed, tc.price), func(t *testing.T) {
			require.Equal(
				t,
				tc.wantChange,
				getchange.GetChange(tc.payed, tc.price, billsAndCoinsTypes))
		})
	}
}
