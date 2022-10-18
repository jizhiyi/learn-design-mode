package decorator

import (
	"learn/decorator/sample"
	"testing"
)

func Test_Display(t *testing.T) {
	d1 := sample.NewDisplay(sample.NewStringDisplay("hello world"))
	d2 := sample.NewDisplay(sample.NewSideBorder(d1.Inner, '#'))
	d3 := sample.NewDisplay(sample.NewFullBorder(d2.Inner))
	d1.Show()
	d2.Show()
	d3.Show()
	d4 := sample.NewDisplay(
		sample.NewSideBorder(
			sample.NewFullBorder(
				sample.NewFullBorder(
					sample.NewSideBorder(
						sample.NewFullBorder(
							sample.NewStringDisplay("hello world"),
						),
						'*',
					),
				),
			),
			'/',
		),
	)
	d4.Show()
}
