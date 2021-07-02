// Test Seafile filesystem interface
package seafile_test

import (
	"testing"

	"github.com/ThePolarNight/rclone/backend/seafile"
	"github.com/ThePolarNight/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestSeafile:",
		NilObject:  (*seafile.Object)(nil),
	})
}
