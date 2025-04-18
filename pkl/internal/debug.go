//===----------------------------------------------------------------------===//
// Copyright © 2024-2025 Apple Inc. and the Pkl project authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//===----------------------------------------------------------------------===//

package internal

import (
	"fmt"
	"os"
)

var DebugEnabled bool

func init() {
	if value, exists := os.LookupEnv("PKL_DEBUG"); exists && value == "1" {
		DebugEnabled = true
	}
}

// Debug writes debugging messages if PKL_DEBUG is set to 1.
func Debug(format string, a ...any) {
	if DebugEnabled {
		_, _ = os.Stderr.WriteString("[pkl-go] " + fmt.Sprintf(format, a...) + "\n")
	}
}
