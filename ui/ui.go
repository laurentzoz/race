package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/gosuri/uilive"
	"github.com/rivo/tview"

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

func Race() {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	words := fast.StaticWordList()
	p := 0
	inputField := tview.NewInputField().
		SetFieldBackgroundColor(tcell.ColorBlack)
	inputField.SetChangedFunc(func(s string) {
		if inputField.GetText() == "" {
			inputField.SetFieldBackgroundColor(tcell.ColorBlack)
			return
		}
		if words[p] == s {
			p++
			textView.Highlight(strconv.Itoa(p))
			inputField.SetText("")
			return
		}
		if strings.HasPrefix(words[p], s) {
			inputField.SetFieldBackgroundColor(tcell.ColorGreen)
			return
		}
		inputField.SetFieldBackgroundColor(tcell.ColorRed)
	})
	grid := tview.NewGrid().
		SetBorders(true).
		SetRows(15, 1)
	grid.
		AddItem(textView, 0, 0, 2, 1, 0, 0, false).
		AddItem(inputField, 1, 0, 1, 1, 0, 0, true)

	numSelections := 0
	for _, word := range words {
		fmt.Fprintf(textView, `["%d"]%s[""] `, numSelections, word)
		numSelections++
	}
	textView.Highlight("0")

	if err := app.SetRoot(grid, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}
