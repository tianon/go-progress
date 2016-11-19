package progress_test

import (
	"fmt"

	"go.tianon.xyz/progress"
)

func ExampleBar() {
	bar := progress.NewBar(nil)
	bar.Min = -50
	bar.Max = 50

	bar.Prefix = func(_ *progress.Bar) string {
		return "["
	}
	bar.Suffix = func(b *progress.Bar) string {
		return fmt.Sprintf("]%3.0f%%", b.Progress()*100)
	}

	bar.Phases = []string{
		" ",
		"-",
		"=",
	}

	for bar.Val = bar.Min; bar.Val <= bar.Max; bar.Val += 25 {
		fmt.Println(bar.TickString(8))
	}

	// Output:
	// [  ]  0%
	// [- ] 25%
	// [= ] 50%
	// [=-] 75%
	// [==]100%
}
