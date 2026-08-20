package scheduler

import (
	"errors"
)

var ErrDeliveryRetryTransient = errors.New("deliveryretry temporarily unavailable")
var ErrDeliveryRetryPermanent = errors.New("deliveryretry permanently rejected")

type DeliveryRetryRetry struct {
	Attempts  int
	Permanent bool
}

func (r *DeliveryRetryRetry) Record(err error) {
	if err == nil {
		r.Permanent = false
		return
	}
	if errors.Is(err, ErrDeliveryRetryTransient) {
		r.Attempts++
		return
	}
	r.Permanent = true
}
