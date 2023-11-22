package integration

import (
	"testing"

	"github.com/Pyrinpyi/pyipad/app/appmessage"
)

func setOnBlockAddedHandler(t *testing.T, harness *appHarness, handler func(notification *appmessage.BlockAddedNotificationMessage)) {
	err := harness.rpcClient.RegisterForBlockAddedNotifications(handler)
	if err != nil {
		t.Fatalf("Error from RegisterForBlockAddedNotifications: %s", err)
	}
}
