// Test HDFS filesystem interface

// +build !plan9

package hdfs_test

import (
	"testing"

	"github.com/ThePolarNight/rclone/backend/hdfs"
	"github.com/ThePolarNight/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestHdfs:",
		NilObject:  (*hdfs.Object)(nil),
	})
}
