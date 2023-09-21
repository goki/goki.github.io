+++
linkTitle = 'About'
title = 'About GoKi'

[menu]
  [menu.main]
    weight = 10
+++

{{% blocks/cover title="About GoKi" image_anchor="bottom" height="auto" %}}

A framework for trees (Ki in Japanese) in Go, including the GoGi 2D & 3D GUI framework.
{.mt-5}

{{% /blocks/cover %}}

{{% blocks/lead %}}

GoKi is an open-source project that provides a set of frameworks for constructing cross-platform GUIs and other useful tools in pure Go using full-strength tree structures. The name GoKi is derived from the word tree in Japanese (木), which is pronounced Ki (き). 

{{% /blocks/lead %}}

{{% blocks/section %}}

### The core package of GoKi is Ki, which provides trees using the `Ki` interface implemented by the `Node` struct. This supports arbitrarily directed structural trees and standard operations on them.
{.text-center}

{{% /blocks/section %}}

{{% blocks/section %}}

### Gi uses the Ki tree infrastructure to implement a scenegraph-based GUI framework in full native idiomatic Go, with minimal OS-specific backend interfaces based originally on the Shiny drivers, now using go-gl/glfw and vulkan-based vgpu, and supporting MacOS, Linux, and Windows.
{.text-center}

{{% /blocks/section %}}