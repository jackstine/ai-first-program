package ai

import (
	"math/rand"
	"time"

	"tic-tac-toe/internal/game"
)

// Difficulty represents AI difficulty levels
type Difficulty int

const (
	Easy Difficulty = iota
	Normal
	Hard
	INeverLose
)

// AI represents the AI player
type AI struct {
	difficulty   Difficulty
	player       game.Player
	opponent     game.Player
	randomSource *rand.Rand
}

// New creates a new AI with specified difficulty and player
func New(difficulty Difficulty, player game.Player) *AI {
	opponent := game.PlayerX
	if player == game.PlayerX {
		opponent = game.PlayerO
	}

	return &AI{
		difficulty:   difficulty,
		player:       player,
		opponent:     opponent,
		randomSource: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// GetMove returns the AI's next move
func (ai *AI) GetMove(g *game.Game) (int, int, error) {
	availableMoves := g.GetAvailableMoves()
	if len(availableMoves) == 0 {
		return -1, -1, nil // No moves available
	}

	switch ai.difficulty {
	case Easy:
		return ai.getEasyMove(g, availableMoves)
	case Normal:
		return ai.getNormalMove(g, availableMoves)
	case Hard:
		return ai.getHardMove(g, availableMoves)
	case INeverLose:
		return ai.getPerfectMove(g, availableMoves)
	default:
		return ai.getEasyMove(g, availableMoves)
	}
}

// getEasyMove - mostly random with occasional basic moves
func (ai *AI) getEasyMove(g *game.Game, availableMoves []game.Position) (int, int, error) {
	// 20% chance to make a smart move, 80% random
	if ai.randomSource.Float32() < 0.2 {
		// Try to win or block
		if row, col := ai.findWinningMove(g, ai.player); row != -1 {
			return row, col, nil
		}
		if row, col := ai.findWinningMove(g, ai.opponent); row != -1 {
			return row, col, nil
		}
	}

	// Random move
	move := availableMoves[ai.randomSource.Intn(len(availableMoves))]
	return move.Row, move.Col, nil
}

// getNormalMove - basic strategy: block wins, take center, corners
func (ai *AI) getNormalMove(g *game.Game, availableMoves []game.Position) (int, int, error) {
	// 1. Try to win
	if row, col := ai.findWinningMove(g, ai.player); row != -1 {
		return row, col, nil
	}

	// 2. Block opponent win
	if row, col := ai.findWinningMove(g, ai.opponent); row != -1 {
		return row, col, nil
	}

	// 3. Take center if available
	if g.IsValidMove(1, 1) {
		return 1, 1, nil
	}

	// 4. Take corners
	corners := []game.Position{{0, 0}, {0, 2}, {2, 0}, {2, 2}}
	for _, corner := range corners {
		if g.IsValidMove(corner.Row, corner.Col) {
			return corner.Row, corner.Col, nil
		}
	}

	// 5. Take any available move
	move := availableMoves[ai.randomSource.Intn(len(availableMoves))]
	return move.Row, move.Col, nil
}

// getHardMove - minimax algorithm with limited depth
func (ai *AI) getHardMove(g *game.Game, availableMoves []game.Position) (int, int, error) {
	bestMove := game.Position{Row: -1, Col: -1}
	bestScore := -1000

	for _, move := range availableMoves {
		// Create a copy of the game for simulation
		testGame := ai.copyGame(g)
		testGame.MakeMove(move.Row, move.Col)

		score := ai.minimax(testGame, 4, false) // Look ahead 4 moves
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}

	if bestMove.Row == -1 {
		// Fallback to normal strategy
		return ai.getNormalMove(g, availableMoves)
	}

	return bestMove.Row, bestMove.Col, nil
}

// getPerfectMove - perfect minimax, never loses
func (ai *AI) getPerfectMove(g *game.Game, availableMoves []game.Position) (int, int, error) {
	bestMove := game.Position{Row: -1, Col: -1}
	bestScore := -1000

	for _, move := range availableMoves {
		// Create a copy of the game for simulation
		testGame := ai.copyGame(g)
		testGame.MakeMove(move.Row, move.Col)

		score := ai.minimax(testGame, 10, false) // Look ahead deeply
		if score > bestScore {
			bestScore = score
			bestMove = move
		}
	}

	if bestMove.Row == -1 {
		// Should never happen, but fallback
		return ai.getHardMove(g, availableMoves)
	}

	return bestMove.Row, bestMove.Col, nil
}

// findWinningMove finds a move that would win for the specified player
func (ai *AI) findWinningMove(g *game.Game, player game.Player) (int, int) {
	board := g.GetBoard()

	// Check rows
	for row := 0; row < 3; row++ {
		count := 0
		emptyCol := -1
		for col := 0; col < 3; col++ {
			if board[row][col] == player {
				count++
			} else if board[row][col] == game.Empty {
				emptyCol = col
			}
		}
		if count == 2 && emptyCol != -1 {
			return row, emptyCol
		}
	}

	// Check columns
	for col := 0; col < 3; col++ {
		count := 0
		emptyRow := -1
		for row := 0; row < 3; row++ {
			if board[row][col] == player {
				count++
			} else if board[row][col] == game.Empty {
				emptyRow = row
			}
		}
		if count == 2 && emptyRow != -1 {
			return emptyRow, col
		}
	}

	// Check diagonal (top-left to bottom-right)
	count := 0
	emptyRow, emptyCol := -1, -1
	for i := 0; i < 3; i++ {
		if board[i][i] == player {
			count++
		} else if board[i][i] == game.Empty {
			emptyRow, emptyCol = i, i
		}
	}
	if count == 2 && emptyRow != -1 {
		return emptyRow, emptyCol
	}

	// Check diagonal (top-right to bottom-left)
	count = 0
	emptyRow, emptyCol = -1, -1
	for i := 0; i < 3; i++ {
		if board[i][2-i] == player {
			count++
		} else if board[i][2-i] == game.Empty {
			emptyRow, emptyCol = i, 2-i
		}
	}
	if count == 2 && emptyRow != -1 {
		return emptyRow, emptyCol
	}

	return -1, -1 // No winning move found
}

// minimax implements the minimax algorithm
func (ai *AI) minimax(g *game.Game, depth int, isMaximizing bool) int {
	status := g.GetStatus()

	// Terminal conditions
	if status == game.StatusWon {
		if g.GetWinner() == ai.player {
			return 10 + depth // Prefer quicker wins
		}
		return -10 - depth // Prefer delayed losses
	}
	if status == game.StatusDraw {
		return 0
	}
	if depth == 0 {
		return 0 // Neutral when depth limit reached
	}

	availableMoves := g.GetAvailableMoves()
	if isMaximizing {
		maxScore := -1000
		for _, move := range availableMoves {
			testGame := ai.copyGame(g)
			testGame.MakeMove(move.Row, move.Col)
			score := ai.minimax(testGame, depth-1, false)
			maxScore = max(maxScore, score)
		}
		return maxScore
	} else {
		minScore := 1000
		for _, move := range availableMoves {
			testGame := ai.copyGame(g)
			testGame.MakeMove(move.Row, move.Col)
			score := ai.minimax(testGame, depth-1, true)
			minScore = min(minScore, score)
		}
		return minScore
	}
}

// copyGame creates a deep copy of the game state
func (ai *AI) copyGame(g *game.Game) *game.Game {
	newGame := game.New()
	newGame.Board = g.GetBoard()
	newGame.CurrentPlayer = g.GetCurrentPlayer()
	newGame.Status = g.GetStatus()
	newGame.Winner = g.GetWinner()
	newGame.Mode = g.GetMode()
	return newGame
}

// GetDifficultyName returns the string name of the difficulty
func (ai *AI) GetDifficultyName() string {
	switch ai.difficulty {
	case Easy:
		return "Easy"
	case Normal:
		return "Normal"
	case Hard:
		return "Hard"
	case INeverLose:
		return "I Never Lose"
	default:
		return "Easy"
	}
}

// GetPlayer returns the AI's player
func (ai *AI) GetPlayer() game.Player {
	return ai.player
}

// SetDifficulty updates the AI difficulty
func (ai *AI) SetDifficulty(difficulty Difficulty) {
	ai.difficulty = difficulty
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}