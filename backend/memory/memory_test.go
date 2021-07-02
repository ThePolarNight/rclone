// Test memory filesystem interface
package memory

import (
	"testing"

	"github.com/ThePolarNight/rclone/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: ":memory:",
		NilObject:  (*Object)(nil),
	})
}
