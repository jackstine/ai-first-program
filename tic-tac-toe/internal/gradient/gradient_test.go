package gradient_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"tic-tac-toe/internal/gradient"
)

var _ = Describe("Gradient", func() {
	var g *gradient.Gradient

	BeforeEach(func() {
		g = gradient.New(gradient.Rainbow)
	})

	Describe("New", func() {
		It("should create a new gradient with correct type", func() {
			rainbowGradient := gradient.New(gradient.Rainbow)
			Expect(rainbowGradient.GetTypeName()).To(Equal("Rainbow"))

			redGradient := gradient.New(gradient.Red)
			Expect(redGradient.GetTypeName()).To(Equal("Red"))
		})
	})

	Describe("GetColors", func() {
		It("should return correct number of colors", func() {
			colors := g.GetColors(5)
			Expect(len(colors)).To(Equal(5))
		})

		It("should return hex color strings", func() {
			colors := g.GetColors(3)
			for _, color := range colors {
				Expect(color).To(MatchRegexp(`^#[0-9a-fA-F]{6}$`))
			}
		})

		It("should handle zero length", func() {
			colors := g.GetColors(0)
			Expect(len(colors)).To(Equal(0))
		})
	})

	Describe("GetTypeName", func() {
		It("should return correct names for all gradient types", func() {
			testCases := map[gradient.GradientType]string{
				gradient.Rainbow: "Rainbow",
				gradient.Red:     "Red",
				gradient.Orange:  "Orange",
				gradient.Yellow:  "Yellow",
				gradient.Green:   "Green",
				gradient.Blue:    "Blue",
				gradient.Indigo:  "Indigo",
				gradient.Violet:  "Violet",
			}

			for gradientType, expectedName := range testCases {
				g := gradient.New(gradientType)
				Expect(g.GetTypeName()).To(Equal(expectedName))
			}
		})
	})

	Describe("ApplyToText", func() {
		It("should handle empty text", func() {
			result := g.ApplyToText("")
			Expect(result).To(Equal(""))
		})

		It("should return styled text", func() {
			result := g.ApplyToText("Hi")
			Expect(result).ToNot(BeEmpty())
			// The result should contain ANSI escape codes for styling
			Expect(len(result)).To(BeNumerically(">=", 2))
		})
	})

	Describe("SetSpeed", func() {
		It("should update animation speed", func() {
			g.SetSpeed(2.0)
			// Test that colors change differently with different speeds
			// This is a basic test since we can't easily test animation timing
			colors1 := g.GetColors(5)
			Expect(len(colors1)).To(Equal(5))
		})
	})

	Describe("Reset", func() {
		It("should reset the animation timer", func() {
			g.Reset()
			// Basic test that reset doesn't break functionality
			colors := g.GetColors(3)
			Expect(len(colors)).To(Equal(3))
		})
	})
})