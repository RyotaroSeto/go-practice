package basicpackage

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{4, 6, 3, 2, 4}
	s := []string{"s", "r", "e"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nam", 20},
		{"ken", 25},
		{"yon", 21},
		{"hon", 40},
	}
	fmt.Println(i, s, p)
	sort.Ints(i)
	fmt.Println(i)
	sort.Strings(s)
	fmt.Println(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(p)
	fmt.Println(i, s, p)
}
