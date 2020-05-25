// Copyright 2020 xwi88.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Build and get info for your project
package version

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"time"
)

// Info contains version information.
type Info struct {
	GitBranch    string `json:"git_branch"`
	GitTag       string `json:"git_tag"`
	GitCommit    string `json:"git_commit"`
	GitTreeState string `json:"git_tree_state"`
	BuildTime    string `json:"build_time"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"go_version"`
	Platform     string `json:"platform"`
	ExecTime     string `json:"exec_time"`
}

// String returns info as a human-friendly json version string.
func (info Info) String() string {
	vv, _ := json.Marshal(info)
	return fmt.Sprintf("%v", string(vv))
}

// String returns info as a human-friendly indent json version string.
func (info Info) StringWithIndent() string {
	vv, _ := json.MarshalIndent(info, "", " ")
	return fmt.Sprintf("%v", string(vv))
}

// Get get the version info
func Get() Info {
	loc, _ := time.LoadLocation("Local")
	buildTimeNow := time.Now().In(loc)
	execTime = buildTimeNow.Format(buildTimeLayout)

	if len(buildTime) == 0 || strings.Trim(buildTime, "") == "" || strings.Trim(buildTime, "''") == "" {
		buildTime = buildTimeLayout
	}
	return Info{
		GitBranch:    gitBranch,
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildTime:    buildTime,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		ExecTime:     execTime,
	}
}
