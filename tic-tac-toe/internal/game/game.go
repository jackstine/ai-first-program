package game

import "fmt"

// Player represents a game player
type Player string

const (
	PlayerX Player = "X"
	PlayerO Player = "O"
	Empty   Player = " "
)

// GameMode represents different game modes
type GameMode int

const (
	PlayerVsPlayer GameMode = iota
	PlayerVsAI
)

// GameStatus represents the current state of the game
type GameStatus int

const (
	StatusPlaying GameStatus = iota
	StatusWon
	StatusDraw
)

// Position represents a position on the board
type Position struct {
	Row int
	Col int
}

// Game represents the tic-tac-toe game state
type Game struct {
	Board         [3][3]Player `json:"board"`
	CurrentPlayer Player       `json:"current_player"`
	Status        GameStatus   `json:"status"`
	Winner        Player       `json:"winner"`
	Mode          GameMode     `json:"mode"`
	MoveHistory   []Position   `json:"move_history"`
}

// New creates a new game instance
func New() *Game {
	board := [3][3]Player{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = Empty
		}
	}
	
	return &Game{
		Board:         board,
		CurrentPlayer: PlayerX,
		Status:        StatusPlaying,
		Winner:        Empty,
		Mode:          PlayerVsPlayer,
		MoveHistory:   make([]Position, 0),
	}
}

// MakeMove attempts to make a move at the specified position
func (g *Game) MakeMove(row, col int) error {
	if g.Status != StatusPlaying {
		return fmt.Errorf("game is not in playing state")
	}

	if row < 0 || row > 2 || col < 0 || col > 2 {
		return fmt.Errorf("invalid position: (%d, %d)", row, col)
	}

	if g.Board[row][col] != Empty {
		return fmt.Errorf("position (%d, %d) is already occupied", row, col)
	}

	// Make the move
	g.Board[row][col] = g.CurrentPlayer
	g.MoveHistory = append(g.MoveHistory, Position{Row: row, Col: col})

	// Check for win or draw
	g.checkGameStatus()

	// Switch players if game is still playing
	if g.Status == StatusPlaying {
		g.switchPlayer()
	}

	return nil
}

// Reset resets the game to initial state
func (g *Game) Reset() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			g.Board[i][j] = Empty
		}
	}
	g.CurrentPlayer = PlayerX
	g.Status = StatusPlaying
	g.Winner = Empty
	g.MoveHistory = make([]Position, 0)
}

// GetBoard returns a copy of the current board
func (g *Game) GetBoard() [3][3]Player {
	return g.Board
}

// GetCurrentPlayer returns the current player
func (g *Game) GetCurrentPlayer() Player {
	return g.CurrentPlayer
}

// GetStatus returns the current game status
func (g *Game) GetStatus() GameStatus {
	return g.Status
}

// GetWinner returns the winner (if any)
func (g *Game) GetWinner() Player {
	return g.Winner
}

// SetMode sets the game mode
func (g *Game) SetMode(mode GameMode) {
	g.Mode = mode
}

// GetMode returns the current game mode
func (g *Game) GetMode() GameMode {
	return g.Mode
}

// GetMoveHistory returns the move history
func (g *Game) GetMoveHistory() []Position {
	return g.MoveHistory
}

// switchPlayer switches the current player
func (g *Game) switchPlayer() {
	if g.CurrentPlayer == PlayerX {
		g.CurrentPlayer = PlayerO
	} else {
		g.CurrentPlayer = PlayerX
	}
}

// checkGameStatus checks if the game has ended (win or draw)
func (g *Game) checkGameStatus() {
	// Check for win
	if winner := g.checkWinner(); winner != Empty {
		g.Status = StatusWon
		g.Winner = winner
		return
	}

	// Check for draw
	if g.isBoardFull() {
		g.Status = StatusDraw
		return
	}
}

// checkWinner checks if there's a winner
func (g *Game) checkWinner() Player {
	// Check rows
	for row := 0; row < 3; row++ {
		if g.Board[row][0] != Empty &&
			g.Board[row][0] == g.Board[row][1] &&
			g.Board[row][1] == g.Board[row][2] {
			return g.Board[row][0]
		}
	}

	// Check columns
	for col := 0; col < 3; col++ {
		if g.Board[0][col] != Empty &&
			g.Board[0][col] == g.Board[1][col] &&
			g.Board[1][col] == g.Board[2][col] {
			return g.Board[0][col]
		}
	}

	// Check diagonals
	if g.Board[0][0] != Empty &&
		g.Board[0][0] == g.Board[1][1] &&
		g.Board[1][1] == g.Board[2][2] {
		return g.Board[0][0]
	}

	if g.Board[0][2] != Empty &&
		g.Board[0][2] == g.Board[1][1] &&
		g.Board[1][1] == g.Board[2][0] {
		return g.Board[0][2]
	}

	return Empty
}

// isBoardFull checks if the board is full
func (g *Game) isBoardFull() bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if g.Board[row][col] == Empty {
				return false
			}
		}
	}
	return true
}

// GetAvailableMoves returns all available positions
func (g *Game) GetAvailableMoves() []Position {
	var moves []Position
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if g.Board[row][col] == Empty {
				moves = append(moves, Position{Row: row, Col: col})
			}
		}
	}
	return moves
}

// IsValidMove checks if a move is valid
func (g *Game) IsValidMove(row, col int) bool {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return false
	}
	return g.Board[row][col] == Empty && g.Status == StatusPlaying
}