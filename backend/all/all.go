// Package all imports all the backends
package all

import (
	_ "github.com/rclone/rclone/backend/combine"
	_ "github.com/rclone/rclone/backend/compress"
	_ "github.com/rclone/rclone/backend/crypt"
	_ "github.com/rclone/rclone/backend/drive"
	_ "github.com/rclone/rclone/backend/ftp"
	_ "github.com/rclone/rclone/backend/local"
	_ "github.com/rclone/rclone/backend/s3"
	_ "github.com/rclone/rclone/backend/sftp"
)
