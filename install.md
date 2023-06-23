---
layout: page
title: Install
permalink: /install/
---

1. Install the platform-specific dependencies for your platform:
    * **MacOS**:
        1. Install the xcode command-line tools if you don't already have them by running `xcode-select --install`
        2. Install the Vulkan SDK from the [Vulkan Website](https://vulkan.lunarg.com/sdk/home)
    * **Windows**:
        1. Download and install git from the [git website](https://git-scm.com/download/win)
        2. Download and install TDM-GCC from [this website](https://jmeubank.github.io/tdm-gcc/)
        3. Open Windows Command Prompt and run `cd C:\TDM-GCC-64`
        4. Then, run `mingwvars.bat` 
    * **Linux**:
        1. Install necessary dependencies by running one of the following commands:
            * If you are on Ubuntu or Debian, run `sudo apt-get install libgl1-mesa-dev xorg-dev`
            * If you are on CentOS or Fedora, run `sudo dnf install libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel libXi-devel`
        2. Install the Vulkan SDK from the [Vulkan Website](https://vulkan.lunarg.com/sdk/home)
2. Install the Go programming language (Golang) if you have not already. You can see how to do this on [the go website](https://go.dev/doc/install)
3. Clone the main GoGi repository by running `git clone https://github.com/goki/gi`

## Run the Widgets Example

4. Navigate to the widgets example by running `cd gi/examples/widgets`
5. Build the widgets example by running `go build`
6. Run the widgets example by running `./widgets` if you are on MacOS or Linux and `./widgets.exe` if you are on Windows. This should create a window with a variety of widgets, similar to the screenshot below:
![Screenshot of Widgets demo](/images/screenshot.png?raw=true "Screenshot of Widgets demo")