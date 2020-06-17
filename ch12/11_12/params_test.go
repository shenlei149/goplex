package params

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func lessThan256(v reflect.Value) error {
	i, ok := v.Interface().(int)
	if !ok {
		return fmt.Errorf("%v is not int.", v)
	}

	if i < 256 {
		return nil
	} else {
		return fmt.Errorf("%d is not less than 256.", i)
	}
}

func TestPackAndUnpack(t *testing.T) {
	type data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max" check:"lt256"`
		Exact      bool     `http:"x"`
	}
	r := httptest.NewRequest(http.MethodPost, "/?l=le149&max=2000&x=true", nil)
	var data1 data
	err := Unpack(r, &data1, map[string]Check{"lt256": lessThan256})
	if err == nil {
		t.Fatalf("lessThan256 should report an error.")
	}

	r = httptest.NewRequest(http.MethodPost, "/?l=le149&max=20&x=true", nil)
	var data2 data
	err = Unpack(r, &data2, map[string]Check{"lt256": lessThan256})
	if err != nil {
		t.Fatalf("Unpack failed: %v\n", err)
	}
	fmt.Println(data2)

	url, err := Pack(data2)
	if err != nil {
		t.Fatalf("Pack failed: %v\n", err)
	}

	want := "l=%5Ble149%5D&max=20&x=true"
	if url.RawQuery != want {
		t.Fatalf("got=%v, want=%v\n", url.RawQuery, want)
	}
}
