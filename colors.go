package main

import "github.com/fatih/color"

//  //  //

// Estructura con los colores utilizados por esta aplicación.
type appColors struct {
	clrHighlighted *color.Color
	clrBold        *color.Color
	clrError       *color.Color
}

//  //  //

var colors appColors = appColors{
	clrHighlighted: color.New(color.FgCyan),
	clrBold:        color.New(color.Bold),
	clrError:       color.New(color.FgRed),
}
