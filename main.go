// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"
	"io/fs"
	"maps"

	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/gimain"
	"goki.dev/girl/styles"
	"goki.dev/girl/units"
	"goki.dev/glide/gidom"
	"goki.dev/grr"
	"goki.dev/icons"
	"goki.dev/webki"
	"golang.org/x/net/html"
)

//go:embed content/en
var content embed.FS

func main() { gimain.Run(app) }

func app() {
	maps.Copy(gidom.ElementHandlers, elementHandlers)
	b := gi.NewBody("goki")
	pg := webki.NewPage(b).SetSource(grr.Log(fs.Sub(content, "content/en")))
	grr.Log0(pg.OpenURL(""))
	gi.DefaultTopAppBar = pg.TopAppBar
	gi.NewWindow(gi.NewScene(b)).Run().Wait()
}

var elementHandlers = map[string]gidom.Handler{
	"feature-block": func(par gi.Widget, n *html.Node) (w gi.Widget, handleChildren bool) {
		f := gi.NewFrame(par).Style(func(s *styles.Style) {
			s.Direction = styles.Column
		})
		ic := icons.Icon(gidom.GetAttr(n, "icon"))
		if !ic.IsNil() {
			gi.NewIcon(f).SetIcon(ic).Style(func(s *styles.Style) {
				s.Min.Set(units.Em(3))
			})
		}
		gi.NewLabel(f).SetType(gi.LabelHeadlineSmall).SetText(gidom.GetAttr(n, "title"))
		return f, true
	},
}
