---
layout: home
title: GoKi
permalink: /
---

GoKi is an open-source project that provides a set of frameworks for constructing cross-platform GUIs and other useful tools in pure Go using full-strength tree structures. The name GoKi is derived from the word tree in Japanese (木), which is pronounced Ki (き). 


## Ki
The core package of GoKi is [Ki](/ki), which provides trees using the `Ki` interface implemented by the `Node` struct. This supports arbitrarily directed structural trees and standard operations on them.

## Gi
[Gi](/gi) is a pure Go 2D and 3D GUI framework, built on [Ki](/ki) and widely used standards in the web, like CSS for styling and SVG for vector graphics.

## Gide
[Gide](/gide) is an IDE and IDE framework built using [Gi](/gi) and [Pi](/pi). It has standard editor features like syntax highlighting, completion, and version control built-in.

## Pi
[Pi](/pi) is an interactive parsing library that uses a simple and robust form of lexing and parsing based on top-down recursive descent.

## vGPU

[vGPU](/vgpu) is a Vulkan-based framework for both Graphics and Compute Engine use of GPU hardware in Go.

## GoSL

[GoSL](/gosl) implements Go as a shader language for GPU compute shaders by converting Go code to HLSL, and then using the glslc compiler to compile into an `.spv` SPIR-V file that can be loaded into a vulkan compute shader.

## Grid
Grid is a Go SVG vector drawing program, built using Gi and based on Inkscape.

## GoPix 

[GoPix](/gopix) is a Go picture management app.

## Mat32

[Mat32](/mat32) is a float32 based vector and matrix package for 2D and 3D graphics that uses a value-based design.

**Coming Soon**

___

## Glide

[Glide](/glide) will be a lightweight internet display engine (HTML renderer and web browser), built using Gi.

## Grail

[Grail](/grail) will be an email app, using Glide for HTML rendering of messages, featuring full keyboard-based navigation and markup-based message formatting.
