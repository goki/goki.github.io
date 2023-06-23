---
title: "Gi"
---

Gi is a 2D and 3D GUI framework, built on GoKi, providing a fully native Go experience, built upon widely-used and familiar standards in the web (CSS-based styling and layout, SVG-based vector graphics) and other GUIs (e.g., the Qt Widget and model-view framework).

Gi has feature parity (or better) compared to the industry standard Qt framework in most respects, including support for things like tooltips, fully-formatted HTML text with hypertext links, syntax-highlighted text display and full-featured editing, custom keymaps, color preferences (including light and dark mode), etc.  The model-view framework uses reflection to render complex structured GUI elements like edit dialogs, tables, lists, choosers, etc based on standard Go structured types (`struct`, slices, `map`), including a `TreeView` based on `Ki` trees, with a full-featured `FileTree` and `FileTreeView` for file system trees.

The 3D framework supports a standard scenegraph-based framework that can import standard Collada `.obj` files, and custom dynamic 3D display elements.  This 3D scenegraph can be embedded anywhere in the 2D scenegraph, and another 2D scenegraph can also be embedded within that 3D scenegraph, allowing nicely rendered, complex 2D GUI controls to be embedded within 3D scenes.