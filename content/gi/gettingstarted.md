---
title: "Getting Started"
---

_Note: this page assumes that you have already installed GoGi by completing the steps in [Install](/gi/install)_

## Create a new Go project

1. Navigate back to your home directory by running `cd`
2. Create a new directory called myapp by running `mkdir myapp`
3. Navigate to your newly created directory by running `cd myapp`
4. Create a new Go module by running `go mod init myapp`
5. Create a new Go file by running `touch main.go`
6. Open main.go using an editor of your choice

## Make a simple app

1. Add the following code to your editor:

    ```go
    package main

    func main() {
        // Run the window event loop function as the main function
        gimain.Main(func() {
            mainrun()
        })
    }

    func mainrun() {
        // Create a window called My App Window with width 1024 and height 768
        win := gi.NewMainWindow("myapp", "My App Window", 1024, 768)
        // Get the viewport within our window
        vp := win.WinViewport2D()
        // Start a protect update on the viewport
        updt := vp.UpdateStart()
        // Create a standard frame within the window and make it the main widget
        mfr := win.SetMainFrame()

        // Add a label to the main frame with the text "Hello, World!"
        label := gi.AddNewLabel(mfr, "label", "Hello, World!")

        // End the protected update on the viewport without a signal.
        // Update signals cause things to be redrawn, which is unnecessary at the start
        // because it is already drawing everything new.
        vp.UpdateEndNoSig(updt)
        // Start the event loop that keeps the window rendering.
        // This is a blocking call, and it will not return until
        // the user quits the app or gi.Quit() is called
        win.StartEventLoop()
    }
    ```

2. Update your dependencies by running `go mod tidy`
3. Build the code by running `go build`
4. Run the app by running `./myapp` if you are on MacOS or Linux and `./myapp.exe` if you are on Windows.

