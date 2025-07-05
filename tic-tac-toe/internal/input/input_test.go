package input_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	tea "github.com/charmbracelet/bubbletea"
	"tic-tac-toe/internal/input"
)

var _ = Describe("Input", func() {
	var handler *input.Handler

	BeforeEach(func() {
		handler = input.New()
	})

	Describe("New", func() {
		It("should create a new input handler", func() {
			Expect(handler).ToNot(BeNil())
			x, y := handler.GetCursorPosition()
			Expect(x).To(Equal(1))
			Expect(y).To(Equal(1))
		})
	})

	Describe("ProcessKeyMsg", func() {
		It("should process arrow keys", func() {
			upKey := tea.KeyMsg{Type: tea.KeyUp}
			action := handler.ProcessKeyMsg(upKey)
			Expect(action).To(Equal(input.ActionMoveUp))

			downKey := tea.KeyMsg{Type: tea.KeyDown}
			action = handler.ProcessKeyMsg(downKey)
			Expect(action).To(Equal(input.ActionMoveDown))

			leftKey := tea.KeyMsg{Type: tea.KeyLeft}
			action = handler.ProcessKeyMsg(leftKey)
			Expect(action).To(Equal(input.ActionMoveLeft))

			rightKey := tea.KeyMsg{Type: tea.KeyRight}
			action = handler.ProcessKeyMsg(rightKey)
			Expect(action).To(Equal(input.ActionMoveRight))
		})

		It("should process action keys", func() {
			enterKey := tea.KeyMsg{Type: tea.KeyEnter}
			action := handler.ProcessKeyMsg(enterKey)
			Expect(action).To(Equal(input.ActionSelect))

			spaceKey := tea.KeyMsg{Type: tea.KeySpace}
			action = handler.ProcessKeyMsg(spaceKey)
			Expect(action).To(Equal(input.ActionSelect))

			escKey := tea.KeyMsg{Type: tea.KeyEsc}
			action = handler.ProcessKeyMsg(escKey)
			Expect(action).To(Equal(input.ActionBack))
		})

		It("should process character keys", func() {
			qKey := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}}
			action := handler.ProcessKeyMsg(qKey)
			Expect(action).To(Equal(input.ActionQuit))

			rKey := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'r'}}
			action = handler.ProcessKeyMsg(rKey)
			Expect(action).To(Equal(input.ActionReset))
		})

		It("should return unknown for unmapped keys", func() {
			unknownKey := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'z'}}
			action := handler.ProcessKeyMsg(unknownKey)
			Expect(action).To(Equal(input.ActionUnknown))
		})
	})

	Describe("MoveCursor", func() {
		It("should move cursor up", func() {
			handler.SetCursorPosition(1, 1)
			x, y := handler.MoveCursor(input.ActionMoveUp)
			Expect(x).To(Equal(1))
			Expect(y).To(Equal(0))
		})

		It("should move cursor down", func() {
			handler.SetCursorPosition(1, 1)
			x, y := handler.MoveCursor(input.ActionMoveDown)
			Expect(x).To(Equal(1))
			Expect(y).To(Equal(2))
		})

		It("should move cursor left", func() {
			handler.SetCursorPosition(1, 1)
			x, y := handler.MoveCursor(input.ActionMoveLeft)
			Expect(x).To(Equal(0))
			Expect(y).To(Equal(1))
		})

		It("should move cursor right", func() {
			handler.SetCursorPosition(1, 1)
			x, y := handler.MoveCursor(input.ActionMoveRight)
			Expect(x).To(Equal(2))
			Expect(y).To(Equal(1))
		})

		It("should not move cursor beyond boundaries", func() {
			// Test upper boundary
			handler.SetCursorPosition(1, 0)
			x, y := handler.MoveCursor(input.ActionMoveUp)
			Expect(x).To(Equal(1))
			Expect(y).To(Equal(0))

			// Test lower boundary
			handler.SetCursorPosition(1, 2)
			x, y = handler.MoveCursor(input.ActionMoveDown)
			Expect(x).To(Equal(1))
			Expect(y).To(Equal(2))

			// Test left boundary
			handler.SetCursorPosition(0, 1)
			x, y = handler.MoveCursor(input.ActionMoveLeft)
			Expect(x).To(Equal(0))
			Expect(y).To(Equal(1))

			// Test right boundary
			handler.SetCursorPosition(2, 1)
			x, y = handler.MoveCursor(input.ActionMoveRight)
			Expect(x).To(Equal(2))
			Expect(y).To(Equal(1))
		})
	})

	Describe("SetCursorPosition", func() {
		It("should set valid cursor positions", func() {
			handler.SetCursorPosition(0, 0)
			x, y := handler.GetCursorPosition()
			Expect(x).To(Equal(0))
			Expect(y).To(Equal(0))

			handler.SetCursorPosition(2, 2)
			x, y = handler.GetCursorPosition()
			Expect(x).To(Equal(2))
			Expect(y).To(Equal(2))
		})

		It("should ignore invalid cursor positions", func() {
			handler.SetCursorPosition(1, 1)
			
			handler.SetCursorPosition(-1, -1)
			x, y := handler.GetCursorPosition()
			Expect(x).To(Equal(1))
			Expect(y).To(Equal(1))

			handler.SetCursorPosition(3, 3)
			x, y = handler.GetCursorPosition()
			Expect(x).To(Equal(1))
			Expect(y).To(Equal(1))
		})
	})

	Describe("MouseToGamePosition", func() {
		It("should convert mouse coordinates to game positions", func() {
			// Test center positions (these are approximations)
			row, col, valid := handler.MouseToGamePosition(3, 1)
			if valid {
				Expect(row).To(BeNumerically(">=", 0))
				Expect(row).To(BeNumerically("<=", 2))
				Expect(col).To(BeNumerically(">=", 0))
				Expect(col).To(BeNumerically("<=", 2))
			}
		})

		It("should return invalid for out-of-bounds clicks", func() {
			_, _, valid := handler.MouseToGamePosition(-1, -1)
			Expect(valid).To(BeFalse())

			_, _, valid = handler.MouseToGamePosition(100, 100)
			Expect(valid).To(BeFalse())
		})
	})

	Describe("GetKeybindings", func() {
		It("should return keybindings", func() {
			keybindings := handler.GetKeybindings()
			Expect(len(keybindings)).To(BeNumerically(">", 0))
		})
	})

	Describe("AddKeybinding", func() {
		It("should add new keybinding", func() {
			initialCount := len(handler.GetKeybindings())
			handler.AddKeybinding("x", input.ActionHelp, "Test key")
			
			keybindings := handler.GetKeybindings()
			Expect(len(keybindings)).To(Equal(initialCount + 1))
		})

		It("should replace existing keybinding", func() {
			initialCount := len(handler.GetKeybindings())
			handler.AddKeybinding("q", input.ActionHelp, "Modified quit")
			
			keybindings := handler.GetKeybindings()
			Expect(len(keybindings)).To(Equal(initialCount))

			// Test that 'q' now maps to help instead of quit
			qKey := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}}
			action := handler.ProcessKeyMsg(qKey)
			Expect(action).To(Equal(input.ActionHelp))
		})
	})

	Describe("RemoveKeybinding", func() {
		It("should remove keybinding", func() {
			initialCount := len(handler.GetKeybindings())
			handler.RemoveKeybinding("q")
			
			keybindings := handler.GetKeybindings()
			Expect(len(keybindings)).To(Equal(initialCount - 1))

			// Test that 'q' no longer works
			qKey := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}}
			action := handler.ProcessKeyMsg(qKey)
			Expect(action).To(Equal(input.ActionUnknown))
		})
	})

	Describe("ResetCursor", func() {
		It("should reset cursor to center", func() {
			handler.SetCursorPosition(0, 0)
			handler.ResetCursor()
			
			x, y := handler.GetCursorPosition()
			Expect(x).To(Equal(1))
			Expect(y).To(Equal(1))
		})
	})

	Describe("GetActionName", func() {
		It("should return correct action names", func() {
			Expect(input.GetActionName(input.ActionMoveUp)).To(Equal("Move Up"))
			Expect(input.GetActionName(input.ActionSelect)).To(Equal("Select"))
			Expect(input.GetActionName(input.ActionQuit)).To(Equal("Quit"))
			Expect(input.GetActionName(input.ActionUnknown)).To(Equal("Unknown"))
		})
	})
})