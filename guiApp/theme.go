package guiApp

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type dejaVuSansMonoTheme struct{}

//var _ fyne.Theme = (*myTheme)(nil)

func (m dejaVuSansMonoTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m dejaVuSansMonoTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m dejaVuSansMonoTheme) Font(style fyne.TextStyle) fyne.Resource {
	return resourceAssetsDejaVuSansMonoTtf
}

func (m dejaVuSansMonoTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
