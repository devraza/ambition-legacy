package fonts

import (
	_ "embed"
)

var (
	//go:embed iosevka-bold.ttf
	IosevkaBold_ttf []byte

	//go:embed iosevka-regular.ttf
	IosevkaRegular_ttf []byte
)
