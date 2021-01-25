package splitstring

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	//测试组
	// type testCase struct {
	// 	str  string
	// 	sep  string
	// 	want []string
	// }
	// testGroup := []testCase{
	// 	testCase{"abca", "b", []string{"a", "ca"}},
	// 	testCase{"a:b:c:a", ":", []string{"a", "b", "c", "a"}},
	// 	testCase{"abcef", "bc", []string{"a", "ef"}},
	// 	testCase{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	// }
	// for _, tc := range testGroup {
	// 	got := Split(tc.str, tc.sep)
	// 	if !reflect.DeepEqual(got, tc.want) {
	// 		t.Fatalf("want:%#v got:%#v\n",tc.want,got)
	// 	}
	// }

	//子测试
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		//go test -run=TestSplit/case_3 单独测试某个用例
		"case_1": testCase{"abca", "b", []string{"a", "ca"}},
		"case_2": testCase{"a:b:c:a", ":", []string{"a", "b", "c", "a"}},
		"case_3": testCase{"abcef", "bc", []string{"a", "ef"}},
		"case_4": testCase{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}

//go test -bench=Split -benchmem
//BenchmarkSplit 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
