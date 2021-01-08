// Copyright 2020 Pichu Chen, The PTT APP Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Encoding convertion for Big5 to UTF-8,
// Actual BBS using "BIG5-UAO", so BBS in Taiwan supports Japanese
// charset, and golang traditionalchinese supports BIG5-UAO already.

package bbs

import (
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func Utf8ToBig5(input string) []byte {
	utf8ToBig5 := traditionalchinese.Big5.NewEncoder()
	big5, _, _ := transform.Bytes(utf8ToBig5, []byte(input))
	return big5
}

func Big5ToUtf8(input []byte) string {
	big5ToUTF8 := traditionalchinese.Big5.NewDecoder()
	utf8, _, _ := transform.String(big5ToUTF8, string(input))
	return utf8
}
