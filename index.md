[GoKi](https://github.com/goki/ki) is a tree package based on the `Ki` interface implemented by the `Node` struct, supporting arbitrary directed (no loops!) structural trees and standard operations thereon, for primary use in constructing the scenegraph in the GoGi GUI framework.  It can also be used for representing file system trees, web DOM trees, or any other such structural tree, which are so commonly used to represent structured information.

[GoGi](https://github.com/goki/gi) is a 2D and 3D GUI framework, built on GoKi, providing a fully native Go experience, built upon widely-used and familiar standards in the web (CSS-based styling and layout, SVG-based vector graphics) and other GUIs (e.g., the Qt Widget and model-view framework).

GoGi has feature parity (or better) compared to the industry standard Qt framework in most respects, including support for things like tooltips, fully-formatted HTML text with hypertext links, syntax-highlighted text display and full-featured editing, custom keymaps, color preferences (including light and dark mode), etc.  The model-view framework uses reflection to render complex structured GUI elements like edit dialogs, tables, lists, choosers, etc based on standard Go structured types (`struct`, slices, `map`), including a `TreeView` based on `Ki` trees, with a full-featured `FileTree` and `FileTreeView` for file system trees.

The 3D framework supports a standard scenegraph-based framework that can import standard Collada `.obj` files, and custom dynamic 3D display elements.  This 3D scenegraph can be embedded anywhere in the 2D scenegraph, and another 2D scenegraph can also be embedded within that 3D scenegraph, allowing nicely rendered, complex 2D GUI controls to be embedded within 3D scenes.

The https://github.com/goki github site has the repositories, full README and wiki docs, etc.

See the [GoGi Wiki](https://github.com/goki/gi/wiki) for more detailed documentation on the GoGi GUI.

There are extensive `examples` demo applications showing and testing the GUI elements included in the GoGi repository, and the following major applications:

* The [emergent](https://EmerSim.org) neural network simulation system ([on github](https://github.com/emer/emergent)) uses 2D and 3D GoGi GUI elements to provide interactive control and visualization of complex neural models of the brain.  These models can be run directly in Go, or via Python, and GoGi itself can be fully accessed directly through Python using the [GoPy](https://github.com/go-python/gopy) tool.

* [Gide](https://github.com/goki/gide) is a Go-based IDE (interactive development environment), which serves as the daily testing platform for GoGi, and features a strongly *emacs*-inspired design and functionality.  It is not fancy, but everything can be accessed by keyboard shortcuts so it is very efficient once learned (and emacs users should be fluent right away).  It has full knowledge of the Go language (via the [GoPi interactive parser](https://github.com/goki/pi)) and supports completion, lookup, interactive debugging via the [delve](https://github.com/go-delve/delve) debugger, etc.

* [GoPix](https://github.com/gopix) is a picture viewing and organizing app, under development but currently usable for basic tasks.  Its primary advantage is in providing full keyboard-based usability, and a generic file-based framework, that avoids any kind of lock-in into complex databases.  For example, it creates symbolic links to create folders of selected images from the full collection of files, and it takes full advantage of the `exif` metadata to update and organize picture files.  Image files can be systematically renamed by their date taken to avoid having a 100's of files all named `Image-1.jpg`.

* [Grid](https://github.com/goki/grid) is an interactive drawing program based on SVG vector-based format, like Inkscape.  Inkscape has never worked very well on Mac, and is currently almost unusable due to extremely slow GUI rendering.  Initial basic functionality is in place, with more complete features to be implemented gradually.

* [Glide](https://github.com/gok/glide) will be a lightweight internet display engine (HTML renderer and web browser), using the GoGi scenegraph as the DOM.  Because GoGi uses CSS natively and already supports HTML-based text formatting, basic functionality should be relatively easy.  Only an idea and a name at this point.

* [Grail](https://github.com/goki/grail) will be an email app, using Glide for HTML rendering of messages, featuring full keyboard-based navigation (emacs-style of course) and markup-based message formatting.  Only an idea and a name at this point.

These examples provide a decent codebase to see how to accomplish various things you might want to do.

As of now, GoGi supports full functionality across the three major desktop platforms: Mac, Linux and Windows, but it should be portable to mobile platforms and wasm with a bit of work by an interested party.  The vast majority of the system works directly on the standard Go `image.Image` interface, and it doesn't use any platform-specific widgets, so it is generally highly portable.

# Screenshots

![Screenshot of Widgets demo](/images/screenshot.png?raw=true "Screenshot of Widgets demo")

![Screenshot of Gi3D demo](/images/screenshot_gi3d.png?raw=true "Screenshot of Gi3D demo")

![Screenshot of GiEditor, Dark mode](/images/screenshot_dark.png?raw=true "Screenshot of GiEditor, Dark Mode")
