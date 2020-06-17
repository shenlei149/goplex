package display

func Example_map() {
	type s2 struct {
		i int
	}

	type s1 struct {
		s  s2
		ss string
	}

	mapping := make(map[s1]int)
	mapping[s1{s: s2{i: 3}, ss: "l"}] = 10

	Display("map", mapping)

	mapping2 := make(map[[5]int]string)
	arr := [5]int{1, 2, 3, 4, 5}
	mapping2[arr] = "1,2,3,4,5"
	Display("map", mapping2)

	// Output:
	// Display map (map[display.s1]int):
	// map[{s:display.s2 value, ss:"l"}] = 10
	// Display map (map[[5]int]string):
	// map[[1, 2, 3, 4, 5]] = "1,2,3,4,5"
}

func Example_cycle() {
	type cycle struct {
		Value int
		Tail  *cycle
	}
	var c cycle
	c = cycle{42, &c}
	Display("c", c)

	// Output:
	// Display c (display.cycle):
	// c.Value = 42
	// (*c.Tail).Value = 42
	// (*(*c.Tail).Tail).Value = 42
	// (*(*(*c.Tail).Tail).Tail).Value = 42
}
