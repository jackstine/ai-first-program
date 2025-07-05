package persistence

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"tic-tac-toe/internal/ai"
	"tic-tac-toe/internal/game"
	"tic-tac-toe/internal/gradient"
)

const (
	saveDir        = ".tic-tac-toe"
	gameStateFile  = "gamestate.json"
	settingsFile   = "settings.json"
	scoresFile     = "scores.json"
)

// GameState represents the serializable game state
type GameState struct {
	Board         [3][3]string `json:"board"`
	CurrentPlayer string       `json:"current_player"`
	Status        int          `json:"status"`
	Winner        string       `json:"winner"`
	Mode          int          `json:"mode"`
	MoveHistory   []Position   `json:"move_history"`
}

// Position represents a move position
type Position struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

// Settings represents application settings
type Settings struct {
	GradientType     int     `json:"gradient_type"`
	AIDifficulty     int     `json:"ai_difficulty"`
	AnimationSpeed   float64 `json:"animation_speed"`
	SoundEnabled     bool    `json:"sound_enabled"`
	LastGameMode     int     `json:"last_game_mode"`
	AutoSaveEnabled  bool    `json:"auto_save_enabled"`
}

// Scores represents game statistics
type Scores struct {
	PlayerVsPlayer PlayerVsPlayerStats `json:"player_vs_player"`
	PlayerVsAI     PlayerVsAIStats     `json:"player_vs_ai"`
	TotalGames     int                 `json:"total_games"`
	LastPlayed     string              `json:"last_played"`
}

// PlayerVsPlayerStats represents PvP statistics
type PlayerVsPlayerStats struct {
	XWins  int `json:"x_wins"`
	OWins  int `json:"o_wins"`
	Draws  int `json:"draws"`
	Games  int `json:"games"`
}

// PlayerVsAIStats represents Player vs AI statistics
type PlayerVsAIStats struct {
	Easy       DifficultyStats `json:"easy"`
	Normal     DifficultyStats `json:"normal"`
	Hard       DifficultyStats `json:"hard"`
	INeverLose DifficultyStats `json:"i_never_lose"`
}

// DifficultyStats represents stats for a specific AI difficulty
type DifficultyStats struct {
	PlayerWins int `json:"player_wins"`
	AIWins     int `json:"ai_wins"`
	Draws      int `json:"draws"`
	Games      int `json:"games"`
}

// Manager handles all persistence operations
type Manager struct {
	saveDirectory string
}

// New creates a new persistence manager
func New() *Manager {
	homeDir, _ := os.UserHomeDir()
	saveDir := filepath.Join(homeDir, saveDir)
	
	// Create save directory if it doesn't exist
	os.MkdirAll(saveDir, 0755)
	
	return &Manager{
		saveDirectory: saveDir,
	}
}

// SaveGameState saves the current game state immediately
func (m *Manager) SaveGameState(g *game.Game) error {
	gameState := &GameState{
		CurrentPlayer: string(g.GetCurrentPlayer()),
		Status:        int(g.GetStatus()),
		Winner:        string(g.GetWinner()),
		Mode:          int(g.GetMode()),
		MoveHistory:   []Position{},
	}

	// Convert board
	board := g.GetBoard()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			gameState.Board[i][j] = string(board[i][j])
		}
	}

	// Convert move history
	// Note: We'll need to add a GetMoveHistory method to game.Game
	// For now, initialize empty array

	return m.saveJSON(gameStateFile, gameState)
}

// LoadGameState loads the saved game state
func (m *Manager) LoadGameState() (*game.Game, error) {
	var gameState GameState
	err := m.loadJSON(gameStateFile, &gameState)
	if err != nil {
		// Return new game if no save file exists
		return game.New(), nil
	}

	g := game.New()
	
	// Restore board
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			g.Board[i][j] = game.Player(gameState.Board[i][j])
		}
	}
	
	g.CurrentPlayer = game.Player(gameState.CurrentPlayer)
	g.Status = game.GameStatus(gameState.Status)
	g.Winner = game.Player(gameState.Winner)
	g.SetMode(game.GameMode(gameState.Mode))

	return g, nil
}

// SaveSettings saves application settings immediately
func (m *Manager) SaveSettings(gradientType gradient.GradientType, aiDifficulty ai.Difficulty, animationSpeed float64) error {
	settings := &Settings{
		GradientType:    int(gradientType),
		AIDifficulty:    int(aiDifficulty),
		AnimationSpeed:  animationSpeed,
		SoundEnabled:    false, // Default for now
		AutoSaveEnabled: true,  // Always enabled for this app
	}

	return m.saveJSON(settingsFile, settings)
}

// LoadSettings loads application settings
func (m *Manager) LoadSettings() (*Settings, error) {
	settings := &Settings{
		GradientType:    int(gradient.Rainbow), // Default
		AIDifficulty:    int(ai.Normal),        // Default
		AnimationSpeed:  1.0,                   // Default
		SoundEnabled:    false,
		AutoSaveEnabled: true,
	}

	err := m.loadJSON(settingsFile, settings)
	if err != nil {
		// Return defaults if no settings file exists
		return settings, nil
	}

	return settings, nil
}

// SaveScores saves game statistics immediately
func (m *Manager) SaveScores(scores *Scores) error {
	return m.saveJSON(scoresFile, scores)
}

// LoadScores loads game statistics
func (m *Manager) LoadScores() (*Scores, error) {
	scores := &Scores{
		PlayerVsPlayer: PlayerVsPlayerStats{},
		PlayerVsAI: PlayerVsAIStats{
			Easy:       DifficultyStats{},
			Normal:     DifficultyStats{},
			Hard:       DifficultyStats{},
			INeverLose: DifficultyStats{},
		},
		TotalGames: 0,
	}

	err := m.loadJSON(scoresFile, scores)
	if err != nil {
		// Return empty scores if no file exists
		return scores, nil
	}

	return scores, nil
}

// UpdatePlayerVsPlayerScore updates PvP statistics and saves immediately
func (m *Manager) UpdatePlayerVsPlayerScore(winner game.Player) error {
	scores, err := m.LoadScores()
	if err != nil {
		return err
	}

	scores.PlayerVsPlayer.Games++
	scores.TotalGames++

	switch winner {
	case game.PlayerX:
		scores.PlayerVsPlayer.XWins++
	case game.PlayerO:
		scores.PlayerVsPlayer.OWins++
	case game.Empty:
		scores.PlayerVsPlayer.Draws++
	}

	return m.SaveScores(scores)
}

// UpdatePlayerVsAIScore updates Player vs AI statistics and saves immediately
func (m *Manager) UpdatePlayerVsAIScore(difficulty ai.Difficulty, winner game.Player, aiPlayer game.Player) error {
	scores, err := m.LoadScores()
	if err != nil {
		return err
	}

	var diffStats *DifficultyStats
	switch difficulty {
	case ai.Easy:
		diffStats = &scores.PlayerVsAI.Easy
	case ai.Normal:
		diffStats = &scores.PlayerVsAI.Normal
	case ai.Hard:
		diffStats = &scores.PlayerVsAI.Hard
	case ai.INeverLose:
		diffStats = &scores.PlayerVsAI.INeverLose
	}

	diffStats.Games++
	scores.TotalGames++

	if winner == aiPlayer {
		diffStats.AIWins++
	} else if winner == game.Empty {
		diffStats.Draws++
	} else {
		diffStats.PlayerWins++
	}

	return m.SaveScores(scores)
}

// ClearAllData removes all saved data
func (m *Manager) ClearAllData() error {
	files := []string{gameStateFile, settingsFile, scoresFile}
	
	for _, file := range files {
		filePath := filepath.Join(m.saveDirectory, file)
		if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove %s: %w", file, err)
		}
	}
	
	return nil
}

// GetSaveDirectory returns the save directory path
func (m *Manager) GetSaveDirectory() string {
	return m.saveDirectory
}

// saveJSON saves data as JSON to a file
func (m *Manager) saveJSON(filename string, data interface{}) error {
	filePath := filepath.Join(m.saveDirectory, filename)
	
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}

// loadJSON loads data from a JSON file
func (m *Manager) loadJSON(filename string, target interface{}) error {
	filePath := filepath.Join(m.saveDirectory, filename)
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	err = json.Unmarshal(data, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}