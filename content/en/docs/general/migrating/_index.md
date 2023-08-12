---
title: Migrating to v2
weight: 1
description: How to migrate to v2 (currently incomplete)
---

## Breaking Changes

* `Defaults()` method removed on several widgets (slider, spinbox, scrollbar, etc); it is no longer needed and all calls of it can be deleted.
* `MenuButton` widget removed; use `Button` instead, as you can put a menu on any button. If you want the arrow indicator again, put `button.Indicator = icons.KeyboardArrowDown`
* `Inactive` renamed to `Disabled` and `Active` to `Enabled`. Many node flag functions (eg, `SetInactive`, `IsActive`, etc) as well as the actual enum constants must be renamed.
* `ClearAct` field on `TextField` converted to `LeadingIcon` and `TrailingIcon`; use `AddClearAction()` to easily replicate the same functionality.
* `gi.IconName` removed and replaced with `icons.Icon`; many functions and fields have changed types that must be updated; also, previous icon names may be broken or have changed to new icons, and it is strongly recommended that you use the new icon constants instead of untyped string literals. 
* Prop-based styling and configuration removed; set the `Style` field using `AddStyleFunc` for styling and directly set the configuration struct fields in your code for configuration.
* `KiT_*` global variables have been renamed to `Type*` (for example, `KiT_Button` changed to `TypeButton`); fixing this should be a simple find and replace for `KiT_` 
* `units.New*` functions have been renamed to `units.*` (for example, `units.NewPx` changed to `units.Px`); fixing this should be a simple find and replace for `units.New` (although if you use `units.NewValue`, which is unchanged, you will have to avoid changing that).
* Prefs colors removed; use `gi.ColorScheme` instead.
* Color transformation functions like `Highlight` now take non-pointer receivers to support function chaining. 