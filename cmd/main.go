package main

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
)

// func main() {
// 	debug := flag.Bool("debug", false, "Debug mode")

// 	flag.Parse()

// 	if *debug {
// 		fmt.Println("not handled!")
// 		return
// 	}

// 	ui.ScanAndAdd()
// }

// Demo code for the TextView primitive.

const corporate = `Leverage agile frameworks to provide a robust synopsis for high level overviews. Iterative approaches to corporate strategy foster collaborative thinking to further the overall value proposition. Organically grow the holistic world view of disruptive innovation via workplace diversity and empowerment.
Bring to the table win-win survival strategies to ensure proactive domination. At the end of the day, going forward, a new normal that has evolved from generation X is on the runway heading towards a streamlined cloud solution. User generated content in real-time will have multiple touchpoints for offshoring.
Capitalize on low hanging fruit to identify a ballpark value added activity to beta test. Override the digital divide with additional clickthroughs from DevOps. Nanotechnology immersion along the information highway will close the loop on focusing solely on the bottom line.
[yellow]Press Enter, then Tab/Backtab for word selections`

func main() {
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	inputField := tview.NewInputField().
		SetChangedFunc(func(s string) {
			textView.Highlight("2")
		})
	grid := tview.NewGrid().
		SetBorders(true).
		SetRows(15, 1)
	grid.
		AddItem(textView, 0, 0, 2, 1, 0, 0, false).
		AddItem(inputField, 1, 0, 1, 1, 0, 0, true)
	numSelections := 0
	for _, word := range strings.Split(corporate, " ") {
		fmt.Fprintf(textView, `["%d"]%s[""] `, numSelections, word)
		numSelections++
	}

	if err := app.SetRoot(grid, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}
