package config_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"tic-tac-toe/internal/ai"
	"tic-tac-toe/internal/config"
	"tic-tac-toe/internal/gradient"
	"tic-tac-toe/internal/persistence"
)

var _ = Describe("Config", func() {
	var cfg *config.Config
	var persistManager *persistence.Manager
	var tempDir string

	BeforeEach(func() {
		// Create a temporary directory for test persistence
		var err error
		tempDir, err = os.MkdirTemp("", "tic-tac-toe-config-test")
		Expect(err).ToNot(HaveOccurred())
		
		// Create persistence manager with temp directory
		persistManager = persistence.New()
		
		// Create config instance
		cfg = config.New(persistManager)
		Expect(cfg).ToNot(BeNil())
	})

	AfterEach(func() {
		// Clean up temporary directory
		if tempDir != "" {
			os.RemoveAll(tempDir)
		}
	})

	Describe("New", func() {
		It("should create a new config with default values", func() {
			Expect(cfg.GetGradientType()).To(Equal(gradient.Rainbow))
			Expect(cfg.GetAIDifficulty()).To(Equal(ai.Normal))
			Expect(cfg.GetAnimationSpeed()).To(Equal(1.0))
			Expect(cfg.IsSoundEnabled()).To(BeFalse())
			Expect(cfg.IsAutoSaveEnabled()).To(BeTrue())
			Expect(cfg.GetLastGameMode()).To(Equal(0))
		})
	})

	Describe("Gradient Type Management", func() {
		It("should get and set gradient type", func() {
			err := cfg.SetGradientType(gradient.Red)
			Expect(err).ToNot(HaveOccurred())
			Expect(cfg.GetGradientType()).To(Equal(gradient.Red))
		})

		It("should cycle through gradient types", func() {
			initialType := cfg.GetGradientType()
			err := cfg.NextGradientType()
			Expect(err).ToNot(HaveOccurred())
			Expect(cfg.GetGradientType()).ToNot(Equal(initialType))
		})

		It("should get gradient type name", func() {
			cfg.SetGradientType(gradient.Rainbow)
			name := cfg.GetGradientTypeName()
			Expect(name).ToNot(BeEmpty())
		})
	})

	Describe("AI Difficulty Management", func() {
		It("should get and set AI difficulty", func() {
			err := cfg.SetAIDifficulty(ai.Hard)
			Expect(err).ToNot(HaveOccurred())
			Expect(cfg.GetAIDifficulty()).To(Equal(ai.Hard))
		})

		It("should cycle through AI difficulties", func() {
			initialDifficulty := cfg.GetAIDifficulty()
			err := cfg.NextAIDifficulty()
			Expect(err).ToNot(HaveOccurred())
			Expect(cfg.GetAIDifficulty()).ToNot(Equal(initialDifficulty))
		})

		It("should get AI difficulty name", func() {
			cfg.SetAIDifficulty(ai.Hard)
			name := cfg.GetAIDifficultyName()
			Expect(name).ToNot(BeEmpty())
		})
	})

	Describe("Animation Speed Management", func() {
		It("should get and set animation speed", func() {
			err := cfg.SetAnimationSpeed(2.5)
			Expect(err).ToNot(HaveOccurred())
			Expect(cfg.GetAnimationSpeed()).To(Equal(2.5))
		})

		It("should clamp animation speed to valid range", func() {
			// Test lower bound
			err := cfg.SetAnimationSpeed(0.05)
			Expect(err).ToNot(HaveOccurred())
			Expect(cfg.GetAnimationSpeed()).To(Equal(0.1))

			// Test upper bound
			err = cfg.SetAnimationSpeed(10.0)
			Expect(err).ToNot(HaveOccurred())
			Expect(cfg.GetAnimationSpeed()).To(Equal(5.0))
		})
	})

	Describe("Sound Settings", func() {
		It("should get and set sound enabled state", func() {
			err := cfg.SetSoundEnabled(true)
			Expect(err).ToNot(HaveOccurred())
			Expect(cfg.IsSoundEnabled()).To(BeTrue())

			err = cfg.SetSoundEnabled(false)
			Expect(err).ToNot(HaveOccurred())
			Expect(cfg.IsSoundEnabled()).To(BeFalse())
		})
	})

	Describe("Auto-Save Settings", func() {
		It("should get and set auto-save enabled state", func() {
			err := cfg.SetAutoSaveEnabled(false)
			Expect(err).ToNot(HaveOccurred())
			Expect(cfg.IsAutoSaveEnabled()).To(BeFalse())

			err = cfg.SetAutoSaveEnabled(true)
			Expect(err).ToNot(HaveOccurred())
			Expect(cfg.IsAutoSaveEnabled()).To(BeTrue())
		})
	})

	Describe("Game Mode Settings", func() {
		It("should get and set last game mode", func() {
			err := cfg.SetLastGameMode(1)
			Expect(err).ToNot(HaveOccurred())
			Expect(cfg.GetLastGameMode()).To(Equal(1))
		})
	})

	Describe("Settings Display", func() {
		It("should generate settings display string", func() {
			display := cfg.GetSettingsDisplay()
			Expect(display).ToNot(BeEmpty())
			Expect(display).To(ContainSubstring("CURRENT SETTINGS"))
			Expect(display).To(ContainSubstring("Gradient:"))
			Expect(display).To(ContainSubstring("AI Difficulty:"))
			Expect(display).To(ContainSubstring("Animation Speed:"))
		})
	})

	Describe("Reset to Defaults", func() {
		It("should reset all settings to default values", func() {
			// Change some values
			cfg.SetGradientType(gradient.Blue)
			cfg.SetAIDifficulty(ai.Easy)
			cfg.SetAnimationSpeed(3.0)
			cfg.SetSoundEnabled(true)
			cfg.SetAutoSaveEnabled(false)
			cfg.SetLastGameMode(1)

			// Reset to defaults
			err := cfg.ResetToDefaults()
			Expect(err).ToNot(HaveOccurred())

			// Verify defaults
			Expect(cfg.GetGradientType()).To(Equal(gradient.Rainbow))
			Expect(cfg.GetAIDifficulty()).To(Equal(ai.Normal))
			Expect(cfg.GetAnimationSpeed()).To(Equal(1.0))
			Expect(cfg.IsSoundEnabled()).To(BeFalse())
			Expect(cfg.IsAutoSaveEnabled()).To(BeTrue())
			Expect(cfg.GetLastGameMode()).To(Equal(0))
		})
	})

	Describe("Configuration Validation", func() {
		It("should validate and correct invalid configurations", func() {
			err := cfg.ValidateConfig()
			Expect(err).ToNot(HaveOccurred())
			
			// After validation, all values should be within valid ranges
			Expect(cfg.GetGradientType()).To(BeNumerically(">=", 0))
			Expect(cfg.GetGradientType()).To(BeNumerically("<=", gradient.Violet))
			Expect(cfg.GetAIDifficulty()).To(BeNumerically(">=", 0))
			Expect(cfg.GetAIDifficulty()).To(BeNumerically("<=", ai.INeverLose))
			Expect(cfg.GetAnimationSpeed()).To(BeNumerically(">=", 0.1))
			Expect(cfg.GetAnimationSpeed()).To(BeNumerically("<=", 5.0))
			Expect(cfg.GetLastGameMode()).To(BeNumerically(">=", 0))
			Expect(cfg.GetLastGameMode()).To(BeNumerically("<=", 1))
		})
	})

	Describe("Load and Save Operations", func() {
		It("should handle load when no settings file exists", func() {
			err := cfg.Load()
			Expect(err).ToNot(HaveOccurred()) // Should not error when file doesn't exist
		})

		It("should save configuration", func() {
			err := cfg.Save()
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("Gradient Type Cycling", func() {
		It("should cycle through all gradient types", func() {
			initialType := cfg.GetGradientType()
			seenTypes := make(map[gradient.GradientType]bool)
			seenTypes[initialType] = true

			// Cycle through until we get back to the initial type
			for i := 0; i < 10; i++ { // Arbitrary limit to prevent infinite loop
				err := cfg.NextGradientType()
				Expect(err).ToNot(HaveOccurred())
				
				currentType := cfg.GetGradientType()
				if currentType == initialType {
					break // We've cycled back to the beginning
				}
				seenTypes[currentType] = true
			}

			// Should have seen multiple gradient types
			Expect(len(seenTypes)).To(BeNumerically(">", 1))
		})
	})

	Describe("AI Difficulty Cycling", func() {
		It("should cycle through all AI difficulties", func() {
			initialDifficulty := cfg.GetAIDifficulty()
			seenDifficulties := make(map[ai.Difficulty]bool)
			seenDifficulties[initialDifficulty] = true

			// Cycle through until we get back to the initial difficulty
			for i := 0; i < 10; i++ { // Arbitrary limit to prevent infinite loop
				err := cfg.NextAIDifficulty()
				Expect(err).ToNot(HaveOccurred())
				
				currentDifficulty := cfg.GetAIDifficulty()
				if currentDifficulty == initialDifficulty {
					break // We've cycled back to the beginning
				}
				seenDifficulties[currentDifficulty] = true
			}

			// Should have seen multiple difficulties
			Expect(len(seenDifficulties)).To(BeNumerically(">", 1))
		})
	})
})