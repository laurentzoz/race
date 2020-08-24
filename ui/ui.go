package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/gosuri/uilive"

	"github.com/laurentzoz/race/fast"
)

func ScanAndAdd() {
	fmt.Println("enter integers to be added")

	s := bufio.NewScanner(os.Stdin)
	w := uilive.New()
	w.Start()

	s.Split(bufio.ScanWords)
	total := 0
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println("this is not a string!")
			break
		}
		total = fast.Add(total, i)
		fmt.Fprintf(w, "total: %d\n", total)
	}
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	w.Stop()
}
