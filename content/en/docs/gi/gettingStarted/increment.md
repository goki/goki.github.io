---
title: Increment
description: Extend the Hello World example to support incrementing a number by clicking on a button.
categories: [Examples]
weight: 2
---

In the last section, you built a simple app that says, "Hello, World!" In this section, you will add to that app by creating a button that increments a label. To do that, you need to add the following code after the "Hello, World!" label:

```go
	// Add a new label to the main frame with the text "0"
	// This label will track the number of times the button has been clicked
	numLabel := gi.AddNewLabel(mfr, "numLabel", "0")
	// Make the label redrawable so that it can be updated when the
	// number of times the button has been clicked changes
	numLabel.Redrawable = true
	// Add a new button to the main frame
	button := gi.AddNewButton(mfr, "button")
	// Set the text of the button to "Increment"
	button.Text = "Increment"
	// Keep track of the number of times that the button has been clicked
	numClicked := 0
	// The OnClicked function is called every time the button is clicked
	button.OnClicked(func() {
		// In it, we increment the number of times the button has been clicked
		numClicked++
		// Then, we set the text of the number label to the number of the times
		// the button has been clicked. strconv.Itoa converts integers to strings.
		numLabel.SetText(strconv.Itoa(numClicked))
	})
```

Then, run `go build` and `./myapp` if you are on MacOS or Linux and `./myapp.exe` if you are on Windows. This should create a window similar to the last one, except with a new label and button. Each time you click the `Increment` button, the number in the new label should increase by 1. If you click the button 7 times, the app should look like this:

![Screenshot of Increment App](/screenshots/increment.png)