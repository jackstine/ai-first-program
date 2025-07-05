package audio

import (
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/generators"
	"github.com/faiface/beep/speaker"
)

// SoundType represents different types of sounds
type SoundType int

const (
	SoundWin SoundType = iota
	SoundDraw
	SoundMove
	SoundError
)

// Manager handles audio playback
type Manager struct {
	enabled   bool
	sampleRate beep.SampleRate
}

// New creates a new audio manager
func New() *Manager {
	manager := &Manager{
		enabled:    true,
		sampleRate: beep.SampleRate(44100),
	}

	// Initialize speaker
	err := speaker.Init(manager.sampleRate, manager.sampleRate.N(time.Second/10))
	if err != nil {
		// If speaker init fails, disable audio
		manager.enabled = false
	}

	return manager
}

// IsEnabled returns whether audio is enabled
func (m *Manager) IsEnabled() bool {
	return m.enabled
}

// SetEnabled sets the audio enabled state
func (m *Manager) SetEnabled(enabled bool) {
	m.enabled = enabled
}

// PlaySound plays a sound based on the sound type
func (m *Manager) PlaySound(soundType SoundType) {
	if !m.enabled {
		return
	}

	var sound beep.Streamer

	switch soundType {
	case SoundWin:
		sound = m.generatePartySound()
	case SoundDraw:
		sound = m.generateNeutralSound()
	case SoundMove:
		sound = m.generateClickSound()
	case SoundError:
		sound = m.generateErrorSound()
	default:
		return
	}

	// Play sound in background
	if sound != nil {
		speaker.Play(sound)
	}
}

// generatePartySound creates a celebratory melody for wins
func (m *Manager) generatePartySound() beep.Streamer {
	// Create a party-like sequence of notes
	notes := []float64{
		523.25, // C5
		659.25, // E5
		783.99, // G5
		1046.5, // C6
		783.99, // G5
		1046.5, // C6
	}

	var sounds []beep.Streamer

	for i, freq := range notes {
		duration := time.Millisecond * 150
		if i == len(notes)-1 {
			duration = time.Millisecond * 400 // Last note longer
		}

		tone, err := generators.SinTone(m.sampleRate, int(freq))
		if err != nil {
			continue
		}

		// Add some volume envelope for better sound
		envelope := &volumeEnvelope{
			Streamer: beep.Take(int(m.sampleRate.N(duration)), tone),
			volume:   0.3,
		}

		sounds = append(sounds, envelope)

		// Add short pause between notes (except last)
		if i < len(notes)-1 {
			silence := beep.Take(int(m.sampleRate.N(time.Millisecond*50)), beep.Silence(1))
			sounds = append(sounds, silence)
		}
	}

	return beep.Seq(sounds...)
}

// generateNeutralSound creates a neutral sound for draws
func (m *Manager) generateNeutralSound() beep.Streamer {
	tone, err := generators.SinTone(m.sampleRate, 440) // A4
	if err != nil {
		return nil
	}

	return &volumeEnvelope{
		Streamer: beep.Take(int(m.sampleRate.N(time.Millisecond*300)), tone),
		volume:   0.2,
	}
}

// generateClickSound creates a click sound for moves
func (m *Manager) generateClickSound() beep.Streamer {
	tone, err := generators.SinTone(m.sampleRate, 800)
	if err != nil {
		return nil
	}

	return &volumeEnvelope{
		Streamer: beep.Take(int(m.sampleRate.N(time.Millisecond*100)), tone),
		volume:   0.1,
	}
}

// generateErrorSound creates an error sound
func (m *Manager) generateErrorSound() beep.Streamer {
	tone, err := generators.SinTone(m.sampleRate, 200) // Low frequency for error
	if err != nil {
		return nil
	}

	return &volumeEnvelope{
		Streamer: beep.Take(int(m.sampleRate.N(time.Millisecond*500)), tone),
		volume:   0.3,
	}
}

// volumeEnvelope applies volume control to a streamer
type volumeEnvelope struct {
	beep.Streamer
	volume float64
}

func (v *volumeEnvelope) Stream(samples [][2]float64) (n int, ok bool) {
	n, ok = v.Streamer.Stream(samples)
	for i := range samples[:n] {
		samples[i][0] *= v.volume
		samples[i][1] *= v.volume
		
		// Apply fade in/out envelope
		fadeLength := len(samples) / 10 // 10% fade
		if i < fadeLength {
			fade := float64(i) / float64(fadeLength)
			samples[i][0] *= fade
			samples[i][1] *= fade
		} else if i > n-fadeLength {
			fade := float64(n-i) / float64(fadeLength)
			samples[i][0] *= fade
			samples[i][1] *= fade
		}
	}
	return n, ok
}

func (v *volumeEnvelope) Err() error {
	return v.Streamer.Err()
}

// Close shuts down the audio manager
func (m *Manager) Close() {
	if m.enabled {
		speaker.Close()
	}
}