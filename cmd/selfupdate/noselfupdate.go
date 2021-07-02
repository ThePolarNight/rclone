// +build noselfupdate

package selfupdate

import (
	"github.com/ThePolarNight/rclone/lib/buildinfo"
)

func init() {
	buildinfo.Tags = append(buildinfo.Tags, "noselfupdate")
}
