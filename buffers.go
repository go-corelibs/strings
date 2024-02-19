// Copyright (c) 2024  The Go-CoreLibs Authors
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

package strings

import (
	"bytes"
	"io"
)

var _ io.WriteCloser = (*ByteBuffer)(nil)

// ByteBuffer is a wrapper around bytes.Buffer which implements the io.Closer
// interface so that it can be used in io.WriteCloser contexts
type ByteBuffer struct {
	bytes.Buffer
}

// NewByteBuffer returns a new ByteBuffer instance
func NewByteBuffer() (c *ByteBuffer) {
	c = &ByteBuffer{}
	return
}

// Close fulfils the io.Closer interface and always returns nil because there
// isn't anything to actually close
func (c *ByteBuffer) Close() error {
	// Noop
	return nil
}
