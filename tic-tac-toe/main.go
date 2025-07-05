package main

import (
	"fmt"
	"log"

	"tic-tac-toe/internal/game"
	"tic-tac-toe/internal/ui"
)

func main() {
	fmt.Println("🎮 Tic-Tac-Toe Game Starting...")
	
	// Initialize game
	g := game.New()
	
	// Start UI
	if err := ui.Start(g); err != nil {
		log.Fatal("Failed to start game:", err)
	}
}