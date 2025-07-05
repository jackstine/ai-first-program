package graphics

import (
	"strings"

	"tic-tac-toe/internal/gradient"
)

// StartupGraphic contains the 8-bit style startup screen
type StartupGraphic struct {
	gradient *gradient.Gradient
}

// New creates a new startup graphic with the specified gradient
func New(g *gradient.Gradient) *StartupGraphic {
	return &StartupGraphic{
		gradient: g,
	}
}

// GetStartupScreen returns the 8-bit style startup screen
func (sg *StartupGraphic) GetStartupScreen() string {
	art := sg.getTicTacToeArt()
	title := sg.getTitle()
	menu := sg.getMainMenu()

	return title + "\n\n" + art + "\n\n" + menu
}

// getTicTacToeArt returns 8-bit style tic-tac-toe board art
func (sg *StartupGraphic) getTicTacToeArt() string {
	// 8-bit style tic-tac-toe board with some moves
	board := []string{
		"┌─────┬─────┬─────┐",
		"│  ██ │     │ ██  │",
		"│ ████│     │████ │",
		"│  ██ │     │ ██  │",
		"├─────┼─────┼─────┤",
		"│     │ ███ │     │",
		"│     │█   █│     │",
		"│     │ ███ │     │",
		"├─────┼─────┼─────┤",
		"│ ██  │     │  ██ │",
		"│████ │     │ ████│",
		"│ ██  │     │  ██ │",
		"└─────┴─────┴─────┘",
	}

	// Apply gradient to each line
	result := make([]string, len(board))
	for i, line := range board {
		result[i] = sg.gradient.ApplyToText(line)
	}

	return strings.Join(result, "\n")
}

// getTitle returns the gradient-styled title
func (sg *StartupGraphic) getTitle() string {
	title := []string{
		"████████ ██  ██████     ████████  █████   ██████     ████████  ██████  ███████",
		"   ██    ██ ██             ██    ██   ██ ██             ██    ██    ██ ██     ",
		"   ██    ██ ██        ███  ██    ███████ ██        ███  ██    ██    ██ █████  ",
		"   ██    ██ ██    ██       ██    ██   ██ ██    ██       ██    ██    ██ ██     ",
		"   ██    ██  ██████        ██    ██   ██  ██████        ██     ██████  ███████",
	}

	result := make([]string, len(title))
	for i, line := range title {
		result[i] = sg.gradient.ApplyToText(line)
	}

	return strings.Join(result, "\n")
}

// getMainMenu returns the main menu options
func (sg *StartupGraphic) getMainMenu() string {
	menu := []string{
		"╔════════════════════════════════════════╗",
		"║                MAIN MENU               ║",
		"╠════════════════════════════════════════╣",
		"║  [1] Player vs Player                  ║",
		"║  [2] Player vs AI                      ║",
		"║  [3] Settings                          ║",
		"║  [4] Quit                              ║",
		"╚════════════════════════════════════════╝",
		"",
		"Press number key to select...",
	}

	result := make([]string, len(menu))
	for i, line := range menu {
		if i < len(menu)-2 { // Apply gradient to menu box
			result[i] = sg.gradient.ApplyToText(line)
		} else { // Keep instruction text normal
			result[i] = line
		}
	}

	return strings.Join(result, "\n")
}

// GetGameBoard returns styled game board
func (sg *StartupGraphic) GetGameBoard(board [3][3]string) string {
	// Create visual board with current game state
	boardLines := []string{
		"┌─────┬─────┬─────┐",
		"│     │     │     │",
		"│  %s  │  %s  │  %s  │",
		"│     │     │     │",
		"├─────┼─────┼─────┤",
		"│     │     │     │",
		"│  %s  │  %s  │  %s  │",
		"│     │     │     │",
		"├─────┼─────┼─────┤",
		"│     │     │     │",
		"│  %s  │  %s  │  %s  │",
		"│     │     │     │",
		"└─────┴─────┴─────┘",
	}

	// Fill in the board state
	boardStr := ""
	for i, line := range boardLines {
		if strings.Contains(line, "%s") {
			// Calculate which row we're on
			row := (i - 2) / 4
			if row >= 0 && row < 3 {
				formatted := line
				for col := 0; col < 3; col++ {
					symbol := board[row][col]
					if symbol == "" || symbol == " " {
						symbol = " "
					}
					// Replace first %s
					formatted = strings.Replace(formatted, "%s", symbol, 1)
				}
				boardStr += sg.gradient.ApplyToText(formatted) + "\n"
			}
		} else {
			boardStr += sg.gradient.ApplyToText(line) + "\n"
		}
	}

	return boardStr
}

// GetKeybindingPanel returns the keybinding help panel
func (sg *StartupGraphic) GetKeybindingPanel() string {
	keybindings := []string{
		"╔═══════════════════════╗",
		"║      KEYBINDINGS      ║",
		"╠═══════════════════════╣",
		"║ ↑↓←→  Navigate board  ║",
		"║ ENTER Place mark      ║",
		"║ SPACE Place mark      ║",
		"║ R     Reset game      ║",
		"║ Q     Quit game       ║",
		"║ S     Settings        ║",
		"║ ESC   Back to menu    ║",
		"╚═══════════════════════╝",
	}

	result := make([]string, len(keybindings))
	for i, line := range keybindings {
		result[i] = sg.gradient.ApplyToText(line)
	}

	return strings.Join(result, "\n")
}

// GetStatusMessage returns styled status message
func (sg *StartupGraphic) GetStatusMessage(message string) string {
	return sg.gradient.ApplyToText(">>> " + message + " <<<")
}

// SetGradient updates the gradient
func (sg *StartupGraphic) SetGradient(g *gradient.Gradient) {
	sg.gradient = g
}
