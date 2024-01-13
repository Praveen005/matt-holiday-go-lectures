package hello

import "testing"

func TestSayHello(t *testing.T){
	subtests := []struct{
		items []string
		result string
	}{
		{
			result: "Hello World!",
		},
		{
			items: []string{"Pratyush"},
			result: "Hello Pratyush!",
		},
		{
			items: []string{"Prateek", "Praveen"},
			result: "Hello Prateek, Praveen!",
		},
	}

	//looping over each struct in the slice
	for _, str := range subtests{
		if s := SayHello(str.items); s != str.result{
			t.Errorf("Wanted: %s\n Got: %s\n", str.result, s)
		}
	}
	
}