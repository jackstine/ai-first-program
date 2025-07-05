package game_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"tic-tac-toe/internal/game"
)

var _ = Describe("Game", func() {
	var g *game.Game

	BeforeEach(func() {
		g = game.New()
	})

	Describe("New", func() {
		It("should create a new game with correct initial state", func() {
			Expect(g.GetCurrentPlayer()).To(Equal(game.PlayerX))
			Expect(g.GetStatus()).To(Equal(game.StatusPlaying))
			Expect(g.GetWinner()).To(Equal(game.Empty))
			Expect(g.GetMode()).To(Equal(game.PlayerVsPlayer))
			
			board := g.GetBoard()
			for row := 0; row < 3; row++ {
				for col := 0; col < 3; col++ {
					Expect(board[row][col]).To(Equal(game.Empty))
				}
			}
		})
	})

	Describe("MakeMove", func() {
		It("should make a valid move", func() {
			err := g.MakeMove(0, 0)
			Expect(err).ToNot(HaveOccurred())
			
			board := g.GetBoard()
			Expect(board[0][0]).To(Equal(game.PlayerX))
			Expect(g.GetCurrentPlayer()).To(Equal(game.PlayerO))
		})

		It("should return error for invalid position", func() {
			err := g.MakeMove(-1, 0)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("invalid position"))

			err = g.MakeMove(3, 0)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("invalid position"))
		})

		It("should return error for occupied position", func() {
			err := g.MakeMove(0, 0)
			Expect(err).ToNot(HaveOccurred())

			err = g.MakeMove(0, 0)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("already occupied"))
		})

		It("should detect win condition - row", func() {
			g.MakeMove(0, 0) // X
			g.MakeMove(1, 0) // O
			g.MakeMove(0, 1) // X
			g.MakeMove(1, 1) // O
			g.MakeMove(0, 2) // X wins

			Expect(g.GetStatus()).To(Equal(game.StatusWon))
			Expect(g.GetWinner()).To(Equal(game.PlayerX))
		})

		It("should detect win condition - column", func() {
			g.MakeMove(0, 0) // X
			g.MakeMove(0, 1) // O
			g.MakeMove(1, 0) // X
			g.MakeMove(0, 2) // O
			g.MakeMove(2, 0) // X wins

			Expect(g.GetStatus()).To(Equal(game.StatusWon))
			Expect(g.GetWinner()).To(Equal(game.PlayerX))
		})

		It("should detect win condition - diagonal", func() {
			g.MakeMove(0, 0) // X
			g.MakeMove(0, 1) // O
			g.MakeMove(1, 1) // X
			g.MakeMove(0, 2) // O
			g.MakeMove(2, 2) // X wins

			Expect(g.GetStatus()).To(Equal(game.StatusWon))
			Expect(g.GetWinner()).To(Equal(game.PlayerX))
		})

		It("should detect draw condition", func() {
			// X O X
			// O O X
			// O X O
			moves := [][]int{
				{0, 0}, {0, 1}, {0, 2}, // X O X
				{1, 1}, {1, 0}, {1, 2}, // O X X -> O O X
				{2, 1}, {2, 0}, {2, 2}, // X O O -> O X O
			}

			for _, move := range moves {
				g.MakeMove(move[0], move[1])
			}

			Expect(g.GetStatus()).To(Equal(game.StatusDraw))
		})
	})

	Describe("Reset", func() {
		It("should reset the game to initial state", func() {
			g.MakeMove(0, 0)
			g.MakeMove(1, 1)
			
			g.Reset()
			
			Expect(g.GetCurrentPlayer()).To(Equal(game.PlayerX))
			Expect(g.GetStatus()).To(Equal(game.StatusPlaying))
			Expect(g.GetWinner()).To(Equal(game.Empty))
			
			board := g.GetBoard()
			for row := 0; row < 3; row++ {
				for col := 0; col < 3; col++ {
					Expect(board[row][col]).To(Equal(game.Empty))
				}
			}
		})
	})

	Describe("IsValidMove", func() {
		It("should return true for valid moves", func() {
			Expect(g.IsValidMove(0, 0)).To(BeTrue())
			Expect(g.IsValidMove(1, 1)).To(BeTrue())
			Expect(g.IsValidMove(2, 2)).To(BeTrue())
		})

		It("should return false for invalid positions", func() {
			Expect(g.IsValidMove(-1, 0)).To(BeFalse())
			Expect(g.IsValidMove(3, 0)).To(BeFalse())
			Expect(g.IsValidMove(0, -1)).To(BeFalse())
			Expect(g.IsValidMove(0, 3)).To(BeFalse())
		})

		It("should return false for occupied positions", func() {
			g.MakeMove(0, 0)
			Expect(g.IsValidMove(0, 0)).To(BeFalse())
		})
	})

	Describe("GetAvailableMoves", func() {
		It("should return all positions when board is empty", func() {
			moves := g.GetAvailableMoves()
			Expect(len(moves)).To(Equal(9))
		})

		It("should return correct available moves", func() {
			g.MakeMove(0, 0)
			g.MakeMove(1, 1)
			
			moves := g.GetAvailableMoves()
			Expect(len(moves)).To(Equal(7))
		})
	})
})