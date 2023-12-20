package main

import (
	"testing"
)

// func TestSayHello(t *testing.T) {
// 	want := "Hello, test!"
// 	got := SayHello([]string{"test"})

// 	if want != got {
// 		t.Errorf("Wanted %s, got %s", want, got)
// 	}
// }

func TestSayHello(t *testing.T) {
	subtests := []struct{
		items []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items: []string{"Matt"},
			result: "Hello, Matt!",
		},
	}

	for _, st := range subtests {
		if s := SayHello(st.items); s != st.result {
			t.Errorf("Wanted %s, got %s", st.result, s)
		}
	}

}
