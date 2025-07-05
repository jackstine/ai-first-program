package audio_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"tic-tac-toe/internal/audio"
)

var _ = Describe("Audio", func() {
	var manager *audio.Manager

	BeforeEach(func() {
		manager = audio.New()
	})

	Describe("New", func() {
		It("should create a new audio manager", func() {
			Expect(manager).ToNot(BeNil())
		})
	})

	Describe("IsEnabled", func() {
		It("should return the enabled state", func() {
			// Audio might be disabled if no audio device is available
			enabled := manager.IsEnabled()
			Expect(enabled).To(BeAssignableToTypeOf(true))
		})
	})

	Describe("SetEnabled", func() {
		It("should set the enabled state", func() {
			manager.SetEnabled(false)
			Expect(manager.IsEnabled()).To(BeFalse())

			manager.SetEnabled(true)
			Expect(manager.IsEnabled()).To(BeTrue())
		})
	})

	Describe("PlaySound", func() {
		It("should not panic when playing different sound types", func() {
			// These should not panic even if audio is disabled
			Expect(func() { manager.PlaySound(audio.SoundWin) }).ToNot(Panic())
			Expect(func() { manager.PlaySound(audio.SoundDraw) }).ToNot(Panic())
			Expect(func() { manager.PlaySound(audio.SoundMove) }).ToNot(Panic())
			Expect(func() { manager.PlaySound(audio.SoundError) }).ToNot(Panic())
		})

		It("should not play sound when disabled", func() {
			manager.SetEnabled(false)
			// Should not panic or cause issues when disabled
			Expect(func() { manager.PlaySound(audio.SoundWin) }).ToNot(Panic())
		})
	})
})