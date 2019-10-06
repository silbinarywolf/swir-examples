package playback

import (
	"github.com/silbinarywolf/swir"
	"github.com/silbinarywolf/swir-examples/squaregame/internal/game/input"
)

var (
	recordPlayer *swir.Reader
)

func RecordInit(data []byte) {
	recordPlayer = swir.NewReader(input.Last, data)

	// Disable updating inputs from humans for tests
	input.DisableUpdate()
}

func RecordUpdate() bool {
	if recordPlayer == nil {
		// Not initialized, ignore
		return false
	}
	keyDowns := recordPlayer.ReadFrame()
	if keyDowns == nil {
		return true
	}
	for i, isDown := range keyDowns {
		buttonIndex := input.Button(i)
		if isDown {
			switch input.ButtonGetState(buttonIndex) {
			case input.StateNone:
				input.ButtonOverride(buttonIndex, input.StatePressed)
			case input.StatePressed:
				input.ButtonOverride(buttonIndex, input.StateHeld)
			}
		} else {
			input.ButtonOverride(buttonIndex, input.StateNone)
		}
	}
	return false
}
