// Solved by Github YoonBaek
package main

import (
	"bufio"
	. "fmt"
	"os"
)

func solve(a, b *int) int {
	return *a + *b
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var a, b int
	Fscanf(r, "%d\n%d", &a, &b)
	Fprintln(w, solve(&a, &b))
}
