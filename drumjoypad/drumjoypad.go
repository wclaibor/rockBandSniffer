package drumjoypad

import (
	"github.com/artman41/vjoy"
	"github.com/dunkalunk/drumpacket"
)

const (
	minJoyID = 1
	maxJoyID = 16
)

// Button IDs
const (
	upperGreen uint = iota
	upperRed
	upperYellow
	upperBlue
	upperOrange
	lowerGreen
	lowerRed
	lowerYellow
	lowerBlue
	lowerOrange
	dpadUp
	dpadDown
	dpadLeft
	dpadRight
	buttonMenu
	buttonOptions
)

// Axis IDs
const (
	slider = vjoy.AxisX
	whammy = vjoy.AxisY
	tilt   = vjoy.AxisZ
)

// DrumJoypad is a Container for the JoypadDevice
// with utility functions baked in to retrieve the
// specific Buttons for drums and cymbals
type DrumJoypad struct {
	joypad *vjoy.Device
	rID    uint
}

// UpperGreen retrieves the Upper Green Fret
func (drumJoypad DrumJoypad) UpperGreen() *vjoy.Button {
	return drumJoypad.joypad.Button(upperGreen)
}

// UpperRed retrieves the Upper Red Fret
func (drumJoypad DrumJoypad) UpperRed() *vjoy.Button {
	return drumJoypad.joypad.Button(upperRed)
}

// UpperYellow retrieves the Upper Yellow Fret
func (drumJoypad DrumJoypad) UpperYellow() *vjoy.Button {
	return drumJoypad.joypad.Button(upperYellow)
}

// UpperBlue retrieves the Upper Blue Fret
func (drumJoypad DrumJoypad) UpperBlue() *vjoy.Button {
	return drumJoypad.joypad.Button(upperBlue)
}

// UpperOrange retrieves the Upper Orange Fret
func (drumJoypad DrumJoypad) UpperOrange() *vjoy.Button {
	return drumJoypad.joypad.Button(upperOrange)
}

// LowerGreen retrieves the Lower Green Fret
func (drumJoypad DrumJoypad) LowerGreen() *vjoy.Button {
	return drumJoypad.joypad.Button(lowerGreen)
}

// LowerRed retrieves the Lower Red Fret
func (drumJoypad DrumJoypad) LowerRed() *vjoy.Button {
	return drumJoypad.joypad.Button(lowerRed)
}

// LowerYellow retrieves the Lower Yellow Fret
func (drumJoypad DrumJoypad) LowerYellow() *vjoy.Button {
	return drumJoypad.joypad.Button(lowerYellow)
}

// LowerBlue retrieves the Lower Blue Fret
func (drumJoypad DrumJoypad) LowerBlue() *vjoy.Button {
	return drumJoypad.joypad.Button(lowerBlue)
}

// LowerOrange retrieves the Lower Orange Fret
func (drumJoypad DrumJoypad) LowerOrange() *vjoy.Button {
	return drumJoypad.joypad.Button(lowerOrange)
}

// DpadUp retrieves the Upper Dpad button
func (drumJoypad DrumJoypad) DpadUp() *vjoy.Button {
	return drumJoypad.joypad.Button(dpadUp)
}

// DpadDown retrieves the Down Dpad button
func (drumJoypad DrumJoypad) DpadDown() *vjoy.Button {
	return drumJoypad.joypad.Button(dpadDown)
}

// DpadLeft retrieves the Left Dpad button
func (drumJoypad DrumJoypad) DpadLeft() *vjoy.Button {
	return drumJoypad.joypad.Button(dpadLeft)
}

// DpadRight retrieves the Right Dpad button
func (drumJoypad DrumJoypad) DpadRight() *vjoy.Button {
	return drumJoypad.joypad.Button(dpadRight)
}

// ButtonMenu retrieves the Menu button
func (drumJoypad DrumJoypad) ButtonMenu() *vjoy.Button {
	return drumJoypad.joypad.Button(buttonMenu)
}

// ButtonOptions retrieves the Options button
func (drumJoypad DrumJoypad) ButtonOptions() *vjoy.Button {
	return drumJoypad.joypad.Button(buttonOptions)
}

// Slider retrieves the Slider Axis
func (drumJoypad DrumJoypad) Slider() *vjoy.Axis {
	return drumJoypad.joypad.Axis(slider)
}

// Tilt retrieves the Tilt Axis
func (drumJoypad DrumJoypad) Tilt() *vjoy.Axis {
	return drumJoypad.joypad.Axis(tilt)
}

// Whammy retrieves the Whammy Axis
func (drumJoypad DrumJoypad) Whammy() *vjoy.Axis {
	return drumJoypad.joypad.Axis(whammy)
}

func (drumJoypad DrumJoypad) SetUpperFretValues(frets drumpacket.Frets) {
	drumJoypad.UpperGreen().Set(frets.Green)
	drumJoypad.UpperRed().Set(frets.Red)
	drumJoypad.UpperYellow().Set(frets.Yellow)
	drumJoypad.UpperBlue().Set(frets.Blue)
	drumJoypad.UpperOrange().Set(frets.Orange)
}

func (drumJoypad DrumJoypad) SetLowerFretValues(frets drumpacket.Frets) {
	drumJoypad.LowerGreen().Set(frets.Green)
	drumJoypad.LowerRed().Set(frets.Red)
	drumJoypad.LowerYellow().Set(frets.Yellow)
	drumJoypad.LowerBlue().Set(frets.Blue)
	drumJoypad.LowerOrange().Set(frets.Orange)
}

func (drumJoypad DrumJoypad) SetDpadValues(dpad drumpacket.Dpad) {
	drumJoypad.DpadUp().Set(dpad.Up)
	drumJoypad.DpadDown().Set(dpad.Down)
	drumJoypad.DpadLeft().Set(dpad.Left)
	drumJoypad.DpadRight().Set(dpad.Right)
}

const maxFloat int = 0x7fff

func convertByte(b byte) int {
	fraction := float32(b) / float32(0xFF)
	return int(fraction * float32(maxFloat))
}

func (drumJoypad DrumJoypad) SetAxesValues(axes drumpacket.Axes) {
	sliderVal := axes.Slider / 16
	var fixedSliderVal float32
	if sliderVal == 0 {
		fixedSliderVal = 0
	} else {
		sliderFraction := float32(sliderVal) / 4
		fixedSliderVal = sliderFraction * float32(0xFF)
	}
	drumJoypad.Slider().Setc(convertByte(byte(fixedSliderVal)))
	drumJoypad.Whammy().Setc(convertByte(axes.Whammy))
	drumJoypad.Tilt().Setc(convertByte(axes.Tilt))
}

func (drumJoypad DrumJoypad) SetButtonValues(buttons drumpacket.Buttons) {
	drumJoypad.ButtonMenu().Set(buttons.Menu)
	drumJoypad.ButtonOptions().Set(buttons.Options)
}

func (drumJoypad DrumJoypad) SetValues(drumPacket drumpacket.DrumPacket) {
	drumJoypad.SetUpperFretValues(drumPacket.UpperFrets)
	drumJoypad.SetLowerFretValues(drumPacket.LowerFrets)
	drumJoypad.SetDpadValues(drumPacket.Dpad)
	drumJoypad.SetAxesValues(drumPacket.Axes)
	drumJoypad.SetButtonValues(drumPacket.Buttons)
}

// Update the vJoyDevice with the set
// Button & Axis values
func (drumJoypad DrumJoypad) Update() error {
	return drumJoypad.joypad.Update()
}

// Reset centers all Axes & resets all Buttons
func (drumJoypad DrumJoypad) Reset() {
	drumJoypad.joypad.Reset()
}

// Relinquish closes the joypad device
func (drumJoypad DrumJoypad) Relinquish() {
	drumJoypad.joypad.Relinquish()
}

// GetVirtualID returns the rID assigned by vJoy
func (drumJoypad DrumJoypad) GetVirtualID() uint {
	return drumJoypad.rID
}

// GetJoypad attempts to obtain a free Joypad
// with a Virtual Device ID between 1 and 16
func GetJoypad() (*DrumJoypad, error) {
	if !vjoy.Available() {
		return nil, ErrUnavailable
	}
	dev, rID, err := accquireJoypad()
	if err != nil {
		return nil, err
	}
	return &DrumJoypad{
		joypad: dev,
		rID:    rID,
	}, nil
}

func accquireJoypad() (dev *vjoy.Device, rID uint, err error) {
	var currentID uint = minJoyID
	dev, err = vjoy.Acquire(currentID)
	if err != nil {
		currentID++
	}
	for err == vjoy.ErrDeviceAlreadyOwned && currentID <= maxJoyID {
		dev, err = vjoy.Acquire(currentID)
		currentID++
	}
	if err != nil {
		return nil, 0, err
	}

	return dev, currentID, nil
}
