// Test filesystem interface
package premiumizeme_test

import (
	"testing"

	"github.com/ThePolarNight/rclone/rclone/backend/premiumizeme"
	"github.com/ThePolarNight/rclone/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestPremiumizeMe:",
		NilObject:  (*premiumizeme.Object)(nil),
	})
}
