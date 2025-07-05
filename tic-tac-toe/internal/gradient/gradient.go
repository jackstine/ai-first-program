package gradient

import (
	"math"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

// GradientType represents different gradient color schemes
type GradientType int

const (
	Rainbow GradientType = iota
	Red
	Orange
	Yellow
	Green
	Blue
	Indigo
	Violet
)

// Gradient represents a color gradient system
type Gradient struct {
	Type      GradientType
	StartTime time.Time
	Speed     float64 // Animation speed multiplier
}

// New creates a new gradient with the specified type
func New(gradientType GradientType) *Gradient {
	return &Gradient{
		Type:      gradientType,
		StartTime: time.Now(),
		Speed:     1.0,
	}
}

// GetColors returns the current gradient colors based on animation state
func (g *Gradient) GetColors(length int) []string {
	elapsed := time.Since(g.StartTime).Seconds() * g.Speed

	switch g.Type {
	case Rainbow:
		return g.getRainbowColors(length, elapsed)
	case Red:
		return g.getSolidGradient(length, elapsed, colorful.Color{R: 1.0, G: 0.0, B: 0.0})
	case Orange:
		return g.getSolidGradient(length, elapsed, colorful.Color{R: 1.0, G: 0.5, B: 0.0})
	case Yellow:
		return g.getSolidGradient(length, elapsed, colorful.Color{R: 1.0, G: 1.0, B: 0.0})
	case Green:
		return g.getSolidGradient(length, elapsed, colorful.Color{R: 0.0, G: 1.0, B: 0.0})
	case Blue:
		return g.getSolidGradient(length, elapsed, colorful.Color{R: 0.0, G: 0.0, B: 1.0})
	case Indigo:
		return g.getSolidGradient(length, elapsed, colorful.Color{R: 0.3, G: 0.0, B: 0.5})
	case Violet:
		return g.getSolidGradient(length, elapsed, colorful.Color{R: 0.5, G: 0.0, B: 1.0})
	default:
		return g.getRainbowColors(length, elapsed)
	}
}

// getRainbowColors generates rainbow gradient colors
func (g *Gradient) getRainbowColors(length int, elapsed float64) []string {
	colors := make([]string, length)
	
	for i := 0; i < length; i++ {
		// Create smooth rainbow transition with animation
		hue := (float64(i)/float64(length) + elapsed*0.1) 
		hue = math.Mod(hue, 1.0) // Keep hue in [0, 1] range
		
		color := colorful.Hsv(hue*360, 0.8, 0.9)
		colors[i] = color.Hex()
	}
	
	return colors
}

// getSolidGradient generates gradient from dark to bright for a single color
func (g *Gradient) getSolidGradient(length int, elapsed float64, baseColor colorful.Color) []string {
	colors := make([]string, length)
	
	for i := 0; i < length; i++ {
		// Create brightness variation with animation
		brightness := 0.3 + 0.7*(float64(i)/float64(length))
		
		// Add animation wave effect
		wave := math.Sin(elapsed*2 + float64(i)*0.5) * 0.2
		brightness += wave
		
		// Clamp brightness
		if brightness < 0.1 {
			brightness = 0.1
		}
		if brightness > 1.0 {
			brightness = 1.0
		}
		
		// Apply brightness to base color
		color := colorful.Color{
			R: baseColor.R * brightness,
			G: baseColor.G * brightness,
			B: baseColor.B * brightness,
		}
		
		colors[i] = color.Hex()
	}
	
	return colors
}

// ApplyToText applies gradient colors to text
func (g *Gradient) ApplyToText(text string) string {
	if text == "" {
		return text
	}
	
	colors := g.GetColors(len(text))
	
	if len(colors) == 0 {
		return text
	}
	
	result := ""
	for i, char := range text {
		colorIndex := i
		if colorIndex >= len(colors) {
			colorIndex = len(colors) - 1
		}
		
		style := lipgloss.NewStyle().Foreground(lipgloss.Color(colors[colorIndex]))
		result += style.Render(string(char))
	}
	
	return result
}

// ApplyToBlock applies gradient as background to a text block
func (g *Gradient) ApplyToBlock(text string, width, height int) string {
	colors := g.GetColors(width)
	
	style := lipgloss.NewStyle().
		Width(width).
		Height(height).
		Background(lipgloss.AdaptiveColor{
			Light: colors[0],
			Dark:  colors[len(colors)-1],
		})
	
	return style.Render(text)
}

// SetSpeed sets the animation speed
func (g *Gradient) SetSpeed(speed float64) {
	g.Speed = speed
}

// Reset resets the animation timer
func (g *Gradient) Reset() {
	g.StartTime = time.Now()
}

// GetTypeName returns the string name of the gradient type
func (g *Gradient) GetTypeName() string {
	switch g.Type {
	case Rainbow:
		return "Rainbow"
	case Red:
		return "Red"
	case Orange:
		return "Orange"
	case Yellow:
		return "Yellow"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	case Indigo:
		return "Indigo"
	case Violet:
		return "Violet"
	default:
		return "Rainbow"
	}
}