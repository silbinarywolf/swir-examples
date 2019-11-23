package input

import (
	"io/ioutil"

	"github.com/hajimehoshi/ebiten"
	"github.com/silbinarywolf/swir"
)

type Button int32

const (
	None Button = -1
)

const (
	Right Button = 0 + iota
	Up           = 1
	Left         = 2
	Down         = 3
	Last         = 4
)

type ButtonState int32

const (
	StateNone    ButtonState = 0
	StatePressed ButtonState = 1
	StateHeld    ButtonState = 2
)

var (
	initialized  bool
	buttonStates [Last]ButtonState
	disabled     bool
	recorder     *swir.Writer
)

func buttonUpdate(button Button, isPressed bool) {
	if isPressed {
		switch buttonStates[button] {
		case StateNone:
			buttonStates[button] = StatePressed
		case StatePressed:
			buttonStates[button] = StateHeld
		}
	} else {
		buttonStates[button] = StateNone
	}
}

func ButtonInit() {
	if initialized {
		panic("Cannot call ButtonInit more than once")
	}
	recorder = swir.NewWriter(int(Last))
	initialized = true
}

// DisableUpdate will disable updating inputs so they can be mocked for tests
func DisableUpdate() {
	disabled = true
}

func ButtonUpdate() {
	if !initialized {
		panic("Must call ButtonInit before calling ButtonUpdate")
	}
	if disabled {
		return
	}

	buttonUpdate(Left, ebiten.IsKeyPressed(ebiten.KeyLeft))
	buttonUpdate(Right, ebiten.IsKeyPressed(ebiten.KeyRight))
	buttonUpdate(Up, ebiten.IsKeyPressed(ebiten.KeyUp))
	buttonUpdate(Down, ebiten.IsKeyPressed(ebiten.KeyDown))

	recorder.WriteFrame([]bool{
		Left:  buttonStates[Left] != StateNone,
		Right: buttonStates[Right] != StateNone,
		Up:    buttonStates[Up] != StateNone,
		Down:  buttonStates[Down] != StateNone,
	})

	// DEBUG: Write out a record file, this is so we can update the test!
	if ebiten.IsKeyPressed(ebiten.Key1) {
		data := recorder.Bytes()
		if err := ioutil.WriteFile("internal/test/record.swirf", data, 0644); err != nil {
			panic(err)
		}
	}
}

// ButtonState is used by the input recorder to be used with ButtonOverride
func ButtonGetState(button Button) ButtonState {
	return buttonStates[button]
}

// ButtonOverride is used to mock inputs for tests
func ButtonOverride(button Button, state ButtonState) {
	buttonStates[button] = state
}

// ButtonCheck will return true if the button is currently being held
func ButtonCheck(button Button) bool {
	return buttonStates[button] == StatePressed || buttonStates[button] == StateHeld
}

// ButtonCheckPressed will return true if the button pressed this frame and on this frame only
func ButtonCheckPressed(button Button) bool {
	return buttonStates[button] == StatePressed
}
