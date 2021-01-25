package splitstring

// import (
// 	"reflect"
// 	"testing"
// )

// func TestSplit(t *testing.T) {
// 	got := Split("abca", "b")
// 	want := []string{"a", "ca"}
// 	if !reflect.DeepEqual(got, want) {
// 		//测试用例失败
// 		t.Errorf("want:%v but got:%v", want, got)
// 	}
// }

// func Test2Split(t *testing.T) {
// 	got := Split("a:b:c:a", ":")
// 	want := []string{"a", "b", "c", "a"}
// 	if !reflect.DeepEqual(got, want) {
// 		//测试用例失败
// 		t.Errorf("want:%v but got:%v", want, got)
// 	}
// }

// func Test3Split(t *testing.T) {
// 	got := Split("abcef", "bc")
// 	want := []string{"a", "ef"}
// 	if !reflect.DeepEqual(got, want) {
// 		t.Fatalf("want:%v but got:%v", want, got)
// 	}
// }
