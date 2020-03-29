// Copyright 2015,2016,2017,2018,2019,2020 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/kasworld/htmlcolors"
)

func main() {
	filename := "color24list.html"
	fd, err := os.Create(filename)
	if err != nil {
		fmt.Printf("err in create %v\n", err)
		return
	}
	defer fd.Close()

	fmt.Fprintf(fd, `<html> <head>
	<title>Color24 list</title>
	</head><body>`)

	for _, v := range htmlcolors.Color24List {
		fmt.Fprintf(fd, `
		<div style="color:%s ; background-color:%s">%s %08x %s </div>
		`, v.Neg().ToHTMLColorString(), v.ToHTMLColorString(),
			v, uint32(v), v.ToHTMLColorString(),
		)
	}

	fmt.Fprintf(fd, `</body> </html>`)

}
