// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"
	"io/fs"
	"maps"

	"goki.dev/colors"
	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/gimain"
	"goki.dev/girl/styles"
	"goki.dev/girl/units"
	"goki.dev/glide/gidom"
	"goki.dev/grr"
	"goki.dev/icons"
	"goki.dev/webki"
)

//go:embed content/en
var content embed.FS

func main() { gimain.Run(app) }

func app() {
	maps.Copy(gidom.ElementHandlers, elementHandlers)
	b := gi.NewBody("goki")
	pg := webki.NewPage(b).SetSource(grr.Log(fs.Sub(content, "content/en")))
	grr.Log0(pg.OpenURL(""))
	b.AddTopBar(func(pw gi.Widget) {
		pg.TopAppBar(b.TopAppBar(pw))
	})
	b.NewWindow().Run().Wait()
}

var elementHandlers = map[string]func(ctx gidom.Context){
	"feature-block": func(ctx gidom.Context) {
		f := gidom.New[*gi.Frame](ctx).Style(func(s *styles.Style) {
			s.Direction = styles.Column
			s.Align.Items = styles.Center
			s.Grow.Set(1, 0)
		}).OnWidgetAdded(func(w gi.Widget) {
			switch w := w.(type) {
			case *gi.Label:
				w.Style(func(s *styles.Style) {
					s.Text.Align = styles.Center
				})
			}
		})
		ic := icons.Icon(gidom.GetAttr(ctx.Node(), "icon"))
		if !ic.IsNil() {
			gi.NewIcon(f).SetIcon(ic).Style(func(s *styles.Style) {
				s.Min.Set(units.Em(3))
			})
		}
		gi.NewLabel(f).SetType(gi.LabelHeadlineSmall).SetText(gidom.GetAttr(ctx.Node(), "title"))
		ctx.SetNewParent(f)
	},
	"page-info": func(ctx gidom.Context) {
		f := gidom.New[*gi.Frame](ctx).Style(func(s *styles.Style) {
			s.BackgroundColor.SetSolid(colors.Scheme.Select.Container)
			s.Border.Radius = styles.BorderRadiusMedium
			s.Padding.Set(units.Dp(8))
			s.Grow.Set(1, 0)
		})
		ctx.SetNewParent(f)
	},
}
