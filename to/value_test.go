package to

import (
	"fmt"
	"testing"
)

func TestValues_ArrayInt32(t *testing.T) {
	// tests := []struct {
	// 	name  string
	// 	v     Values
	// 	wantI []int32
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if gotI := tt.v.ArrayInt32(); !reflect.DeepEqual(gotI, tt.wantI) {
	// 			t.Errorf("Values.ArrayInt32() = %v, want %v", gotI, tt.wantI)
	// 		}
	// 	})
	// }

	var att = []int32{1, 2, 3}

	vv := ValuesF(att)

	// vv := Values[int32](att)

	fmt.Println(vv.String())

	var mtt = map[int]string{
		1: "Hello",
		2: "kk",
		3: "3",
	}

	vm := ValueM[int, string](mtt)

	fmt.Printf("vm: %+v\n", vm.Int())
}
