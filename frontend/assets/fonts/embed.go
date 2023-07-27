package fonts

import (
	_ "embed"
)

var (
	// Iosevka
	//go:embed Iosevka/iosevka-bold.ttf
	IosevkaBold_ttf []byte
	//go:embed Iosevka/iosevka-regular.ttf
	IosevkaRegular_ttf []byte

	/// FiraCode
	//go:embed FiraMono/FiraMono-Bold.ttf
	FiraBold_ttf []byte
	//go:embed FiraMono/FiraMono-Medium.ttf
	FiraMedium_ttf []byte
	//go:embed FiraMono/FiraMono-Regular.ttf
	FiraRegular_ttf []byte

	/// Victor Mono
	//go:embed VictorMono/VictorMono-ExtraLight.ttf
	VictorExtraLight_ttf []byte
	//go:embed VictorMono/VictorMono-Thin.ttf
	VictorThin_ttf []byte
	//go:embed VictorMono/VictorMono-Light.ttf
	VictorLight_ttf []byte
	//go:embed VictorMono/VictorMono-Regular.ttf
	VictorRegular_ttf []byte
	//go:embed VictorMono/VictorMono-Medium.ttf
	VictorMedium_ttf []byte
	//go:embed VictorMono/VictorMono-SemiBold.ttf
	VictorSemiBold_ttf []byte
	//go:embed VictorMono/VictorMono-Bold.ttf
	VictorBold_ttf []byte
)
