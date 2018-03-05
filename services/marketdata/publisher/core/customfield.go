package core

import (
	"github.com/quickfixgo/fix44/marketdatarequest"
	"github.com/quickfixgo/quickfix"
)

// CustomeFieldDataRequest alias of MarketDataRequest
type CustomeFieldDataRequest marketdatarequest.MarketDataRequest

//AttributedPricesField is a custom field provided by currenex
type AttributedPricesField struct{ quickfix.FIXBoolean }

const (
	// AttributedPricesTag Currenex customized tag. Must be set to 'Y' and sent with MarketDepth (264) = 0 and AggregatedBook (266) = N to receive attributed pricing.
	AttributedPricesTag quickfix.Tag = 7560
)

//Tag returns tag.AggregatedBook (266)
func (f AttributedPricesField) Tag() quickfix.Tag { return AttributedPricesTag }

//NewAttributedPrices returns a new NewAttributedPricesField initialized with val
func NewAttributedPrices(val bool) AttributedPricesField {
	return AttributedPricesField{quickfix.FIXBoolean(val)}
}

// SetNewAttributedPrices set NewAttributedPrices
func (r CustomeFieldDataRequest) SetNewAttributedPrices(v bool) {
	r.Set(NewAttributedPrices(v))
}
