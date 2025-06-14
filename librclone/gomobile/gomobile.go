// Package gomobile exports shims for gomobile use
package gomobile

import (
	"github.com/rclone/rclone/librclone/librclone"

	_ "github.com/rclone/rclone/backend/all"
	_ "github.com/rclone/rclone/cmd/about"
	_ "github.com/rclone/rclone/cmd/bisync"
	_ "github.com/rclone/rclone/cmd/backend"
	_ "github.com/rclone/rclone/cmd/check"
	_ "github.com/rclone/rclone/cmd/config"
	_ "github.com/rclone/rclone/cmd/copy"
	_ "github.com/rclone/rclone/cmd/delete"
	_ "github.com/rclone/rclone/cmd/link"
	_ "github.com/rclone/rclone/cmd/ls"
	_ "github.com/rclone/rclone/cmd/move"
	_ "github.com/rclone/rclone/cmd/purge"
	_ "github.com/rclone/rclone/cmd/serve"
	_ "github.com/rclone/rclone/cmd/size"
	_ "github.com/rclone/rclone/cmd/sync"
	_ "github.com/rclone/rclone/cmd/tree"
	_ "github.com/rclone/rclone/cmd/version"
	_ "github.com/rclone/rclone/fs/operations"
	_ "github.com/rclone/rclone/fs/sync"
	_ "github.com/rclone/rclone/fs/walk"
	_ "github.com/rclone/rclone/lib/plugin"

	_ "golang.org/x/mobile/event/key" // make go.mod add this as a dependency
)

// RcloneInitialize initializes rclone as a library
func RcloneInitialize() {
	librclone.Initialize()
}

// RcloneFinalize finalizes the library
func RcloneFinalize() {
	librclone.Finalize()
}

// RcloneRPCResult is returned from RcloneRPC
//
//	Output will be returned as a serialized JSON object
//	Status is a HTTP status return (200=OK anything else fail)
type RcloneRPCResult struct {
	Output string
	Status int
}

// RcloneRPC has an interface optimised for gomobile, in particular
// the function signature is valid under gobind rules.
//
// https://pkg.go.dev/golang.org/x/mobile/cmd/gobind#hdr-Type_restrictions
func RcloneRPC(method string, input string) (result *RcloneRPCResult) { //nolint:deadcode
	output, status := librclone.RPC(method, input)
	return &RcloneRPCResult{
		Output: output,
		Status: status,
	}
}
