package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/container/layout"
	"fyne.io/fyne/v2/container/tab"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/container/widget"
)

var isoPath string
var ramSize string

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("QEMU GUI")

	basicTabContent := container.NewVBox(
		widget.NewLabel("Basic Configuration"),
		widget.NewButton("Select ISO", func() {
			showFileDialog(myWindow)
		}),
		widget.NewEntryWithData(&isoPath, "ISO Path").SetPlaceHolder("Enter path to ISO"),
		widget.NewLabel("RAM Size (in MB)"),
		widget.NewEntryWithData(&ramSize, "RAM Size").SetPlaceHolder("Enter RAM size"),
		widget.NewButton("Start QEMU", func() {
			startQEMU()
		}),
	)

	advancedTabContent := container.NewVBox(
		widget.NewLabel("Advanced Configuration"),
	)

	tabs := tab.NewContainer(
		tab.NewItem("Basic", basicTabContent),
		tab.NewItem("Advanced", advancedTabContent),
	)

	myWindow.SetContent(container.New(layout.NewMaxLayout(), tabs))
	myWindow.ShowAndRun()
}

func showFileDialog(window fyne.Window) {
	dialog.ShowFileOpen(func(uris []fyne.URI) {
		if len(uris) > 0 {
			isoPath = uris[0].Path()
			fmt.Println("Selected ISO:", isoPath)
		}
	}, window)
}

func startQEMU() {
	if isoPath == "" {
		dialog.ShowError(errors.New("Please select an ISO file"), myWindow)
		return
	}

	if ramSize == "" {
		dialog.ShowError(errors.New("Please enter RAM size"), myWindow)
		return
	}

	cmd := exec.Command("qemu-system-x86_64", "-hda", isoPath, "-m", ramSize)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
