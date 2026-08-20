package scheduler

import "errors"

func (r *DeliveryRetryRetry) ShouldRetry(err error) bool {
	if err == nil || r.Permanent {
		return false
	}
	return errors.Is(err, ErrDeliveryRetryTransient)
}
func (r *DeliveryRetryRetry) State() string {
	if r.Permanent {
		return "permanent"
	}
	if r.Attempts > 0 {
		return "retrying"
	}
	return "pending"
}
