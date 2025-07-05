package input

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// MouseClickMsg represents a mouse click event
type MouseClickMsg struct {
	X, Y   int
	Button tea.MouseButton
}

// KeybindingAction represents different actions that can be triggered
type KeybindingAction int

const (
	ActionMoveUp KeybindingAction = iota
	ActionMoveDown
	ActionMoveLeft
	ActionMoveRight
	ActionSelect
	ActionBack
	ActionQuit
	ActionReset
	ActionSettings
	ActionHelp
	ActionToggleGradient
	ActionMenu1
	ActionMenu2
	ActionMenu3
	ActionMenu4
	ActionMenu5
	ActionMenu6
	ActionCycleDifficulty
	ActionSpeedUp
	ActionSpeedDown
	ActionCycleCursor
	ActionUnknown
)

// Keybinding represents a key and its action
type Keybinding struct {
	Key         string
	Action      KeybindingAction
	Description string
}

// Handler manages input processing
type Handler struct {
	keybindings []Keybinding
	cursorX     int
	cursorY     int
}

// New creates a new input handler
func New() *Handler {
	return &Handler{
		keybindings: getDefaultKeybindings(),
		cursorX:     1, // Start at center
		cursorY:     1,
	}
}

// getDefaultKeybindings returns the default key mappings
func getDefaultKeybindings() []Keybinding {
	return []Keybinding{
		// Vim-style movement (arrow keys handled in special keys section)
		{"k", ActionMoveUp, "Move cursor up (vim)"},
		{"j", ActionMoveDown, "Move cursor down (vim)"},
		{"h", ActionMoveLeft, "Move cursor left (vim)"},
		{"l", ActionMoveRight, "Move cursor right (vim)"},
		// Game actions
		{"q", ActionQuit, "Quit game"},
		{"r", ActionReset, "Reset game"},
		{"t", ActionSettings, "Open settings"}, // Changed from 's' to avoid WASD conflict
		{"?", ActionHelp, "Show help"},
		{"g", ActionToggleGradient, "Toggle gradient"},
		// Menu shortcuts
		{"1", ActionMenu1, "Menu option 1"},
		{"2", ActionMenu2, "Menu option 2"},
		{"3", ActionMenu3, "Menu option 3"},
		{"4", ActionMenu4, "Menu option 4"},
		{"5", ActionMenu5, "Menu option 5"},
		{"6", ActionMenu6, "Menu option 6"},
		{"d", ActionCycleDifficulty, "Cycle AI difficulty"},
		{"+", ActionSpeedUp, "Increase animation speed"},
		{"-", ActionSpeedDown, "Decrease animation speed"},
		{"c", ActionCycleCursor, "Cycle cursor symbol"},
		// Note: space, enter, esc handled in special keys section
	}
}

// ProcessKeyMsg processes keyboard input and returns the action
func (h *Handler) ProcessKeyMsg(msg tea.KeyMsg) KeybindingAction {
	// Handle special keys
	switch msg.Type {
	case tea.KeyUp:
		return ActionMoveUp
	case tea.KeyDown:
		return ActionMoveDown
	case tea.KeyLeft:
		return ActionMoveLeft
	case tea.KeyRight:
		return ActionMoveRight
	case tea.KeyEnter:
		return ActionSelect
	case tea.KeySpace:
		return ActionSelect
	case tea.KeyEsc:
		return ActionBack
	}
	
	// Handle character keys
	var key string
	if msg.Type == tea.KeyRunes && len(msg.Runes) > 0 {
		key = string(msg.Runes[0])
	} else {
		key = msg.String()
	}
	
	for _, binding := range h.keybindings {
		if binding.Key == key {
			return binding.Action
		}
	}
	
	return ActionUnknown
}

// ProcessMouseMsg processes mouse input and returns click position
func (h *Handler) ProcessMouseMsg(msg tea.MouseMsg) *MouseClickMsg {
	if msg.Type == tea.MouseLeft {
		return &MouseClickMsg{
			X:      msg.X,
			Y:      msg.Y,
			Button: msg.Button,
		}
	}
	
	return nil
}

// MoveCursor moves the cursor based on action
func (h *Handler) MoveCursor(action KeybindingAction) (int, int) {
	switch action {
	case ActionMoveUp:
		if h.cursorY > 0 {
			h.cursorY--
		}
	case ActionMoveDown:
		if h.cursorY < 2 {
			h.cursorY++
		}
	case ActionMoveLeft:
		if h.cursorX > 0 {
			h.cursorX--
		}
	case ActionMoveRight:
		if h.cursorX < 2 {
			h.cursorX++
		}
	}
	
	return h.cursorX, h.cursorY
}

// GetCursorPosition returns current cursor position
func (h *Handler) GetCursorPosition() (int, int) {
	return h.cursorX, h.cursorY
}

// SetCursorPosition sets the cursor position
func (h *Handler) SetCursorPosition(x, y int) {
	if x >= 0 && x <= 2 {
		h.cursorX = x
	}
	if y >= 0 && y <= 2 {
		h.cursorY = y
	}
}

// MouseToGamePosition converts mouse coordinates to game board position
// NOTE: This is a basic implementation. For full dynamic support, the UI model
// would need to pass current board dimensions to the input handler.
func (h *Handler) MouseToGamePosition(x, y int) (int, int, bool) {
	// This is a simplified mapping that works reasonably well for most board sizes
	// The actual implementation would ideally receive board dimensions from the UI
	
	// Estimate board positioning (this works for common terminal sizes)
	// Account for centering and borders
	
	// Basic offset estimation
	borderOffset := 6  // Approximate border + padding
	adjustedX := x - borderOffset
	adjustedY := y - (borderOffset / 2)
	
	col := -1
	row := -1
	
	// Adaptive cell width estimation based on click position
	// This provides reasonable accuracy for most board sizes
	// Account for larger dynamic cell sizes
	if adjustedX >= 0 {
		// Estimate cell boundaries more dynamically
		// Assume average cell size around 8-15 chars for most terminals
		if adjustedX <= 15 {
			col = 0
		} else if adjustedX <= 30 {
			col = 1
		} else if adjustedX <= 45 {
			col = 2
		}
	}
	
	// Adaptive row height estimation
	if adjustedY >= 0 {
		if adjustedY <= 4 {
			row = 0
		} else if adjustedY <= 8 {
			row = 1
		} else if adjustedY <= 12 {
			row = 2
		}
	}
	
	valid := row >= 0 && row <= 2 && col >= 0 && col <= 2
	return row, col, valid
}

// GetKeybindings returns all keybindings for display
func (h *Handler) GetKeybindings() []Keybinding {
	return h.keybindings
}

// GetKeybindingDisplay returns formatted keybinding display
func (h *Handler) GetKeybindingDisplay() string {
	display := "KEYBINDINGS:\n"
	display += "────────────\n"
	
	for _, binding := range h.keybindings {
		display += fmt.Sprintf("%-8s %s\n", binding.Key, binding.Description)
	}
	
	return display
}

// AddKeybinding adds a custom keybinding
func (h *Handler) AddKeybinding(key string, action KeybindingAction, description string) {
	// Remove existing binding for this key if it exists
	for i, binding := range h.keybindings {
		if binding.Key == key {
			h.keybindings = append(h.keybindings[:i], h.keybindings[i+1:]...)
			break
		}
	}
	
	// Add new binding
	h.keybindings = append(h.keybindings, Keybinding{
		Key:         key,
		Action:      action,
		Description: description,
	})
}

// RemoveKeybinding removes a keybinding
func (h *Handler) RemoveKeybinding(key string) {
	for i, binding := range h.keybindings {
		if binding.Key == key {
			h.keybindings = append(h.keybindings[:i], h.keybindings[i+1:]...)
			break
		}
	}
}

// ResetCursor resets cursor to center position
func (h *Handler) ResetCursor() {
	h.cursorX = 1
	h.cursorY = 1
}

// GetActionName returns the string name of an action
func GetActionName(action KeybindingAction) string {
	switch action {
	case ActionMoveUp:
		return "Move Up"
	case ActionMoveDown:
		return "Move Down"
	case ActionMoveLeft:
		return "Move Left"
	case ActionMoveRight:
		return "Move Right"
	case ActionSelect:
		return "Select"
	case ActionBack:
		return "Back"
	case ActionQuit:
		return "Quit"
	case ActionReset:
		return "Reset"
	case ActionSettings:
		return "Settings"
	case ActionHelp:
		return "Help"
	case ActionToggleGradient:
		return "Toggle Gradient"
	case ActionMenu1:
		return "Menu Option 1"
	case ActionMenu2:
		return "Menu Option 2"
	case ActionMenu3:
		return "Menu Option 3"
	case ActionMenu4:
		return "Menu Option 4"
	case ActionMenu5:
		return "Menu Option 5"
	case ActionMenu6:
		return "Menu Option 6"
	case ActionCycleDifficulty:
		return "Cycle AI Difficulty"
	case ActionSpeedUp:
		return "Increase Speed"
	case ActionSpeedDown:
		return "Decrease Speed"
	case ActionCycleCursor:
		return "Cycle Cursor"
	default:
		return "Unknown"
	}
}