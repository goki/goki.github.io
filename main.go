// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"
	"io/fs"

	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/gimain"
	"goki.dev/grr"
	"goki.dev/webki"
)

//go:embed all:content/en
var content embed.FS

func main() { gimain.Run(app) }

func app() {
	sc := gi.NewScene("goki").SetTitle("GoKi")
	pg := webki.NewPage(sc).SetSource(grr.Log(fs.Sub(content, "content/en")))
	grr.Log0(pg.OpenURL(""))
	gi.DefaultTopAppBar = pg.TopAppBar
	gi.NewWindow(sc).Run().Wait()
}
