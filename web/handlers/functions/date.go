package functions

import (
	"time"

	"github.com/rickb777/date/v2"
	"github.com/rickb777/date/v2/view"
)

func VDate(dateOrTime interface{}) view.VDate {
	switch d := dateOrTime.(type) {
	case time.Time:
		return view.NewVDate(date.NewAt(d))
	case date.Date:
		return view.NewVDate(d)
	default:
		return view.NewVDate(date.Today())
	}
}
