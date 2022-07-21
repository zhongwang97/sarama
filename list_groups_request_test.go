package sarama

import (
	"testing"

	"go.uber.org/goleak"
)

func TestListGroupsRequest(t *testing.T) {
	t.Cleanup(func() { goleak.IgnoreTopFunction("github.com/rcrowley/go-metrics.(*meterArbiter).tick") })
	testRequest(t, "ListGroupsRequest", &ListGroupsRequest{}, []byte{})
}
