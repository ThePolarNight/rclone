// Test Sftp filesystem interface

// +build !plan9

package sftp_test

import (
	"testing"

	"github.com/ThePolarNight/rclone/backend/sftp"
	"github.com/ThePolarNight/rclone/fstest"
	"github.com/ThePolarNight/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestSFTPOpenssh:",
		NilObject:  (*sftp.Object)(nil),
	})
}

func TestIntegration2(t *testing.T) {
	if *fstest.RemoteName != "" {
		t.Skip("skipping as -remote is set")
	}
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestSFTPRclone:",
		NilObject:  (*sftp.Object)(nil),
	})
}
