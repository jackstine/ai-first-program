package ui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"tic-tac-toe/internal/ai"
	"tic-tac-toe/internal/audio"
	"tic-tac-toe/internal/config"
	"tic-tac-toe/internal/game"
	"tic-tac-toe/internal/gradient"
	"tic-tac-toe/internal/graphics"
	"tic-tac-toe/internal/input"
	"tic-tac-toe/internal/persistence"
)

type GameState int

const (
	StateStartup GameState = iota
	StateMainMenu
	StateGame
	StateSettings
	StateGameOver
	StateHelp
	StateStatistics
	StateQuitConfirm
)

type Model struct {
	state            GameState
	game             *game.Game
	ai               *ai.AI
	inputHandler     *input.Handler
	gradientManager  *gradient.Gradient
	graphics         *graphics.StartupGraphic
	config           *config.Config
	persistManager   *persistence.Manager
	audioManager     *audio.Manager
	
	width            int
	height           int
	cursorPosition   [2]int
	cursorSymbols    []string
	cursorIndex      int
	cellSize         int  // Dynamic cell size based on terminal
	boardPadding     int  // Dynamic padding
	showHelp         bool
	animationTicker  *time.Ticker
	lastUpdateTime   time.Time
	
	statusMessage    string
	errorMessage     string
	showStartupAnim  bool
	startupAnimPhase int
}

func New() (*Model, error) {
	// Initialize persistence manager
	persistManager := persistence.New()
	
	// Initialize configuration
	cfg := config.New(persistManager)
	if err := cfg.Load(); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	
	// Initialize gradient with current config
	gradientManager := gradient.New(cfg.GetGradientType())
	
	// Initialize graphics with gradient
	graphicsManager := graphics.New(gradientManager)
	
	// Initialize input handler
	inputHandler := input.New()
	
	// Initialize AI with current difficulty
	aiPlayer := ai.New(cfg.GetAIDifficulty(), "O")
	
	// Initialize audio manager
	audioManager := audio.New()
	
	// Create new game
	gameInstance := game.New()
	
	// Try to load previous game state
	gameState, err := persistManager.LoadGameState()
	if err == nil && gameState != nil {
		// Note: Game state restoration would need public setters in game module
		// For now, just set the mode if it's available
		gameInstance.SetMode(game.GameMode(gameState.Mode))
	}
	
	model := &Model{
		state:            StateStartup,
		game:             gameInstance,
		ai:               aiPlayer,
		inputHandler:     inputHandler,
		gradientManager:  gradientManager,
		graphics:         graphicsManager,
		config:           cfg,
		persistManager:   persistManager,
		audioManager:     audioManager,
		cursorPosition:   [2]int{1, 0}, // Start on first menu option
		cursorSymbols:    []string{"◆", "●", "◇", "○", "▲", "▼", "◀", "▶", "★", "♦", "♠", "♣", "♥", "⬢", "⬡", "▪"},
		cursorIndex:      0,
		cellSize:         5,  // Will be updated dynamically
		boardPadding:     2,  // Will be updated dynamically
		showHelp:         false,
		lastUpdateTime:   time.Now(),
		showStartupAnim:  true,
		startupAnimPhase: 0,
	}
	
	// Start animation ticker
	model.animationTicker = time.NewTicker(time.Duration(1000.0/cfg.GetAnimationSpeed()) * time.Millisecond)
	
	// Initialize board dimensions (will be updated when window size is received)
	model.updateBoardDimensions()
	
	return model, nil
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		m.tickAnimation(),
		tea.EnableMouseCellMotion,
	)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.updateBoardDimensions()
		
	case tea.KeyMsg:
		action := m.inputHandler.ProcessKeyMsg(msg)
		cmd := m.handleKeyAction(action, msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
		
	case tea.MouseMsg:
		if mouseClick := m.inputHandler.ProcessMouseMsg(msg); mouseClick != nil {
			cmd := m.handleMouseClick(mouseClick)
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
		
	case animationTickMsg:
		m.updateAnimation()
		cmds = append(cmds, m.tickAnimation())
		
	case gameUpdateMsg:
		// Handle game state updates (AI moves, etc.)
		if msg.saveRequired {
			if err := m.saveGameState(); err != nil {
				m.errorMessage = "Failed to save game: " + err.Error()
			}
		}
	}
	
	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	if m.width == 0 || m.height == 0 {
		return "Initializing..."
	}
	
	switch m.state {
	case StateStartup:
		return m.renderStartupScreen()
	case StateMainMenu:
		return m.renderMainMenu()
	case StateGame:
		return m.renderGameScreen()
	case StateSettings:
		return m.renderSettingsScreen()
	case StateGameOver:
		return m.renderGameOverScreen()
	case StateHelp:
		return m.renderHelpScreen()
	case StateStatistics:
		return m.renderStatisticsScreen()
	case StateQuitConfirm:
		return m.renderQuitConfirmScreen()
	default:
		return "Unknown state"
	}
}

func (m *Model) renderStartupScreen() string {
	if !m.showStartupAnim {
		return m.renderMainMenu()
	}
	
	startupScreen := m.graphics.GetStartupScreen()
	
	// Center the startup screen
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Height(m.height).
		Padding(2)
	
	content := startupScreen + "\n\n"
	content += m.gradientManager.ApplyToText("Press any key to continue...")
	
	return style.Render(content)
}

func (m *Model) renderMainMenu() string {
	title := m.graphics.GetStartupScreen()
	
	menuOptions := []string{
		"🎮 Player vs Player",
		"🤖 Player vs AI",
		"⚙️  Settings",
		"📊 Statistics",
		"❓ Help",
		"🚪 Quit",
	}
	
	menu := ""
	for i, option := range menuOptions {
		if i == m.cursorPosition[1] {
			// Highlight selected option
			highlighted := "▶ " + option + " ◀"
			menu += m.gradientManager.ApplyToText(highlighted) + "\n"
		} else {
			menu += "  " + option + "\n"
		}
	}
	
	content := title + "\n\n" + menu
	
	if m.statusMessage != "" {
		content += "\n" + m.gradientManager.ApplyToText(m.statusMessage)
	}
	
	if m.errorMessage != "" {
		content += "\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render(m.errorMessage)
	}
	
	// Add keybinding help
	content += "\n\n" + lipgloss.NewStyle().Faint(true).Render("↑↓ Navigate • Enter Select • q Quit")
	
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Height(m.height).
		Padding(1)
	
	return style.Render(content)
}

func (m *Model) renderGameScreen() string {
	board := m.renderGameBoard()
	status := m.renderGameStatus()
	controls := m.renderGameControls()
	
	leftPanel := lipgloss.JoinVertical(lipgloss.Left, status, controls)
	
	if m.showHelp {
		helpPanel := m.renderHelpPanel()
		leftPanel = lipgloss.JoinVertical(lipgloss.Left, leftPanel, helpPanel)
	}
	
	// Calculate spacing for better centering
	boardWidth := m.estimateBoardWidth()
	panelWidth := 50 // Approximate width of side panel
	totalContentWidth := boardWidth + panelWidth + 4 // +4 for spacing
	
	// Center the content if terminal is wide enough
	var horizontalPadding int
	if m.width > totalContentWidth {
		horizontalPadding = (m.width - totalContentWidth) / 2
	}
	
	// Join board and panel with appropriate spacing
	spacing := "    " // 4 spaces between board and panel
	main := lipgloss.JoinHorizontal(lipgloss.Top, board, spacing, leftPanel)
	
	// Apply centering and responsive layout
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		PaddingLeft(horizontalPadding).
		PaddingTop(1).
		Width(m.width).
		Height(m.height)
	
	return style.Render(main)
}

func (m *Model) renderGameBoard() string {
	board := m.game.GetBoard()
	boardStr := ""
	
	// Create dynamic cell template based on cellSize
	emptyCellTemplate := m.createCellString("")
	xCellTemplate := m.createCellString("X")
	oCellTemplate := m.createCellString("O")
	separatorLine := m.createSeparatorLine()
	
	for row := 0; row < 3; row++ {
		// Create multi-line cells like in the diagram
		cellHeight := m.calculateCellHeight()
		
		// Create the cell content for this row
		var cellLines []string
		for lineIdx := 0; lineIdx < cellHeight; lineIdx++ {
			rowStr := ""
			for col := 0; col < 3; col++ {
				var cell string
				
				// Determine cell content based on line position
				isMiddleLine := lineIdx == cellHeight/2
				
				if isMiddleLine {
					// Middle line contains the actual content (X, O, cursor, etc.)
					if board[row][col] == game.PlayerX {
						cell = xCellTemplate
					} else if board[row][col] == game.PlayerO {
						cell = oCellTemplate
					} else {
						cell = emptyCellTemplate
					}
					
					// Highlight cursor position
					if row == m.cursorPosition[0] && col == m.cursorPosition[1] {
						currentCursor := m.cursorSymbols[m.cursorIndex]
						if board[row][col] == game.Empty {
							// Show hover effect on empty squares with current cursor symbol
							cell = m.gradientManager.ApplyToText(m.createCursorCell(currentCursor))
						} else {
							// Show selection on occupied squares
							centerPos := m.cellSize / 2
							symbol := string(cell[centerPos])
							cell = m.gradientManager.ApplyToText("▶"+m.createCellString(symbol)[1:len(cell)-1]+"◀")
						}
					} else if board[row][col] != game.Empty {
						// Apply gradient to played pieces
						cell = m.gradientManager.ApplyToText(cell)
					} else {
						// Empty squares get subtle highlighting
						style := lipgloss.NewStyle().Faint(true)
						centerPos := m.cellSize / 2
						dotCell := m.createCellString("")
						dotCell = dotCell[:centerPos] + "·" + dotCell[centerPos+1:]
						cell = style.Render(dotCell)
					}
				} else {
					// Empty lines above and below content
					cell = m.createCellString("")
				}
				
				rowStr += cell
				if col < 2 {
					rowStr += m.createSeparator()
				}
			}
			cellLines = append(cellLines, rowStr)
		}
		
		// Add all lines for this row
		for _, line := range cellLines {
			boardStr += line + "\n"
		}
		
		// Add separator line between rows
		if row < 2 {
			boardStr += separatorLine + "\n"
		}
	}
	
	// Frame the board with enhanced styling using dynamic padding
	frameStyle := lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("39")).  // Bright blue border
		Padding(m.boardPadding, m.boardPadding+1).  // Dynamic padding
		MarginTop(1).
		MarginBottom(1)
	
	return frameStyle.Render(boardStr)
}

func (m *Model) renderGameStatus() string {
	status := "GAME STATUS\n"
	status += "───────────\n"
	
	currentPlayer := m.game.GetCurrentPlayer()
	if currentPlayer == game.PlayerX {
		status += "Current: " + m.gradientManager.ApplyToText("Player X") + "\n"
	} else {
		status += "Current: " + m.gradientManager.ApplyToText("Player O") + "\n"
	}
	
	gameStatus := m.game.GetStatus()
	switch gameStatus {
	case game.StatusPlaying:
		status += "Status: In Progress\n"
	case game.StatusWon:
		winner := m.game.GetWinner()
		if winner == game.PlayerX {
			status += "Status: " + m.gradientManager.ApplyToText("X Wins!") + "\n"
		} else {
			status += "Status: " + m.gradientManager.ApplyToText("O Wins!") + "\n"
		}
	case game.StatusDraw:
		status += "Status: " + m.gradientManager.ApplyToText("Draw!") + "\n"
	}
	
	mode := m.game.GetMode()
	if mode == game.PlayerVsPlayer {
		status += "Mode: Player vs Player\n"
	} else {
		status += "Mode: Player vs AI\n"
		status += "AI: " + m.config.GetAIDifficultyName() + "\n"
	}
	
	if m.statusMessage != "" {
		status += "\n" + m.statusMessage + "\n"
	}
	
	if m.errorMessage != "" {
		status += "\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render(m.errorMessage) + "\n"
	}
	
	// Add current session score
	status += "\n" + m.renderCurrentScore() + "\n"
	
	// Add move history if there are moves
	moveHistory := m.game.GetMoveHistory()
	if len(moveHistory) > 0 {
		status += "\nMOVE HISTORY\n"
		status += "────────────\n"
		for i, move := range moveHistory {
			player := "X"
			if i%2 == 1 {
				player = "O"
			}
			status += fmt.Sprintf("%d. %s -> (%d,%d)\n", i+1, player, move.Row, move.Col)
		}
	}
	
	return status
}

func (m *Model) renderGameControls() string {
	controls := "\nCONTROLS\n"
	controls += "────────\n"
	controls += "↑↓←→ Move cursor\n"
	controls += "Enter/Space Place mark\n"
	controls += "r Reset game\n"
	controls += "t Settings\n"
	controls += "? Toggle help\n"
	controls += "g Cycle gradient\n"
	controls += "c Cycle cursor\n"
	controls += "h Show move history\n"
	controls += "esc Back to menu\n"
	controls += "q Quit\n"
	
	// Show current cursor symbol
	controls += "\nCURRENT CURSOR\n"
	controls += "──────────────\n"
	controls += fmt.Sprintf("Symbol: %s\n", m.cursorSymbols[m.cursorIndex])
	controls += fmt.Sprintf("Index: %d/%d\n", m.cursorIndex+1, len(m.cursorSymbols))
	
	return controls
}

func (m *Model) renderHelpPanel() string {
	help := "\nHELP\n"
	help += "────\n"
	help += m.inputHandler.GetKeybindingDisplay()
	return help
}

func (m *Model) renderSettingsScreen() string {
	title := m.gradientManager.ApplyToText("SETTINGS")
	
	settings := m.config.GetSettingsDisplay()
	
	content := title + "\n\n" + settings + "\n\n"
	content += "Controls:\n"
	content += "g - Cycle gradient type\n"
	content += "d - Cycle AI difficulty\n"
	content += "+/- - Adjust animation speed\n"
	content += "r - Reset to defaults\n"
	content += "esc - Back to menu\n"
	
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Height(m.height).
		Padding(2)
	
	return style.Render(content)
}

func (m *Model) renderGameOverScreen() string {
	winner := m.game.GetWinner()
	status := m.game.GetStatus()
	
	var message string
	if status == game.StatusWon {
		if winner == game.PlayerX {
			message = "🎉 PLAYER X WINS! 🎉"
		} else {
			message = "🎉 PLAYER O WINS! 🎉"
		}
	} else {
		message = "🤝 IT'S A DRAW! 🤝"
	}
	
	content := m.gradientManager.ApplyToText(message) + "\n\n"
	content += "Press 'r' to play again\n"
	content += "Press 'esc' for main menu\n"
	content += "Press 'q' to quit\n"
	
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Height(m.height).
		Padding(2)
	
	return style.Render(content)
}

func (m *Model) renderHelpScreen() string {
	content := m.gradientManager.ApplyToText("HELP & CONTROLS") + "\n\n"
	content += m.inputHandler.GetKeybindingDisplay() + "\n\n"
	content += "Press 'esc' to go back\n"
	
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Height(m.height).
		Padding(2)
	
	return style.Render(content)
}

// Animation and update functions
type animationTickMsg struct{}
type gameUpdateMsg struct {
	saveRequired bool
}

func (m *Model) tickAnimation() tea.Cmd {
	return tea.Tick(time.Duration(1000.0/m.config.GetAnimationSpeed())*time.Millisecond, func(t time.Time) tea.Msg {
		return animationTickMsg{}
	})
}

func (m *Model) updateAnimation() {
	m.lastUpdateTime = time.Now()
	
	if m.state == StateStartup {
		m.startupAnimPhase++
		if m.startupAnimPhase > 180 { // 3 seconds at 60fps
			m.showStartupAnim = false
			m.state = StateMainMenu
		}
	}
}

func (m *Model) saveGameState() error {
	return m.persistManager.SaveGameState(m.game)
}

func (m *Model) handleKeyAction(action input.KeybindingAction, keyMsg tea.KeyMsg) tea.Cmd {
	// Clear messages
	m.statusMessage = ""
	m.errorMessage = ""
	
	switch m.state {
	case StateStartup:
		if action != input.ActionUnknown {
			m.showStartupAnim = false
			m.state = StateMainMenu
		}
		
	case StateMainMenu:
		return m.handleMainMenuInput(action)
		
	case StateGame:
		return m.handleGameInput(action)
		
	case StateSettings:
		return m.handleSettingsInput(action)
		
	case StateGameOver:
		return m.handleGameOverInput(action)
		
	case StateHelp:
		if action == input.ActionBack {
			m.state = StateMainMenu
		}
		
	case StateStatistics:
		return m.handleStatisticsInput(action)
		
	case StateQuitConfirm:
		return m.handleQuitConfirmInput(action, keyMsg)
	}
	
	// Global actions
	if action == input.ActionQuit {
		m.state = StateQuitConfirm
		return nil
	}
	
	return nil
}

func (m *Model) handleMainMenuInput(action input.KeybindingAction) tea.Cmd {
	switch action {
	case input.ActionMoveUp:
		if m.cursorPosition[1] > 0 {
			m.cursorPosition[1]--
		}
	case input.ActionMoveDown:
		if m.cursorPosition[1] < 5 { // 0-5 for 6 menu options
			m.cursorPosition[1]++
		}
	case input.ActionSelect:
		return m.selectMainMenuItem()
	case input.ActionMenu1:
		m.cursorPosition[1] = 0
		return m.selectMainMenuItem()
	case input.ActionMenu2:
		m.cursorPosition[1] = 1
		return m.selectMainMenuItem()
	case input.ActionMenu3:
		m.cursorPosition[1] = 2
		return m.selectMainMenuItem()
	case input.ActionMenu4:
		m.cursorPosition[1] = 3
		return m.selectMainMenuItem()
	case input.ActionMenu5:
		m.cursorPosition[1] = 4
		return m.selectMainMenuItem()
	case input.ActionMenu6:
		m.cursorPosition[1] = 5
		return m.selectMainMenuItem()
	case input.ActionBack:
		return tea.Quit
	}
	return nil
}

func (m *Model) selectMainMenuItem() tea.Cmd {
	switch m.cursorPosition[1] {
	case 0: // Player vs Player
		m.game.Reset()
		m.game.SetMode(game.PlayerVsPlayer)
		m.state = StateGame
		m.cursorPosition = [2]int{1, 1}
		return func() tea.Msg {
			return gameUpdateMsg{saveRequired: true}
		}
	case 1: // Player vs AI
		m.game.Reset()
		m.game.SetMode(game.PlayerVsAI)
		m.state = StateGame
		m.cursorPosition = [2]int{1, 1}
		return func() tea.Msg {
			return gameUpdateMsg{saveRequired: true}
		}
	case 2: // Settings
		m.state = StateSettings
	case 3: // Statistics
		m.state = StateStatistics
	case 4: // Help
		m.state = StateHelp
	case 5: // Quit
		return tea.Quit
	}
	return nil
}

func (m *Model) handleGameInput(action input.KeybindingAction) tea.Cmd {
	switch action {
	case input.ActionMoveUp:
		x, y := m.inputHandler.MoveCursor(action)
		m.cursorPosition = [2]int{y, x} // Store as [row, col] for consistent rendering
	case input.ActionMoveDown:
		x, y := m.inputHandler.MoveCursor(action)
		m.cursorPosition = [2]int{y, x} // Store as [row, col] for consistent rendering
	case input.ActionMoveLeft:
		x, y := m.inputHandler.MoveCursor(action)
		m.cursorPosition = [2]int{y, x} // Store as [row, col] for consistent rendering
	case input.ActionMoveRight:
		x, y := m.inputHandler.MoveCursor(action)
		m.cursorPosition = [2]int{y, x} // Store as [row, col] for consistent rendering
	case input.ActionSelect:
		return m.makeMove()
	case input.ActionReset:
		m.game.Reset()
		m.cursorPosition = [2]int{1, 1}
		return func() tea.Msg {
			return gameUpdateMsg{saveRequired: true}
		}
	case input.ActionSettings:
		m.state = StateSettings
	case input.ActionHelp:
		m.showHelp = !m.showHelp
	case input.ActionToggleGradient:
		if err := m.config.NextGradientType(); err != nil {
			m.errorMessage = "Failed to change gradient: " + err.Error()
		} else {
			m.gradientManager = gradient.New(m.config.GetGradientType())
			m.graphics = graphics.New(m.gradientManager)
		}
	case input.ActionCycleCursor:
		m.cursorIndex = (m.cursorIndex + 1) % len(m.cursorSymbols)
		m.statusMessage = fmt.Sprintf("Cursor changed to: %s", m.cursorSymbols[m.cursorIndex])
	case input.ActionBack:
		m.state = StateMainMenu
		m.cursorPosition = [2]int{1, 0}
	}
	return nil
}

func (m *Model) makeMove() tea.Cmd {
	if m.game.GetStatus() != game.StatusPlaying {
		m.statusMessage = "Game is already finished!"
		return nil
	}
	
	// Use the cursor position directly (already stored as [row, col])
	row, col := m.cursorPosition[0], m.cursorPosition[1]
	if err := m.game.MakeMove(row, col); err != nil {
		m.errorMessage = err.Error()
		m.audioManager.PlaySound(audio.SoundError)
		return nil
	}
	
	// Play move sound
	m.audioManager.PlaySound(audio.SoundMove)
	
	// Check if game is over
	if m.game.GetStatus() != game.StatusPlaying {
		m.recordGameScore()
		
		// Play appropriate end game sound
		if m.game.GetStatus() == game.StatusWon {
			m.audioManager.PlaySound(audio.SoundWin)
		} else if m.game.GetStatus() == game.StatusDraw {
			m.audioManager.PlaySound(audio.SoundDraw)
		}
		
		m.state = StateGameOver
		return func() tea.Msg {
			return gameUpdateMsg{saveRequired: true}
		}
	}
	
	// Handle AI move if in AI mode
	if m.game.GetMode() == game.PlayerVsAI && m.game.GetCurrentPlayer() == game.PlayerO {
		return m.makeAIMove()
	}
	
	return func() tea.Msg {
		return gameUpdateMsg{saveRequired: true}
	}
}

func (m *Model) makeAIMove() tea.Cmd {
	aiX, aiY, err := m.ai.GetMove(m.game)
	if err != nil {
		m.errorMessage = "AI move failed: " + err.Error()
		return nil
	}
	
	if err := m.game.MakeMove(aiX, aiY); err != nil {
		m.errorMessage = "AI move error: " + err.Error()
		m.audioManager.PlaySound(audio.SoundError)
		return nil
	}
	
	// Play AI move sound
	m.audioManager.PlaySound(audio.SoundMove)
	
	// Check if game is over after AI move
	if m.game.GetStatus() != game.StatusPlaying {
		m.recordGameScore()
		
		// Play appropriate end game sound
		if m.game.GetStatus() == game.StatusWon {
			m.audioManager.PlaySound(audio.SoundWin)
		} else if m.game.GetStatus() == game.StatusDraw {
			m.audioManager.PlaySound(audio.SoundDraw)
		}
		
		m.state = StateGameOver
	}
	
	return func() tea.Msg {
		return gameUpdateMsg{saveRequired: true}
	}
}

func (m *Model) handleSettingsInput(action input.KeybindingAction) tea.Cmd {
	switch action {
	case input.ActionToggleGradient:
		if err := m.config.NextGradientType(); err != nil {
			m.errorMessage = "Failed to change gradient: " + err.Error()
		} else {
			m.gradientManager = gradient.New(m.config.GetGradientType())
			m.graphics = graphics.New(m.gradientManager)
		}
	case input.ActionCycleDifficulty:
		if err := m.config.NextAIDifficulty(); err != nil {
			m.errorMessage = "Failed to change AI difficulty: " + err.Error()
		} else {
			m.ai = ai.New(m.config.GetAIDifficulty(), "O")
			m.statusMessage = "AI difficulty changed to " + m.config.GetAIDifficultyName()
		}
	case input.ActionSpeedUp:
		if err := m.config.IncreaseAnimationSpeed(); err != nil {
			m.errorMessage = "Failed to increase speed: " + err.Error()
		} else {
			m.statusMessage = "Animation speed increased"
		}
	case input.ActionSpeedDown:
		if err := m.config.DecreaseAnimationSpeed(); err != nil {
			m.errorMessage = "Failed to decrease speed: " + err.Error()
		} else {
			m.statusMessage = "Animation speed decreased"
		}
	case input.ActionBack:
		m.state = StateMainMenu
		m.cursorPosition = [2]int{1, 0}
	}
	
	// Handle other settings actions
	switch action {
	case input.ActionUnknown:
		// Check for specific key presses for settings
		// This is a simplified approach - normally we'd handle this in ProcessKeyMsg
	}
	
	return nil
}

func (m *Model) handleGameOverInput(action input.KeybindingAction) tea.Cmd {
	switch action {
	case input.ActionReset:
		m.game.Reset()
		m.state = StateGame
		m.cursorPosition = [2]int{1, 1}
		return func() tea.Msg {
			return gameUpdateMsg{saveRequired: true}
		}
	case input.ActionBack:
		m.state = StateMainMenu
		m.cursorPosition = [2]int{1, 0}
	}
	return nil
}

func (m *Model) handleMouseClick(mouseClick *input.MouseClickMsg) tea.Cmd {
	if m.state != StateGame {
		return nil
	}
	
	// Convert mouse coordinates to game position
	row, col, valid := m.inputHandler.MouseToGamePosition(mouseClick.X, mouseClick.Y)
	if !valid {
		return nil
	}
	
	// Set cursor position and make move
	m.inputHandler.SetCursorPosition(col, row)
	m.cursorPosition = [2]int{row, col} // Store as [row, col] for consistent rendering
	
	return m.makeMove()
}

func (m *Model) renderStatisticsScreen() string {
	title := m.gradientManager.ApplyToText("GAME STATISTICS")
	
	// Load scores from persistence
	scores, err := m.persistManager.LoadScores()
	if err != nil {
		return title + "\n\nError loading statistics: " + err.Error()
	}
	
	content := title + "\n\n"
	
	// Overall stats
	content += m.gradientManager.ApplyToText("📊 OVERALL STATISTICS") + "\n"
	content += "═══════════════════════\n"
	content += fmt.Sprintf("Total Games Played: %d\n", scores.TotalGames)
	if scores.LastPlayed != "" {
		content += fmt.Sprintf("Last Played: %s\n", scores.LastPlayed)
	}
	content += "\n"
	
	// Player vs Player stats
	content += m.gradientManager.ApplyToText("🎮 PLAYER vs PLAYER") + "\n"
	content += "───────────────────\n"
	pvpStats := scores.PlayerVsPlayer
	if pvpStats.Games > 0 {
		content += fmt.Sprintf("Games: %d\n", pvpStats.Games)
		content += fmt.Sprintf("X Wins: %d (%.1f%%)\n", pvpStats.XWins, float64(pvpStats.XWins)/float64(pvpStats.Games)*100)
		content += fmt.Sprintf("O Wins: %d (%.1f%%)\n", pvpStats.OWins, float64(pvpStats.OWins)/float64(pvpStats.Games)*100)
		content += fmt.Sprintf("Draws: %d (%.1f%%)\n", pvpStats.Draws, float64(pvpStats.Draws)/float64(pvpStats.Games)*100)
	} else {
		content += "No games played yet\n"
	}
	content += "\n"
	
	// Player vs AI stats
	content += m.gradientManager.ApplyToText("🤖 PLAYER vs AI") + "\n"
	content += "─────────────────\n"
	
	aiStats := scores.PlayerVsAI
	difficulties := []struct {
		name  string
		stats persistence.DifficultyStats
	}{
		{"Easy", aiStats.Easy},
		{"Normal", aiStats.Normal},
		{"Hard", aiStats.Hard},
		{"I Never Lose", aiStats.INeverLose},
	}
	
	for _, diff := range difficulties {
		if diff.stats.Games > 0 {
			content += fmt.Sprintf("%s: %d games\n", diff.name, diff.stats.Games)
			content += fmt.Sprintf("  Player: %d wins (%.1f%%)\n", 
				diff.stats.PlayerWins, 
				float64(diff.stats.PlayerWins)/float64(diff.stats.Games)*100)
			content += fmt.Sprintf("  AI: %d wins (%.1f%%)\n", 
				diff.stats.AIWins, 
				float64(diff.stats.AIWins)/float64(diff.stats.Games)*100)
			content += fmt.Sprintf("  Draws: %d (%.1f%%)\n", 
				diff.stats.Draws, 
				float64(diff.stats.Draws)/float64(diff.stats.Games)*100)
		} else {
			content += fmt.Sprintf("%s: No games played\n", diff.name)
		}
	}
	
	content += "\n" + lipgloss.NewStyle().Faint(true).Render("r - Reset all statistics • esc - Back to menu")
	
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Height(m.height).
		Padding(2)
	
	return style.Render(content)
}

func (m *Model) handleStatisticsInput(action input.KeybindingAction) tea.Cmd {
	switch action {
	case input.ActionReset:
		// Reset all statistics
		emptyScores := &persistence.Scores{
			PlayerVsPlayer: persistence.PlayerVsPlayerStats{},
			PlayerVsAI:     persistence.PlayerVsAIStats{},
			TotalGames:     0,
			LastPlayed:     "",
		}
		if err := m.persistManager.SaveScores(emptyScores); err != nil {
			m.errorMessage = "Failed to reset statistics: " + err.Error()
		} else {
			m.statusMessage = "Statistics reset successfully!"
		}
	case input.ActionBack:
		m.state = StateMainMenu
		m.cursorPosition = [2]int{1, 0}
	}
	return nil
}

func (m *Model) recordGameScore() {
	mode := m.game.GetMode()
	winner := m.game.GetWinner()
	
	if mode == game.PlayerVsPlayer {
		// Record Player vs Player score
		if err := m.persistManager.UpdatePlayerVsPlayerScore(winner); err != nil {
			m.errorMessage = "Failed to save score: " + err.Error()
		}
	} else if mode == game.PlayerVsAI {
		// Record Player vs AI score
		difficulty := m.config.GetAIDifficulty()
		if err := m.persistManager.UpdatePlayerVsAIScore(difficulty, winner, game.PlayerO); err != nil {
			m.errorMessage = "Failed to save score: " + err.Error()
		}
	}
}

func (m *Model) renderQuitConfirmScreen() string {
	title := m.gradientManager.ApplyToText("QUIT CONFIRMATION")
	
	content := title + "\n\n"
	content += "Are you sure you want to quit?\n\n"
	content += m.gradientManager.ApplyToText("▶ Yes (Y)") + "\n"
	content += "  No (N)\n\n"
	content += lipgloss.NewStyle().Faint(true).Render("Y/Enter - Quit • N/Esc - Cancel")
	
	style := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.width).
		Height(m.height).
		Padding(2)
	
	return style.Render(content)
}

func (m *Model) handleQuitConfirmInput(action input.KeybindingAction, keyMsg tea.KeyMsg) tea.Cmd {
	// Handle Y/N key presses
	if keyMsg.Type == tea.KeyRunes && len(keyMsg.Runes) > 0 {
		key := string(keyMsg.Runes[0])
		switch key {
		case "y", "Y":
			return tea.Quit
		case "n", "N":
			m.state = StateMainMenu
			m.cursorPosition = [2]int{1, 0}
			return nil
		}
	}
	
	switch action {
	case input.ActionSelect: // Enter key - confirm quit
		return tea.Quit
	case input.ActionBack: // Esc key - cancel
		m.state = StateMainMenu
		m.cursorPosition = [2]int{1, 0}
	}
	return nil
}

// createCellString creates a cell string of the appropriate size with the symbol centered
func (m *Model) createCellString(symbol string) string {
	if m.cellSize < 3 {
		m.cellSize = 3 // Minimum size
	}
	
	// Create padding
	padding := (m.cellSize - 1) / 2
	cell := ""
	
	// Add left padding
	for i := 0; i < padding; i++ {
		cell += " "
	}
	
	// Add symbol (or space if empty)
	if symbol == "" {
		cell += " "
	} else {
		cell += symbol
	}
	
	// Add right padding
	for i := 0; i < padding; i++ {
		cell += " "
	}
	
	return cell
}

// createCursorCell creates a cursor cell with brackets around the cursor symbol
func (m *Model) createCursorCell(cursor string) string {
	// For very small cells, use minimal brackets
	if m.cellSize < 5 {
		return "[ " + cursor + " ]"
	}
	
	// Scale cursor presentation based on cell size
	var cursorContent string
	var bracketStyle string
	
	if m.cellSize >= 15 {
		// Very large cells: use double brackets and spaced cursor
		bracketStyle = "【  "
		cursorContent = " " + cursor + " "
		closeBracket := "  】"
		
		// Calculate padding
		totalBracketSpace := len(bracketStyle) + len(cursorContent) + len(closeBracket)
		availableSpace := m.cellSize - totalBracketSpace
		leftPad := availableSpace / 2
		rightPad := availableSpace - leftPad
		
		cell := ""
		for i := 0; i < leftPad; i++ {
			cell += " "
		}
		cell += bracketStyle + cursorContent + closeBracket
		for i := 0; i < rightPad; i++ {
			cell += " "
		}
		return cell
		
	} else if m.cellSize >= 10 {
		// Large cells: use enhanced brackets and spacing
		bracketStyle = "[ "
		cursorContent = " " + cursor + " "
		closeBracket := " ]"
		
		// Calculate padding
		totalBracketSpace := len(bracketStyle) + len(cursorContent) + len(closeBracket)
		availableSpace := m.cellSize - totalBracketSpace
		leftPad := availableSpace / 2
		rightPad := availableSpace - leftPad
		
		cell := ""
		for i := 0; i < leftPad; i++ {
			cell += " "
		}
		cell += bracketStyle + cursorContent + closeBracket
		for i := 0; i < rightPad; i++ {
			cell += " "
		}
		return cell
		
	} else {
		// Medium cells: standard brackets
		totalBracketSpace := 4 // "[ " + " ]"
		availableSpace := m.cellSize - totalBracketSpace
		leftPad := availableSpace / 2
		rightPad := availableSpace - leftPad
		
		cell := ""
		for i := 0; i < leftPad; i++ {
			cell += " "
		}
		cell += "[ " + cursor + " ]"
		for i := 0; i < rightPad; i++ {
			cell += " "
		}
		return cell
	}
}

// createSeparator creates the column separator
func (m *Model) createSeparator() string {
	return " │ "
}

// createSeparatorLine creates the row separator line
func (m *Model) createSeparatorLine() string {
	// Create thick, continuous horizontal lines like in the diagram
	// Each cell gets cellSize dashes, plus connecting sections
	cellDashes := ""
	for i := 0; i < m.cellSize; i++ {
		cellDashes += "─"
	}
	
	// Create the full separator with thick connecting sections
	// This creates: ───────────────┼───────────────┼───────────────
	return cellDashes + "┼" + cellDashes + "┼" + cellDashes
}

// calculateCellHeight determines how many lines tall each cell should be
func (m *Model) calculateCellHeight() int {
	// Scale cell height based on cell size to create proper proportions
	if m.cellSize >= 15 {
		return 5 // Very large cells get 5 lines (like the diagram)
	} else if m.cellSize >= 10 {
		return 4 // Large cells get 4 lines
	} else if m.cellSize >= 7 {
		return 3 // Medium cells get 3 lines
	} else {
		return 2 // Small cells get 2 lines minimum
	}
}

// estimateBoardWidth calculates the approximate width of the rendered board
func (m *Model) estimateBoardWidth() int {
	// 3 cells + 2 separators + borders + padding
	cellWidth := m.cellSize * 3
	separatorWidth := 6 // 2 * " │ " 
	borderWidth := 6   // Thick border + padding
	paddingWidth := (m.boardPadding + 1) * 2 // Left and right padding
	
	return cellWidth + separatorWidth + borderWidth + paddingWidth
}

func (m *Model) updateBoardDimensions() {
	if m.width == 0 || m.height == 0 {
		return
	}

	// Calculate available space for the board
	// Reserve space for side panels (approximately 40-50 chars)
	availableWidth := m.width - 50

	// For wide terminals (172+ cols), make board 100% larger
	if m.width >= 172 {
		// Calculate cell size to use about 60% of available width for board
		// Board width = 3 cells + 2 separators + borders + padding
		targetBoardWidth := int(float64(availableWidth) * 0.6)
		
		// Each cell: cellSize chars, separator: 3 chars (" │ "), border: ~6 chars
		// Total width = (3 * cellSize) + (2 * 3) + 6 = 3*cellSize + 12
		m.cellSize = (targetBoardWidth - 12) / 3
		
		// Ensure minimum and maximum cell sizes
		if m.cellSize < 8 {
			m.cellSize = 8
		} else if m.cellSize > 20 {
			m.cellSize = 20
		}
		
		// Large padding for bigger terminals
		m.boardPadding = 4
	} else if m.width >= 120 {
		// Medium terminals
		m.cellSize = 8
		m.boardPadding = 3
	} else if m.width >= 80 {
		// Standard terminals  
		m.cellSize = 6
		m.boardPadding = 2
	} else {
		// Small terminals
		m.cellSize = 4
		m.boardPadding = 1
	}

	// Ensure odd cell size for better centering of X/O
	if m.cellSize%2 == 0 {
		m.cellSize++
	}
}

func (m *Model) renderCurrentScore() string {
	// Load current scores from persistence
	scores, err := m.persistManager.LoadScores()
	if err != nil {
		return "SCORE: Error loading"
	}
	
	scoreText := "SESSION SCORE\n"
	scoreText += "─────────────\n"
	
	mode := m.game.GetMode()
	if mode == game.PlayerVsPlayer {
		pvpStats := scores.PlayerVsPlayer
		scoreText += fmt.Sprintf("X Wins: %d\n", pvpStats.XWins)
		scoreText += fmt.Sprintf("O Wins: %d\n", pvpStats.OWins)
		scoreText += fmt.Sprintf("Draws: %d\n", pvpStats.Draws)
		scoreText += fmt.Sprintf("Total: %d games", pvpStats.Games)
	} else {
		// Player vs AI mode
		difficulty := m.config.GetAIDifficulty()
		aiStats := scores.PlayerVsAI
		var diffStats persistence.DifficultyStats
		
		switch difficulty {
		case ai.Easy:
			diffStats = aiStats.Easy
		case ai.Normal:
			diffStats = aiStats.Normal
		case ai.Hard:
			diffStats = aiStats.Hard
		case ai.INeverLose:
			diffStats = aiStats.INeverLose
		}
		
		scoreText += fmt.Sprintf("Player: %d wins\n", diffStats.PlayerWins)
		scoreText += fmt.Sprintf("AI: %d wins\n", diffStats.AIWins)
		scoreText += fmt.Sprintf("Draws: %d\n", diffStats.Draws)
		scoreText += fmt.Sprintf("Total: %d games", diffStats.Games)
	}
	
	return scoreText
}

// Start function to begin the application
func Start(g *game.Game) error {
	m, err := New()
	if err != nil {
		return err
	}
	
	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())
	_, err = p.Run()
	return err
}