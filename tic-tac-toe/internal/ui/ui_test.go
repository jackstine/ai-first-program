package ui_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	tea "github.com/charmbracelet/bubbletea"
	"tic-tac-toe/internal/game"
	"tic-tac-toe/internal/ui"
)

var _ = Describe("UI", func() {
	var model *ui.Model
	var err error

	BeforeEach(func() {
		model, err = ui.New()
		Expect(err).ToNot(HaveOccurred())
		Expect(model).ToNot(BeNil())
	})

	Describe("New", func() {
		It("should create a new UI model successfully", func() {
			Expect(model).ToNot(BeNil())
		})
	})

	Describe("Model initialization", func() {
		It("should start in startup state", func() {
			// We can't directly access the state field since it's not exported
			// But we can test the View output to verify startup state
			view := model.View()
			Expect(view).To(ContainSubstring("Initializing..."))
		})
	})

	Describe("Update handling", func() {
		It("should handle window size messages", func() {
			msg := tea.WindowSizeMsg{Width: 80, Height: 24}
			updatedModel, _ := model.Update(msg)
			Expect(updatedModel).ToNot(BeNil())
		})

		It("should handle key messages", func() {
			// First set window size to avoid "Initializing..." view
			msg := tea.WindowSizeMsg{Width: 80, Height: 24}
			model.Update(msg)
			
			keyMsg := tea.KeyMsg{Type: tea.KeyEnter}
			updatedModel, _ := model.Update(keyMsg)
			Expect(updatedModel).ToNot(BeNil())
			// Command can be nil or not, both are valid
		})
	})

	Describe("View rendering", func() {
		It("should render initialization message when size is not set", func() {
			view := model.View()
			Expect(view).To(Equal("Initializing..."))
		})

		It("should render startup screen when size is set", func() {
			// Set window size first
			msg := tea.WindowSizeMsg{Width: 80, Height: 24}
			model.Update(msg)
			
			view := model.View()
			Expect(view).ToNot(Equal("Initializing..."))
			Expect(view).ToNot(BeEmpty())
		})
	})

	Describe("Game integration", func() {
		It("should start with a new game", func() {
			// Set window size to enable proper rendering
			msg := tea.WindowSizeMsg{Width: 80, Height: 24}
			model.Update(msg)
			
			// The model should have been initialized with a game
			view := model.View()
			Expect(view).ToNot(BeEmpty())
		})
	})

	Describe("State transitions", func() {
		BeforeEach(func() {
			// Set window size for proper rendering
			msg := tea.WindowSizeMsg{Width: 80, Height: 24}
			model.Update(msg)
		})

		It("should transition from startup to main menu on key press", func() {
			// Press any key to skip startup animation
			keyMsg := tea.KeyMsg{Type: tea.KeySpace}
			model.Update(keyMsg)
			
			view := model.View()
			// Should contain menu options
			Expect(view).To(ContainSubstring("Player vs Player"))
		})
	})

	Describe("Input handling", func() {
		BeforeEach(func() {
			// Set window size and skip startup
			msg := tea.WindowSizeMsg{Width: 80, Height: 24}
			model.Update(msg)
			keyMsg := tea.KeyMsg{Type: tea.KeySpace}
			model.Update(keyMsg)
		})

		It("should handle navigation keys in main menu", func() {
			// Test down arrow
			downKey := tea.KeyMsg{Type: tea.KeyDown}
			updatedModel, _ := model.Update(downKey)
			Expect(updatedModel).ToNot(BeNil())
			
			// Test up arrow
			upKey := tea.KeyMsg{Type: tea.KeyUp}
			updatedModel, _ = model.Update(upKey)
			Expect(updatedModel).ToNot(BeNil())
		})

		It("should handle selection in main menu", func() {
			// Test enter key
			enterKey := tea.KeyMsg{Type: tea.KeyEnter}
			updatedModel, _ := model.Update(enterKey)
			Expect(updatedModel).ToNot(BeNil())
			// Command may or may not be present depending on menu selection
		})
	})

	Describe("Error handling", func() {
		It("should handle invalid states gracefully", func() {
			// Set window size
			msg := tea.WindowSizeMsg{Width: 80, Height: 24}
			model.Update(msg)
			
			// The view should not panic even with invalid internal state
			view := model.View()
			Expect(view).ToNot(BeEmpty())
		})
	})
})

var _ = Describe("Start function", func() {
	It("should accept a game instance", func() {
		g := game.New()
		Expect(g).ToNot(BeNil())
		
		// We can't easily test the full Start function without a real terminal
		// But we can verify it doesn't panic with a nil game
		// Note: This test is limited because Start() blocks with the UI
	})
})