/*
Copyright © 2020 HIDETO INAMURA <h.inamura0710@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package action

import (
	"fmt"
	"time"

	"github.com/docker/go-units"
)

// see https://github.com/helm/helm/blob/v3.2.1/internal/experimental/registry/util.go
// byteCountBinary produces a human-readable file size
func byteCountBinary(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

// shortDigest returns first 7 characters of a sha256 digest
func shortDigest(digest string) string {
	if len(digest) == 64 {
		return digest[:7]
	}
	return digest
}

// timeAgo returns a human-readable timestamp representing time that has passed
func timeAgo(t time.Time) string {
	return units.HumanDuration(time.Now().UTC().Sub(t))
}
