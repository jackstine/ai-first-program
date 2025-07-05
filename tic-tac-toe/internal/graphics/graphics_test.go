package graphics_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"tic-tac-toe/internal/gradient"
	"tic-tac-toe/internal/graphics"
)

var _ = Describe("Graphics", func() {
	var (
		g  *gradient.Gradient
		sg *graphics.StartupGraphic
	)

	BeforeEach(func() {
		g = gradient.New(gradient.Rainbow)
		sg = graphics.New(g)
	})

	Describe("New", func() {
		It("should create a new startup graphic", func() {
			Expect(sg).ToNot(BeNil())
		})
	})

	Describe("GetStartupScreen", func() {
		It("should return startup screen with title, art, and menu", func() {
			screen := sg.GetStartupScreen()
			Expect(screen).ToNot(BeEmpty())
			Expect(screen).To(ContainSubstring("MAIN MENU"))
			Expect(screen).To(ContainSubstring("Player vs Player"))
			Expect(screen).To(ContainSubstring("Player vs AI"))
		})
	})

	Describe("GetGameBoard", func() {
		It("should render empty board", func() {
			board := [3][3]string{}
			result := sg.GetGameBoard(board)
			Expect(result).ToNot(BeEmpty())
			Expect(result).To(ContainSubstring("┌─────┬─────┬─────┐"))
		})

		It("should render board with moves", func() {
			board := [3][3]string{
				{"X", " ", "O"},
				{" ", "X", " "},
				{"O", " ", "X"},
			}
			result := sg.GetGameBoard(board)
			Expect(result).ToNot(BeEmpty())
			Expect(result).To(ContainSubstring("X"))
			Expect(result).To(ContainSubstring("O"))
		})
	})

	Describe("GetKeybindingPanel", func() {
		It("should return keybinding panel", func() {
			panel := sg.GetKeybindingPanel()
			Expect(panel).ToNot(BeEmpty())
			Expect(panel).To(ContainSubstring("KEYBINDINGS"))
			Expect(panel).To(ContainSubstring("Navigate board"))
			Expect(panel).To(ContainSubstring("Place mark"))
		})
	})

	Describe("GetStatusMessage", func() {
		It("should return styled status message", func() {
			message := sg.GetStatusMessage("Player X's turn")
			Expect(message).ToNot(BeEmpty())
			Expect(message).To(ContainSubstring("Player X's turn"))
		})
	})

	Describe("SetGradient", func() {
		It("should update the gradient", func() {
			newGradient := gradient.New(gradient.Red)
			sg.SetGradient(newGradient)
			// Basic test that it doesn't crash
			screen := sg.GetStartupScreen()
			Expect(screen).ToNot(BeEmpty())
		})
	})
})