package to

import (
	"fmt"
	"reflect"
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

	vv := Vs(att)

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

func TestSteam_Ints(t *testing.T) {
	var ss *Value

	fmt.Println("is nil", ss.IsNil())

	fmt.Println("is str", ss.Int())

	tests := []struct {
		name    string
		v       any
		want    []int
		wantErr bool
	}{
		{"[]string", []string{"1", "2", "3"}, []int{1, 2, 3}, false},
		{"string", "llll", nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Value{
				v: tt.v,
			}
			got, err := s.IntsE()
			if (err != nil) != tt.wantErr {
				t.Errorf("Steam.Ints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Steam.Ints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_VInts(t *testing.T) {
	fmt.Println("first", V([]string{"1", "2", "3"}).VInts().First().Int())

	tests := []struct {
		name  string
		v     any
		wantI Values[int]
	}{
		{"test", []string{"1", "2", "3"}, Values[int]([]int{1, 2, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Value{
				v: tt.v,
			}
			if gotI := v.VInts(); !reflect.DeepEqual(gotI, tt.wantI) {
				t.Errorf("Value.VInts() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}
