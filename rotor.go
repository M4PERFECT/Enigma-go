package enigmago

import "fmt"

type Rotor struct {
	Position    rune
	Rotation    rune
	wiring      [26]rune // Assuming A=0, B=1, ..., Z=25
	Notch       rune
	RingSetting rune // Optional
}

func newRotor(position, rotation, notch, ringsetting rune, wiring [26]rune) Rotor {
	return Rotor{
		Position:    position,
		Rotation:    rotation,
		wiring:      wiring,
		Notch:       notch,
		RingSetting: ringsetting,
	}
}

func (r *Rotor) encryption() {
	fmt.Print("Starting rotor encryption")
	return
}
