package grifts

import (
	"github.com/Narwhalamongunicorns/skycasino/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
