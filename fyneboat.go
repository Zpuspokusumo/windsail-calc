package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Boat struct {
	image        *canvas.Image
	angle        float64
	desiredAngle float64
}

func (b *Boat) UpdateAngle(delta float64) {
	b.desiredAngle += delta
	if b.desiredAngle > 360 {
		b.desiredAngle -= 360
	} else if b.desiredAngle < 0 {
		b.desiredAngle += 360
	}
}

func (b *Boat) GetImage() *canvas.Image {
	rotatedImage := canvas.NewImageFromResource(b.image.Resource)
	rotatedImage.Rotate(b.angle)
	return rotatedImage
}

func main2() {
	boatImage := canvas.NewImageFromResource(resourceLoad("boat.png")) // Replace "boat.png" with your image resource name
	boat := &Boat{image: boatImage}

	leftButton := widget.NewButton("Left", func() {
		boat.UpdateAngle(-10)
	})
	rightButton := widget.NewButton("Right", func() {
		boat.UpdateAngle(10)
	})

	canvas := canvas.NewCanvas()
	refreshFunc := func() {
		canvas.Clear()
		canvas.DrawImage(boat.GetImage(), canvas.NewPosition(100, 100))
		canvas.Refresh()
	}
	go func() {
		for {
			refreshFunc()
			boat.angle += (boat.desiredAngle - boat.angle) * 0.1
			// Adjust the 0.1 factor to control the rotation speed
		}
	}()

	app := app.New()
	w := container.NewVBox(canvas, container.NewHBox(leftButton, rightButton))
	app.Window.SetContent(w)
	app.Run()
}
