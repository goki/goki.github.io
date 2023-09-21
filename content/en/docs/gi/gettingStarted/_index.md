+++
description = 'Set up prerequisites, install Gi, and run the Widgets example.'
title = 'Getting Started'
weight = 1
+++

## Prerequisites

On all platforms, you must download and install Go from [the Go website](https://go.dev/doc/install) if you do not already have Go 1.18+ installed.

### MacOS

1. Install the xcode command-line tools if you don't already have them by running `xcode-select --install`
2. If you don't already have the Vulkan SDK installed, install it by doing the following:
    * Run `curl -O https://sdk.lunarg.com/sdk/download/latest/mac/vulkan_sdk.dmg`
    * Run `open vulkan_sdk.dmg`
    * Double click `InstallVulkan.app`
    * Follow the installation prompts and ignore all warnings about the Vulkan Portability Enumeration extension

### Windows

1. Download and install Git for Windows from the [git website](https://git-scm.com/download/win) if you don't already have it. You should install Git Bash as part of this process and use it for development.
2. Download and install TDM-GCC from [this website](https://jmeubank.github.io/tdm-gcc/)
3. Open Windows Command Prompt and run `cd C:\TDM-GCC-64`
4. Then, run `mingwvars.bat`

### Linux

* If you are on Ubuntu or Debian, run `sudo apt-get install libgl1-mesa-dev xorg-dev`
* If you are on CentOS or Fedora, run `sudo dnf install libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel libXi-devel libXxf86vm-devel`

## Installation

Clone the Gi repository by running `git clone https://github.com/goki/gi`

## Try it out!

4. Navigate to the widgets example by running `cd gi/examples/widgets`
5. Build the widgets example by running `go build`
6. Run the widgets example by running `./widgets` if you are on MacOS or Linux and `./widgets.exe` if you are on Windows. This should create a window with a variety of widgets, similar to the screenshot below:
![Screenshot of Widgets Demo](/screenshots/widgets.png)