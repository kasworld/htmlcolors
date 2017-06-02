// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package htmlcolors

import (
	gocolor "image/color"
)

func (c Color) ToRGBA() RGBA {
	return [4]uint8{
		uint8((c >> 16) & 0xff),
		uint8((c >> 8) & 0xff),
		uint8((c >> 0) & 0xff),
		0xff,
	}
}

// for image/color Color interface
func (c Color) RGBA() (uint32, uint32, uint32, uint32) {
	rtn := c.ToRGBA()
	return uint32(rtn[0]) << 8, uint32(rtn[1]) << 8, uint32(rtn[2]) << 8, uint32(rtn[3]) << 8
}

type RGBA [4]uint8

func (c RGBA) ToColor() Color {
	var rtn uint32
	for _, v := range c {
		rtn |= uint32(v)
		rtn <<= 8
	}
	return Color(rtn)
}

func (c RGBA) Neg() RGBA {
	return RGBA{255 - c[0], 255 - c[1], 255 - c[2], 255}
}
func (c RGBA) Bg() RGBA {
	g := c.ToGray()
	return RGBA{g + 128, g + 128, g + 128, 255}
	// if g <= 127 {
	// 	return RGBA{255, 255, 255, 255}
	// } else {
	// 	return RGBA{0, 0, 0, 255}
	// }
}
func (c RGBA) Rev() RGBA {
	return RGBA{128 + c[0], 128 + c[1], 128 + c[2], 255}
}

func (c RGBA) ToGray() uint8 {
	Y, _, _ := gocolor.RGBToYCbCr(c[0], c[1], c[2])
	return Y
}
