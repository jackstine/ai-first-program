# Tic-Tac-Toe Bubble Tea Project Plan

## Project Overview
Create a terminal-based tic-tac-toe game using Go and the Bubble Tea framework for learning purposes.

## Key Requirements
1. Two-player tic-tac-toe game (X and O)
2. Terminal user interface using Bubble Tea
3. Keyboard navigation for moves
4. Win detection and game state management
5. Clean, intuitive interface

## Technical Architecture

### Core Components
1. **Game State** - Track board, current player, game status
2. **Game Logic** - Move validation, win checking, turn management
3. **UI Model** - Bubble Tea model for rendering and updates
4. **Input Handling** - Keyboard controls for player moves

### Data Structures
```go
type GameState struct {
    Board        [3][3]string  // Game board
    CurrentPlayer string       // "X" or "O"
    Status       string        // "playing", "won", "draw"
    Winner       string        // Winner if game is won
}

type Model struct {
    Game         GameState
    CursorX      int          // Current cursor position
    CursorY      int
}
```

### Game Rules
- 3x3 grid
- Players alternate turns (X starts first)
- Win conditions: 3 in a row (horizontal, vertical, diagonal)
- Draw condition: board full with no winner

## Implementation Phases

### Phase 1: Basic Setup
- Initialize Go module
- Set up basic project structure
- Create main.go with entry point

### Phase 2: Core Game Logic
- Implement game state structure
- Add move validation
- Create win condition checking
- Add turn management

### Phase 3: Bubble Tea Integration
- Set up Bubble Tea framework
- Create basic model and view
- Implement update logic
- Add keyboard input handling

### Phase 4: UI Polish
- Improve board display
- Add status messages
- Create win/draw screens
- Polish user experience

## Controls Design
- Arrow keys: Move cursor
- Enter/Space: Place mark
- 'r': Reset game
- 'q': Quit game

## Display Layout
```
   1   2   3
1     |   | X
  -----------
2  O  | X |  
  -----------
3     |   |  

Player X's turn
Use arrow keys to move, Enter to place
```

## Technical Dependencies
- Go 1.21+
- Bubble Tea framework (github.com/charmbracelet/bubbletea)
- Lipgloss for styling (github.com/charmbracelet/lipgloss)

## Learning Objectives
1. Understand Bubble Tea's Model-Update-View pattern
2. Practice Go struct design and methods
3. Implement game logic and state management
4. Create interactive terminal applications
5. Handle keyboard input in Go applications

## Success Criteria
- Game runs without errors
- All win conditions work correctly
- Intuitive user controls
- Clean, readable code structure
- Proper error handling

---

## 🆕 **ADVANCED FEATURES EXPANSION**

### 🔊 **Sound Effects System**

#### **Technical Architecture**
- **Audio Backend**: Use `github.com/hajimehoshi/ebiten/v2/audio` or `github.com/faiface/beep`
- **Sound Manager Module**: `internal/audio/` for sound loading and playback
- **Event-Driven System**: Trigger sounds based on game events
- **Configuration**: User-toggleable sound effects in settings

#### **Sound Events to Implement**
```go
type SoundEvent int

const (
    SoundMovePlaced SoundEvent = iota
    SoundGameWon
    SoundGameDraw
    SoundMenuSelect
    SoundMenuNavigate
    SoundError
    SoundStartup
    SoundGradientChange
)
```

#### **Sound Files Needed**
- `move_place.wav` - Soft click when placing X/O
- `game_won.wav` - Victory fanfare
- `game_draw.wav` - Neutral completion sound
- `menu_select.wav` - Menu selection confirmation
- `menu_navigate.wav` - Subtle navigation sound
- `error.wav` - Error notification
- `startup.wav` - Application startup sound
- `gradient_change.wav` - Gradient cycle sound

#### **Implementation Plan**
1. **Audio Module Setup**
   - Create `internal/audio/audio.go`
   - Implement sound loading and caching
   - Add volume control and mute functionality
   - Create sound event dispatcher

2. **Integration Points**
   - Game moves (`makeMove()`, `makeAIMove()`)
   - Menu navigation and selection
   - Settings changes
   - Game state transitions

3. **Configuration Management**
   - Add sound settings to config module
   - Volume slider (0-100%)
   - Master mute toggle
   - Individual sound type toggles

### 🎬 **Game Replay Functionality**

#### **Technical Architecture**
- **Replay Engine**: `internal/replay/` module for recording and playback
- **Replay Format**: JSON-based game recording format
- **Replay Controls**: Play, pause, step forward/backward, speed control
- **Replay UI**: Dedicated replay viewing screen with controls

#### **Data Structures**
```go
type GameReplay struct {
    ID           string          `json:"id"`
    Timestamp    time.Time       `json:"timestamp"`
    GameMode     game.GameMode   `json:"game_mode"`
    AIDifficulty ai.Difficulty   `json:"ai_difficulty,omitempty"`
    Winner       game.Player     `json:"winner"`
    Duration     time.Duration   `json:"duration"`
    Moves        []ReplayMove    `json:"moves"`
    FinalBoard   [3][3]game.Player `json:"final_board"`
}

type ReplayMove struct {
    Player      game.Player   `json:"player"`
    Position    game.Position `json:"position"`
    Timestamp   time.Time     `json:"timestamp"`
    MoveNumber  int           `json:"move_number"`
    TimeElapsed time.Duration `json:"time_elapsed"`
}

type ReplayState struct {
    CurrentMove int           `json:"current_move"`
    Playing     bool          `json:"playing"`
    Speed       float64       `json:"speed"` // 0.5x, 1x, 2x, 4x
    Direction   int           `json:"direction"` // 1 forward, -1 backward
}
```

#### **Replay Controls Interface**
```
┌─ REPLAY CONTROLS ─────────────────────────┐
│ [◀◀] [▶||] [▶▶] [▶▶▶]  Speed: 2.0x       │
│ Move: 5/7  Time: 00:23  Player: O        │
└───────────────────────────────────────────┘

   1   2   3
1  X  |   | O     ◀ Current replay position
  -----------
2     | X |       Move History:
  -----------      1. X -> (0,0)
3     |   |        2. O -> (0,2)
                   3. X -> (1,1)
Current Player: O   4. O -> ...
Status: In Progress 5. X -> ... ◀ Current
                   
Controls:
Space - Play/Pause    ←→ Step moves
+/- Speed control     R Restart
ESC Back to menu      S Save replay
```

#### **Replay Features**
1. **Recording**
   - Auto-record all games
   - Capture timing information
   - Store game metadata
   - Compress and optimize storage

2. **Playback Controls**
   - Play/pause functionality
   - Step forward/backward by move
   - Speed control (0.25x to 4x)
   - Jump to specific move
   - Loop replay option

3. **Replay Management**
   - List saved replays with metadata
   - Search/filter replays by date, mode, winner
   - Delete old replays
   - Export/import replay files
   - Replay statistics and analysis

4. **Enhanced Visualization**
   - Highlight last move during playback
   - Show move timing information
   - Display player thinking time
   - Gradient effects during replay

#### **Integration Points**
1. **Main Menu**: Add "📽️ View Replays" option
2. **Game Over Screen**: Add "📹 Save Replay" and "▶️ Watch Replay" options
3. **Settings**: Replay auto-save preferences
4. **Statistics**: Link replays to game statistics

### 🎯 **Implementation Priority**
1. **Sound Effects** (Medium Priority)
   - Enhances user experience immediately
   - Relatively straightforward to implement
   - Good cross-platform audio library support

2. **Game Replay** (Medium Priority)
   - More complex feature requiring new UI screens
   - Valuable for learning and entertainment
   - Demonstrates advanced state management

### 🔧 **Technical Considerations**
- **Audio Dependencies**: Choose lightweight, cross-platform audio library
- **File I/O**: Efficient replay file format and compression
- **Memory Management**: Avoid memory leaks in audio playback
- **Performance**: Ensure audio doesn't impact game responsiveness
- **Testing**: Audio testing strategies and replay validation