# Tic-Tac-Toe Project Q&A

## Questions
1. What's your familiarity with the Bubble Tea framework? Should I include extra explanations?
2. Do you want any additional features like score tracking, computer AI opponent, different board sizes?
3. Any specific styling preferences for the terminal interface (colors, symbols, layout)?
4. Which aspects are most important for your learning: Bubble Tea framework patterns, Go project structure, game logic implementation, or terminal UI design?
5. Should this be a minimal implementation to understand basics, or a more feature-rich version?

## Answered Questions
1. What's your familiarity with the Bubble Tea framework? Should I include extra explanations?
   - (A) None, I need a summary of what it can do

2. Do you want any additional features like score tracking, computer AI opponent, different board sizes?
   - (A) Yes let's add score tracking, and make sure that the game is persistent. There needs to be a way to quit the game as well. It should be 2 players. Yes let's create a simple AI in the game to play against if single player is selected. Let's keep the game at 3 x 3 for now, but we might change this in the future.

3. Any specific styling preferences for the terminal interface (colors, symbols, layout)?
   - (A) I like lots of rainbow colors so having some gradient like linear gradient would be awesome and if we can add animation so that the colors move down the displayed terminal. If you have questions for this, please ask me more about it. I can give more detailed answers.

4. Which aspects are most important for your learning: Bubble Tea framework patterns, Go project structure, game logic implementation, or terminal UI design?
   - (A) Sure, bubble tea

5. Should this be a minimal implementation to understand basics, or a more feature-rich version?
   - (A) Let's go for minimal implementation right now and create an MVP. We will definitely update this as we go on, but we're starting out with a great set of planning sessions and todos.

## Follow-up Questions

## Answered Follow-up Questions
1. For the rainbow gradient animation - do you want the colors to cycle through the entire rainbow spectrum, or focus on specific color combinations?
   - (A) Yes, the rainbow spectrum will be nice. Actually, let's add an option feature to the application where you can select linear gradient for red, orange, yellow, green, blue, indigo, violet, and rainbow where the default setting will be rainbow

2. Should the gradient animation be constant/continuous, or triggered by specific events (like moves, wins)?
   - (A) It should be continuous

3. For persistence - should we save game state to a file, or just maintain session-based persistence?
   - (A) Sure, let's add to do for preserving the game state

4. For the AI opponent - should it be random moves, or implement basic strategy (blocking wins, taking center, etc.)?
   - (A) Let's implement a basic strategy nothing complex, but let's also implement strategies for easy, normal hard, and I never lose. These will be the difficulty settings for the AI.

## New Requirements Questions

## Answered New Requirements
1. How should the application handle input and controls?
   - (A) I want to focus on the system and how input will be registered into the application, I want the ability to use keybindings. I want the keybindings to display on the terminal with the action it performs. I also want to have the ability to select a square on the tic-tac-toe with your mouse click.

2. What should the application startup experience be?
   - (A) I want the application to start up with a start screen, displaying a small graphic in eight bit picture showing the tic-tac-toe board somewhat filled out. Add the selected color scheme to this as well.

3. How frequently should the application save data?
   - (A) We should store saved files on every operation.

## Additional UI Questions

## Answered Additional UI Questions
1. Should the keybindings show as a permanent panel, or toggle on/off?
   - (A) Permanent panel

2. What visual feedback do you want when hovering/clicking squares?
   - (A) Clicking on the squares, add a little flare just make the square pop a little, be creative about it, use your best judgement.

3. Any specific retro aesthetic preferences for 8-bit graphics?
   - (A) Retro color palettes please.

## Development Planning Questions

## Answered Development Planning Questions
1. What testing framework should we use for the Go application?
   - (A) I want to create tests for the application, and we should always use ginkgo and gomega when testing go for every function that we create, we need to be able to unit test it.