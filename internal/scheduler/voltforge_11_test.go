package scheduler

import (
	"errors"
	"fmt"
	"testing"
)

func TestVoltForge11(t *testing.T) {
	retry := &DeliveryRetryRetry{}
	wrapped := fmt.Errorf("network: %w", ErrDeliveryRetryTransient)
	retry.Record(wrapped)
	if !errors.Is(wrapped, ErrDeliveryRetryTransient) || !retry.ShouldRetry(wrapped) || retry.State() != "retrying" {
		t.Fatalf("retry state=%s attempts=%d", retry.State(), retry.Attempts)
	}
}
