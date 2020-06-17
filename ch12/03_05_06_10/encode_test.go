package sexpr

import (
	"encoding/json"
	"testing"
)

func TestMarshal(t *testing.T) {
	testData := []struct {
		v    interface{}
		want string
	}{
		{123.4567890, "123.456789"},
		{0.0, ""},
		{true, "t"},
		{false, ""},
		{12345679, "12345679"},
		{0, ""},
		{"string", "\"string\""},
		{"", ""},
		{complex(1, 2), "#C(1.000000 2.000000)"},
		{complex(0, 0), ""},
		{[]int{1, 2, 3}, "(1 2 3)"},
	}

	for _, d := range testData {
		data, err := Marshal(d.v)
		if err != nil {
			t.Fatalf("Marshal failed: %v\n", err)
		}
		got := string(data)
		if got != d.want {
			t.Fatalf("got=%v, want=%v\n", got, d.want)
		}
	}
}

func TestJson(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	b, err := MarshalJson(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v\n", err)
	}
	var res Movie
	err = json.Unmarshal(b, &res)
	if err != nil {
		t.Fatalf("Marshal failed: %v\n", err)
	}
	want := "Dr. Strangelove"
	if res.Title != want {
		t.Fatalf("got=%v, want=%v\n", res.Title, want)
	}
}
