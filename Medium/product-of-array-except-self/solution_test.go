package productofarrayexceptself

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{
			in:   []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
	}

	for _, c := range cases {
		got := productExceptSelf(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("productExceptSelf(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
