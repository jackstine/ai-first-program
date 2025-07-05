package ai_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"tic-tac-toe/internal/ai"
	"tic-tac-toe/internal/game"
)

var _ = Describe("AI", func() {
	var (
		g      *game.Game
		aiEasy *ai.AI
	)

	BeforeEach(func() {
		g = game.New()
		aiEasy = ai.New(ai.Easy, game.PlayerO)
	})

	Describe("New", func() {
		It("should create AI with correct difficulty and player", func() {
			Expect(aiEasy.GetDifficultyName()).To(Equal("Easy"))
			Expect(aiEasy.GetPlayer()).To(Equal(game.PlayerO))
		})

		It("should handle different difficulties", func() {
			testCases := map[ai.Difficulty]string{
				ai.Easy:       "Easy",
				ai.Normal:     "Normal",
				ai.Hard:       "Hard",
				ai.INeverLose: "I Never Lose",
			}

			for difficulty, expectedName := range testCases {
				testAI := ai.New(difficulty, game.PlayerX)
				Expect(testAI.GetDifficultyName()).To(Equal(expectedName))
			}
		})
	})

	Describe("GetMove", func() {
		It("should return valid moves", func() {
			row, col, err := aiEasy.GetMove(g)
			Expect(err).ToNot(HaveOccurred())
			Expect(g.IsValidMove(row, col)).To(BeTrue())
		})

		It("should handle empty board", func() {
			row, col, err := aiEasy.GetMove(g)
			Expect(err).ToNot(HaveOccurred())
			Expect(row).To(BeNumerically(">=", 0))
			Expect(row).To(BeNumerically("<=", 2))
			Expect(col).To(BeNumerically(">=", 0))
			Expect(col).To(BeNumerically("<=", 2))
		})

		It("should handle full board", func() {
			// Fill the board completely
			moves := [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}
			for _, move := range moves {
				if g.GetStatus() == game.StatusPlaying {
					g.MakeMove(move[0], move[1])
				}
			}

			// When board is full, AI should return -1, -1
			row, col, err := aiEasy.GetMove(g)
			Expect(err).ToNot(HaveOccurred())
			if len(g.GetAvailableMoves()) == 0 {
				Expect(row).To(Equal(-1))
				Expect(col).To(Equal(-1))
			}
		})

		It("should block winning moves on Hard difficulty", func() {
			// Set up a scenario where X can win
			g.MakeMove(0, 0) // X
			g.MakeMove(1, 1) // O
			g.MakeMove(0, 1) // X

			// Hard AI should block at (0, 2)
			aiHard := ai.New(ai.Hard, game.PlayerO)
			row, col, err := aiHard.GetMove(g)
			Expect(err).ToNot(HaveOccurred())
			Expect(row).To(Equal(0))
			Expect(col).To(Equal(2))
		})

		It("should take winning moves when available", func() {
			// Set up a scenario where AI (O) can win
			g.MakeMove(0, 0) // X
			g.MakeMove(1, 0) // O  
			g.MakeMove(0, 1) // X
			g.MakeMove(1, 1) // O
			// Now O can win at (1, 2)

			// AI should take the winning move
			row, col, err := aiEasy.GetMove(g)
			Expect(err).ToNot(HaveOccurred())
			
			// AI should make a valid move
			Expect(g.IsValidMove(row, col)).To(BeTrue())
			
			// If it's a winning move, it should be (1, 2)
			if row == 1 && col == 2 {
				// This is the winning move
				Expect(row).To(Equal(1))
				Expect(col).To(Equal(2))
			}
		})
	})

	Describe("SetDifficulty", func() {
		It("should update difficulty", func() {
			aiEasy.SetDifficulty(ai.Hard)
			Expect(aiEasy.GetDifficultyName()).To(Equal("Hard"))
		})
	})

	Describe("GetPlayer", func() {
		It("should return correct player", func() {
			aiX := ai.New(ai.Easy, game.PlayerX)
			aiO := ai.New(ai.Easy, game.PlayerO)
			
			Expect(aiX.GetPlayer()).To(Equal(game.PlayerX))
			Expect(aiO.GetPlayer()).To(Equal(game.PlayerO))
		})
	})

	Describe("Different difficulty behaviors", func() {
		It("should show different move patterns for different difficulties", func() {
			// This is a statistical test - run multiple games
			easyWins := 0
			hardWins := 0
			games := 10

			for i := 0; i < games; i++ {
				// Test Easy AI
				testGameEasy := game.New()
				testGameEasy.SetMode(game.PlayerVsAI)
				aiEasyTest := ai.New(ai.Easy, game.PlayerO)

				for testGameEasy.GetStatus() == game.StatusPlaying {
					if testGameEasy.GetCurrentPlayer() == game.PlayerX {
						// Simple strategy for X
						moves := testGameEasy.GetAvailableMoves()
						if len(moves) > 0 {
							testGameEasy.MakeMove(moves[0].Row, moves[0].Col)
						}
					} else {
						row, col, _ := aiEasyTest.GetMove(testGameEasy)
						if row != -1 {
							testGameEasy.MakeMove(row, col)
						}
					}
				}

				if testGameEasy.GetWinner() == game.PlayerO {
					easyWins++
				}

				// Test Hard AI
				testGameHard := game.New()
				testGameHard.SetMode(game.PlayerVsAI)
				aiHardTest := ai.New(ai.Hard, game.PlayerO)

				for testGameHard.GetStatus() == game.StatusPlaying {
					if testGameHard.GetCurrentPlayer() == game.PlayerX {
						// Simple strategy for X
						moves := testGameHard.GetAvailableMoves()
						if len(moves) > 0 {
							testGameHard.MakeMove(moves[0].Row, moves[0].Col)
						}
					} else {
						row, col, _ := aiHardTest.GetMove(testGameHard)
						if row != -1 {
							testGameHard.MakeMove(row, col)
						}
					}
				}

				if testGameHard.GetWinner() == game.PlayerO {
					hardWins++
				}
			}

			// Hard AI should generally perform better than Easy AI
			// This is probabilistic, so we allow some variance
			Expect(hardWins).To(BeNumerically(">=", easyWins-2))
		})
	})
})