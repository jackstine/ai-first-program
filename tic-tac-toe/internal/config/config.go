package config

import (
	"fmt"

	"tic-tac-toe/internal/ai"
	"tic-tac-toe/internal/gradient"
	"tic-tac-toe/internal/persistence"
)

// Config represents the application configuration
type Config struct {
	GradientType    gradient.GradientType `json:"gradient_type"`
	AIDifficulty    ai.Difficulty         `json:"ai_difficulty"`
	AnimationSpeed  float64               `json:"animation_speed"`
	SoundEnabled    bool                  `json:"sound_enabled"`
	AutoSaveEnabled bool                  `json:"auto_save_enabled"`
	LastGameMode    int                   `json:"last_game_mode"`
	persistence     *persistence.Manager
}

// New creates a new configuration manager
func New(persistenceManager *persistence.Manager) *Config {
	return &Config{
		GradientType:    gradient.Rainbow,
		AIDifficulty:    ai.Normal,
		AnimationSpeed:  1.0,
		SoundEnabled:    false,
		AutoSaveEnabled: true,
		LastGameMode:    0, // PlayerVsPlayer
		persistence:     persistenceManager,
	}
}

// Load loads configuration from persistence
func (c *Config) Load() error {
	settings, err := c.persistence.LoadSettings()
	if err != nil {
		// If loading fails, keep defaults
		return nil
	}

	c.GradientType = gradient.GradientType(settings.GradientType)
	c.AIDifficulty = ai.Difficulty(settings.AIDifficulty)
	c.AnimationSpeed = settings.AnimationSpeed
	c.SoundEnabled = settings.SoundEnabled
	c.AutoSaveEnabled = settings.AutoSaveEnabled
	c.LastGameMode = settings.LastGameMode

	return nil
}

// Save saves configuration to persistence
func (c *Config) Save() error {
	return c.persistence.SaveSettings(c.GradientType, c.AIDifficulty, c.AnimationSpeed)
}

// GetGradientType returns the current gradient type
func (c *Config) GetGradientType() gradient.GradientType {
	return c.GradientType
}

// SetGradientType sets the gradient type and saves immediately
func (c *Config) SetGradientType(gradientType gradient.GradientType) error {
	c.GradientType = gradientType
	return c.Save()
}

// GetAIDifficulty returns the current AI difficulty
func (c *Config) GetAIDifficulty() ai.Difficulty {
	return c.AIDifficulty
}

// SetAIDifficulty sets the AI difficulty and saves immediately
func (c *Config) SetAIDifficulty(difficulty ai.Difficulty) error {
	c.AIDifficulty = difficulty
	return c.Save()
}

// GetAnimationSpeed returns the current animation speed
func (c *Config) GetAnimationSpeed() float64 {
	return c.AnimationSpeed
}

// SetAnimationSpeed sets the animation speed and saves immediately
func (c *Config) SetAnimationSpeed(speed float64) error {
	if speed < 0.1 {
		speed = 0.1
	}
	if speed > 5.0 {
		speed = 5.0
	}
	
	c.AnimationSpeed = speed
	return c.Save()
}

// IsSoundEnabled returns whether sound is enabled
func (c *Config) IsSoundEnabled() bool {
	return c.SoundEnabled
}

// SetSoundEnabled sets sound enabled state and saves immediately
func (c *Config) SetSoundEnabled(enabled bool) error {
	c.SoundEnabled = enabled
	return c.Save()
}

// IsAutoSaveEnabled returns whether auto-save is enabled
func (c *Config) IsAutoSaveEnabled() bool {
	return c.AutoSaveEnabled
}

// SetAutoSaveEnabled sets auto-save enabled state and saves immediately
func (c *Config) SetAutoSaveEnabled(enabled bool) error {
	c.AutoSaveEnabled = enabled
	return c.Save()
}

// GetLastGameMode returns the last game mode
func (c *Config) GetLastGameMode() int {
	return c.LastGameMode
}

// SetLastGameMode sets the last game mode and saves immediately
func (c *Config) SetLastGameMode(mode int) error {
	c.LastGameMode = mode
	return c.Save()
}

// NextGradientType cycles to the next gradient type
func (c *Config) NextGradientType() error {
	gradientTypes := []gradient.GradientType{
		gradient.Rainbow,
		gradient.Red,
		gradient.Orange,
		gradient.Yellow,
		gradient.Green,
		gradient.Blue,
		gradient.Indigo,
		gradient.Violet,
	}

	currentIndex := 0
	for i, gType := range gradientTypes {
		if gType == c.GradientType {
			currentIndex = i
			break
		}
	}

	nextIndex := (currentIndex + 1) % len(gradientTypes)
	return c.SetGradientType(gradientTypes[nextIndex])
}

// NextAIDifficulty cycles to the next AI difficulty
func (c *Config) NextAIDifficulty() error {
	difficulties := []ai.Difficulty{
		ai.Easy,
		ai.Normal,
		ai.Hard,
		ai.INeverLose,
	}

	currentIndex := 0
	for i, difficulty := range difficulties {
		if difficulty == c.AIDifficulty {
			currentIndex = i
			break
		}
	}

	nextIndex := (currentIndex + 1) % len(difficulties)
	return c.SetAIDifficulty(difficulties[nextIndex])
}

// IncreaseAnimationSpeed increases animation speed by 0.5x
func (c *Config) IncreaseAnimationSpeed() error {
	newSpeed := c.AnimationSpeed + 0.5
	return c.SetAnimationSpeed(newSpeed)
}

// DecreaseAnimationSpeed decreases animation speed by 0.5x
func (c *Config) DecreaseAnimationSpeed() error {
	newSpeed := c.AnimationSpeed - 0.5
	return c.SetAnimationSpeed(newSpeed)
}

// GetGradientTypeName returns the string name of the current gradient type
func (c *Config) GetGradientTypeName() string {
	grad := gradient.New(c.GradientType)
	return grad.GetTypeName()
}

// GetAIDifficultyName returns the string name of the current AI difficulty
func (c *Config) GetAIDifficultyName() string {
	aiPlayer := ai.New(c.AIDifficulty, "O")
	return aiPlayer.GetDifficultyName()
}

// ResetToDefaults resets configuration to default values and saves
func (c *Config) ResetToDefaults() error {
	c.GradientType = gradient.Rainbow
	c.AIDifficulty = ai.Normal
	c.AnimationSpeed = 1.0
	c.SoundEnabled = false
	c.AutoSaveEnabled = true
	c.LastGameMode = 0

	return c.Save()
}

// GetSettingsDisplay returns a formatted display of current settings
func (c *Config) GetSettingsDisplay() string {
	display := "CURRENT SETTINGS:\n"
	display += "─────────────────\n"
	display += "Gradient: " + c.GetGradientTypeName() + "\n"
	display += "AI Difficulty: " + c.GetAIDifficultyName() + "\n"
	display += "Animation Speed: " + fmt.Sprintf("%.1fx", c.AnimationSpeed) + "\n"
	display += "Sound: "
	if c.SoundEnabled {
		display += "Enabled"
	} else {
		display += "Disabled"
	}
	display += "\n"
	display += "Auto-Save: "
	if c.AutoSaveEnabled {
		display += "Enabled"
	} else {
		display += "Disabled"
	}
	display += "\n"

	return display
}

// ValidateConfig validates the current configuration
func (c *Config) ValidateConfig() error {
	// Validate gradient type
	if c.GradientType < 0 || c.GradientType > gradient.Violet {
		c.GradientType = gradient.Rainbow
	}

	// Validate AI difficulty
	if c.AIDifficulty < 0 || c.AIDifficulty > ai.INeverLose {
		c.AIDifficulty = ai.Normal
	}

	// Validate animation speed
	if c.AnimationSpeed < 0.1 || c.AnimationSpeed > 5.0 {
		c.AnimationSpeed = 1.0
	}

	// Validate game mode
	if c.LastGameMode < 0 || c.LastGameMode > 1 {
		c.LastGameMode = 0
	}

	return c.Save()
}