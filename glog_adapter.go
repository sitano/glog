// Go support for leveled logs, analogous to https://code.google.com/p/google-glog/
//
// Copyright 2015 Ivan Prisyazhnyy <john.koepi@gmail.com>. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Go lang "log" adapter

package glog

import (
	"log"
	"os"
)

type GLogWriter struct {
	s Severity
	depth int
}

func (ui *GLogWriter) Write(p []byte) (n int, err error) {
	Logging.printDepth(ui.s, ui.depth, string(p))
	return len(p), nil
}

func NewGLogWriter(s Severity, depth int) *GLogWriter {
	return &GLogWriter{
		s: s,
		depth: 2 + depth,
	}
}

// s is InfoLog and depth is 1 usually.
func SetLogOutputToGlog(s Severity, depth int) {
	log.SetOutput(NewGLogWriter(s, depth))
}

func SetLogToStderr() {
	log.SetOutput(os.Stderr)
}
