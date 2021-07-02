// +build linux darwin,amd64

package mount2

import (
	"testing"

	"github.com/ThePolarNight/rclone/vfs/vfstest"
)

func TestMount(t *testing.T) {
	vfstest.RunTests(t, false, mount)
}
