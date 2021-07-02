// Test Webdav filesystem interface
package webdav_test

import (
	"testing"

	"github.com/ThePolarNight/rclone/backend/webdav"
	"github.com/ThePolarNight/rclone/fstest"
	"github.com/ThePolarNight/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestWebdavNextcloud:",
		NilObject:  (*webdav.Object)(nil),
	})
}

// TestIntegration runs integration tests against the remote
func TestIntegration2(t *testing.T) {
	if *fstest.RemoteName != "" {
		t.Skip("skipping as -remote is set")
	}
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestWebdavOwncloud:",
		NilObject:  (*webdav.Object)(nil),
	})
}

// TestIntegration runs integration tests against the remote
func TestIntegration3(t *testing.T) {
	if *fstest.RemoteName != "" {
		t.Skip("skipping as -remote is set")
	}
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestWebdavRclone:",
		NilObject:  (*webdav.Object)(nil),
	})
}

// TestIntegration runs integration tests against the remote
func TestIntegration4(t *testing.T) {
	if *fstest.RemoteName != "" {
		t.Skip("skipping as -remote is set")
	}
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestWebdavNTLM:",
		NilObject:  (*webdav.Object)(nil),
	})
}
