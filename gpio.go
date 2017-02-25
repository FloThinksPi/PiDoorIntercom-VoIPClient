package talkiepi

import (
	"github.com/dchote/gpio"
)

func (b *Talkiepi) initGPIO() {
	// gpio status leds
	b.OnlineLED = gpio.NewOutput(OnlineLEDPin, false)
	b.ParticipantsLED = gpio.NewOutput(ParticipantsLEDPin, false)
	b.TransmitLED = gpio.NewOutput(TransmitLEDPin, false)
}

func (b *Talkiepi) LEDOn(LED gpio.Pin) {
	if b.GPIOEnabled == false {
		return
	}

	LED.High()
}

func (b *Talkiepi) LEDOff(LED gpio.Pin) {
	if b.GPIOEnabled == false {
		return
	}

	LED.Low()
}
