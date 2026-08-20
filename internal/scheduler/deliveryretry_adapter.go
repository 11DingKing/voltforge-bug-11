package scheduler

import "errors"

func (r *DeliveryRetryRetry) ShouldRetry(err error) bool {
	return err != nil && !r.Permanent && err.Error() == ErrDeliveryRetryTransient.Error()
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
