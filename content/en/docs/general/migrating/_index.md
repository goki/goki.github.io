---
title: Migrating to v2
weight: 1
description: How to migrate to v2 (currently incomplete)
---

## Breaking Changes

### GoKi-wide
* Import paths have been changed from `github.com/goki/*` to `goki.dev/*` (for example, `github.com/goki/gi` changed to `goki.dev/gi`). All repositories with a changed import URL that were on version 1 are now on version 2 and have a major version URL suffix (for example, `goki.dev/gi/v2`) 
* `KiT_*` global variables have been renamed to `Type*` (for example, `KiT_Button` changed to `TypeButton`); fixing this should be a simple find and replace for `KiT_` => `Type`
* `AddNew*` functions and methods have been renamed to `New*` (for example, `AddNewButton` changed to `NewButton`); fixing this should be a simple find a replace for `AddNew` => `New`

### goki/ki
* Package `ki` has been moved into the root directory of ki, so it is now imported as just `goki.dev/ki/v2`
* Support for automatic Ki field children has been removed, so you now have to manage that yourself (see https://github.com/goki/ki/issues/17)

### goki/gi/gi
* `Defaults()` method removed on several widgets (slider, spinbox, scrollbar, etc); it is no longer needed and all calls of it can be deleted.
* `MenuButton` widget removed; use `Button` instead, as you can put a menu on any button. If you want the arrow indicator again, put `button.Indicator = icons.KeyboardArrowDown`
* `Inactive` renamed to `Disabled` and `Active` to `Enabled`. Many node flag functions (eg, `SetInactive`, `IsActive`, etc) as well as the actual enum constants must be renamed.
* `ClearAct` field on `TextField` converted to `LeadingIcon` and `TrailingIcon`; use `AddClearAction()` to easily replicate the same functionality.
* `gi.IconName` removed and replaced with `icons.Icon`; many functions and fields have changed types that must be updated; also, previous icon names may be broken or have changed to new icons, and it is strongly recommended that you use the new icon constants instead of untyped string literals. 
* Prop-based styling and configuration removed; set the `Style` field using `AddStyleFunc` for styling and directly set the configuration struct fields in your code for configuration.
* Prefs colors removed; use `gi.ColorScheme` instead.

### goki/gi/gist
* Color transformation functions like `Highlight` now take non-pointer receivers to support function chaining. 

### goki/gi/units
* `units.New*` functions have been renamed to `units.*` (for example, `units.NewPx` changed to `units.Px`); fixing this should be a simple find and replace for `units.New` (although if you use `units.NewValue`, which is unchanged, you will have to avoid changing that).
* Renamed `(units.Context).ToDotsFactor` to `(units.Context).Dots`
* Removed `units.Pct` and related things (replaced with `units.Ew`, `units.Eh`, `units.Pw`, `units.Ph`, and related things)
* Renamed `units.Context` fields `ElW`, `ElH`, `VpW`, and `VpH` to `Ew`, `Eh`, `Vw`, and `Vh` respectively.
* Setter functions for `units.Context` now also take parent size.
* `units.Value` contains a `DotsFunc` function (you should use keyed struct literals).