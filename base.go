// Copyright 2020 xwi88.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Build info for your project with arguments
package version

var (
	gitBranch       string
	gitTag          string
	gitCommit       = "$Format:%H$"              // sha1 from git, output of $(git rev-parse HEAD)
	gitTreeState    = "not a git tree"           // state of git tree, either "clean" or "dirty"
	buildTime       = ""                         // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	buildTimeLayout = "2006-01-02T15:04:05Z0700" // ref time.RFC3339: 2006-01-02T15:04:05Z07:00
	execTime        = ""
)
