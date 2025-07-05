package persistence_test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"tic-tac-toe/internal/ai"
	"tic-tac-toe/internal/game"
	"tic-tac-toe/internal/gradient"
	"tic-tac-toe/internal/persistence"
)

var _ = Describe("Persistence", func() {
	var (
		manager *persistence.Manager
		tempDir string
	)

	BeforeEach(func() {
		// Create temporary directory for testing
		var err error
		tempDir, err = os.MkdirTemp("", "tic-tac-toe-test")
		Expect(err).ToNot(HaveOccurred())

		// Set HOME to temp directory for testing
		os.Setenv("HOME", tempDir)
		manager = persistence.New()
	})

	AfterEach(func() {
		// Clean up
		os.RemoveAll(tempDir)
	})

	Describe("New", func() {
		It("should create a new persistence manager", func() {
			Expect(manager).ToNot(BeNil())
			Expect(manager.GetSaveDirectory()).To(ContainSubstring(".tic-tac-toe"))
		})
	})

	Describe("SaveGameState and LoadGameState", func() {
		It("should save and load game state", func() {
			g := game.New()
			g.MakeMove(0, 0)
			g.MakeMove(1, 1)

			err := manager.SaveGameState(g)
			Expect(err).ToNot(HaveOccurred())

			loadedGame, err := manager.LoadGameState()
			Expect(err).ToNot(HaveOccurred())
			Expect(loadedGame.GetCurrentPlayer()).To(Equal(g.GetCurrentPlayer()))
			Expect(loadedGame.GetStatus()).To(Equal(g.GetStatus()))
		})

		It("should return new game when no save file exists", func() {
			loadedGame, err := manager.LoadGameState()
			Expect(err).ToNot(HaveOccurred())
			Expect(loadedGame.GetCurrentPlayer()).To(Equal(game.PlayerX))
			Expect(loadedGame.GetStatus()).To(Equal(game.StatusPlaying))
		})
	})

	Describe("SaveSettings and LoadSettings", func() {
		It("should save and load settings", func() {
			err := manager.SaveSettings(gradient.Red, ai.Hard, 2.0)
			Expect(err).ToNot(HaveOccurred())

			settings, err := manager.LoadSettings()
			Expect(err).ToNot(HaveOccurred())
			Expect(settings.GradientType).To(Equal(int(gradient.Red)))
			Expect(settings.AIDifficulty).To(Equal(int(ai.Hard)))
			Expect(settings.AnimationSpeed).To(Equal(2.0))
		})

		It("should return defaults when no settings file exists", func() {
			settings, err := manager.LoadSettings()
			Expect(err).ToNot(HaveOccurred())
			Expect(settings.GradientType).To(Equal(int(gradient.Rainbow)))
			Expect(settings.AIDifficulty).To(Equal(int(ai.Normal)))
			Expect(settings.AnimationSpeed).To(Equal(1.0))
		})
	})

	Describe("SaveScores and LoadScores", func() {
		It("should save and load scores", func() {
			scores, err := manager.LoadScores()
			Expect(err).ToNot(HaveOccurred())

			scores.TotalGames = 5
			scores.PlayerVsPlayer.XWins = 2

			err = manager.SaveScores(scores)
			Expect(err).ToNot(HaveOccurred())

			loadedScores, err := manager.LoadScores()
			Expect(err).ToNot(HaveOccurred())
			Expect(loadedScores.TotalGames).To(Equal(5))
			Expect(loadedScores.PlayerVsPlayer.XWins).To(Equal(2))
		})
	})

	Describe("UpdatePlayerVsPlayerScore", func() {
		It("should update PvP statistics", func() {
			err := manager.UpdatePlayerVsPlayerScore(game.PlayerX)
			Expect(err).ToNot(HaveOccurred())

			scores, err := manager.LoadScores()
			Expect(err).ToNot(HaveOccurred())
			Expect(scores.PlayerVsPlayer.XWins).To(Equal(1))
			Expect(scores.PlayerVsPlayer.Games).To(Equal(1))
			Expect(scores.TotalGames).To(Equal(1))
		})

		It("should handle draws", func() {
			err := manager.UpdatePlayerVsPlayerScore(game.Empty)
			Expect(err).ToNot(HaveOccurred())

			scores, err := manager.LoadScores()
			Expect(err).ToNot(HaveOccurred())
			Expect(scores.PlayerVsPlayer.Draws).To(Equal(1))
		})
	})

	Describe("UpdatePlayerVsAIScore", func() {
		It("should update Player vs AI statistics", func() {
			err := manager.UpdatePlayerVsAIScore(ai.Hard, game.PlayerX, game.PlayerO)
			Expect(err).ToNot(HaveOccurred())

			scores, err := manager.LoadScores()
			Expect(err).ToNot(HaveOccurred())
			Expect(scores.PlayerVsAI.Hard.PlayerWins).To(Equal(1))
			Expect(scores.PlayerVsAI.Hard.Games).To(Equal(1))
		})

		It("should track AI wins", func() {
			err := manager.UpdatePlayerVsAIScore(ai.Easy, game.PlayerO, game.PlayerO)
			Expect(err).ToNot(HaveOccurred())

			scores, err := manager.LoadScores()
			Expect(err).ToNot(HaveOccurred())
			Expect(scores.PlayerVsAI.Easy.AIWins).To(Equal(1))
		})
	})

	Describe("ClearAllData", func() {
		It("should remove all save files", func() {
			// Create some data first
			g := game.New()
			manager.SaveGameState(g)
			manager.SaveSettings(gradient.Rainbow, ai.Normal, 1.0)

			err := manager.ClearAllData()
			Expect(err).ToNot(HaveOccurred())

			// Check that files are gone
			saveDir := manager.GetSaveDirectory()
			files := []string{"gamestate.json", "settings.json", "scores.json"}
			
			for _, file := range files {
				filePath := filepath.Join(saveDir, file)
				_, err := os.Stat(filePath)
				Expect(os.IsNotExist(err)).To(BeTrue())
			}
		})
	})
})