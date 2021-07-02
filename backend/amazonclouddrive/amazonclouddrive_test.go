// Test AmazonCloudDrive filesystem interface

// +build acd

package amazonclouddrive_test

import (
	"testing"

	"github.com/ThePolarNight/rclone/backend/amazonclouddrive"
	"github.com/ThePolarNight/rclone/fs"
	"github.com/ThePolarNight/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.NilObject = fs.Object((*amazonclouddrive.Object)(nil))
	fstests.RemoteName = "TestAmazonCloudDrive:"
	fstests.Run(t)
}
