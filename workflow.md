# Claude Code Learning Workflow

This document defines the precise workflow for learning Claude Code, managing projects, and maintaining documentation. Follow these procedures exactly as specified.

## Command Reference

### `#organize-wk` - Organize Workflow
When this command is issued:
1. Review current project structure and files
2. Ensure all required files exist (todo.md, plan.md, go.mod for Go projects)
3. Clean up any duplicate main() functions by converting to utility functions
4. Verify project runs with `go run .` (for Go projects)
5. Update todo.md with completion status
6. Document any changes made

## 1. Project Structure Requirements

### Every Project MUST Have:
```
project-name/
├── todo.md      # Task tracking with [ ] and [✅] format
├── plan.md      # Project implementation plan
├── QA.md        # Questions and answers for project planning
├── main.go      # Single entry point (Go projects)
├── go.mod       # Module definition (Go projects)
└── [other files as needed]
```

### File Creation Order:
1. Create project directory
2. Create `todo.md` with initial tasks
3. Create `plan.md` with implementation approach
4. Create `QA.md` with planning questions
5. Create language-specific files (go.mod, main.go, etc.)
6. Test basic functionality before adding features

## 2. Todo.md Management

### Format Requirements:
```markdown
## Task Category
- [ ] Incomplete task description
- [✅] Completed task description
```

### Task Management Rules:
1. **ALWAYS** add tasks to todo.md BEFORE starting work
2. **NEVER** batch update completed tasks - mark each one immediately
3. **EVERY** suggestion from Claude must be added as a task
4. **BREAK DOWN** complex tasks into smaller, trackable steps
5. **UPDATE** immediately when a task is completed with [✅]

### Example Todo.md Structure:
```markdown
## Project Setup
- [✅] Create project directory
- [✅] Initialize go.mod
- [ ] Create main.go structure
- [ ] Add basic functionality
- [ ] Test with go run .

## Bug Fixes
- [ ] Fix duplicate main() function error
- [ ] Remove main() from utility files
- [ ] Update imports in main.go
```

## 3. Go Project Organization

### Main Function Rule:
- **ONE** main() function per project
- **ONLY** in main.go file
- **NEVER** in utility or library files

### Module Organization Rule:
- **ALWAYS** organize code into logical modules for programming projects
- **CREATE** separate directories for different concerns
- **MAINTAIN** clear separation of responsibilities
- **USE** Go package structure for organization

### Recommended Module Structure:
```
project-name/
├── main.go           # Entry point only
├── go.mod           # Module definition
├── internal/        # Private packages
│   ├── game/        # Core game logic
│   ├── ui/          # User interface components
│   ├── ai/          # AI system
│   ├── persistence/ # Save/load functionality
│   └── config/      # Configuration management
├── pkg/             # Public packages (if any)
└── assets/          # Static files, graphics, etc.
```

### File Organization:
```go
// main.go - ONLY file with main()
package main

import (
    "project-name/internal/game"
    "project-name/internal/ui"
)

func main() {
    // Orchestrate program functionality
    game := game.NewGame()
    ui.Start(game)
}

// internal/game/game.go - Game logic module
package game

type Game struct { /* ... */ }
func NewGame() *Game { /* ... */ }

// internal/ui/ui.go - UI module
package ui

func Start(game *game.Game) { /* ... */ }
```

### Compilation Error Resolution:
1. **IDENTIFY** the error (usually duplicate main())
2. **PLAN** the fix in todo.md
3. **EDIT** files to remove duplicate main() - convert to functions
4. **TEST** with `go run .` after each change
5. **NEVER** delete files to fix errors

## 4. Problem Resolution Workflow

### When Errors Occur:
1. **READ** the error message completely
2. **ADD** fix tasks to todo.md
3. **PLAN** the solution approach
4. **EXECUTE** fixes one at a time
5. **TEST** after each change
6. **MARK** completed fixes with [✅]

### File Editing Principles:
- **PRESERVE** all code by converting, not deleting
- **EDIT** conflicting functions to utilities
- **MAINTAIN** all struct definitions and types
- **TEST** incrementally with `go run .`

### Example Fix Process:
```markdown
## Fix Multiple main() Functions
- [ ] Identify which files have main()
- [ ] Convert hello.go main() to PrintHello()
- [ ] Convert person.go main() to DemonstratePerson()
- [ ] Update main.go to call new functions
- [ ] Test with go run .
- [ ] Verify output is correct
```

## 5. Documentation Workflow

### URL Research Process:
1. **SEARCH** URL with WebFetch
2. **EXTRACT** all findings and new URLs
3. **UPDATE** planning.md with comprehensive details
4. **MOVE** searched URLs to visited-urls.md
5. **ADD** new URLs to claude-code-links.md
6. **REMOVE** searched URLs from claude-code-links.md

### File Update Order:
1. `visited-urls.md` - Add searched URLs
2. `claude-code-links.md` - Remove searched, add discovered
3. `planning.md` - Add comprehensive findings
4. `research.md` - Update with insights (if applicable)
5. `todo.md` - Update task completion status

## 6. Question Management (QA.md File)

### QA.md File Requirements:
- **EVERY** project MUST have a `QA.md` file
- Created during project planning phase
- Contains all project-related questions and answers

### Question Workflow:
1. **CREATE** QA.md file with initial planning questions
2. **ADD** new questions to "Questions" section
3. **MOVE** answered questions to "Answered Questions"
4. **FORMAT** answers with `- (A)` prefix, indented
5. **UPDATE** plan.md with insights from answers
6. **CREATE** follow-up questions based on answers

### QA.md Format Example:
```markdown
# Project Name Q&A

## Questions
1. What is your experience with Go?
2. What features do you want in the application?

## Answered Questions
1. What is your experience with Go?
   - (A) I have 2 years of experience with Go development

## Follow-up Questions
1. Should we implement advanced Go patterns?
```

## 7. Learning Path Management

### Phase Progression:
1. **COMPLETE** Phase 1 Foundation before moving to Phase 2
2. **MARK** all tasks in current phase before advancing
3. **DOCUMENT** learnings from each phase
4. **UPDATE** future phases based on learnings

### Success Metrics:
- All tasks marked with [✅] for completed phase
- Project runs successfully with `go run .`
- Can explain concepts learned to others
- Ready to apply knowledge in next phase

## 8. File Management Standards

### Creation Rules:
- **CREATE** files in proper order (todo.md first)
- **INITIALIZE** with proper structure
- **TEST** basic functionality immediately
- **DOCUMENT** purpose in comments or plan.md

### Editing Rules:
- **NEVER** delete files to fix problems
- **ALWAYS** preserve code by converting
- **EDIT** to resolve conflicts
- **TEST** after every change

### Deletion Rules:
- **ONLY** delete if file is completely unused
- **NEVER** delete to fix compilation errors
- **CONFIRM** no references exist before deletion
- **DOCUMENT** reason for deletion in todo.md

## 9. Progress Tracking

### Daily Workflow:
1. **REVIEW** todo.md at start
2. **SELECT** tasks to complete
3. **EXECUTE** tasks one by one
4. **MARK** completed with [✅] immediately
5. **COMMIT** changes (if using git)
6. **PLAN** next day's tasks

### Weekly Review:
1. **COUNT** completed vs incomplete tasks
2. **IDENTIFY** blockers or challenges
3. **UPDATE** learning plan if needed
4. **CELEBRATE** progress made
5. **PLAN** next week's objectives

## 10. Quality Assurance

### Before Marking Task Complete:
- Code compiles without errors
- Functionality works as expected
- Tests pass (if applicable)
- Documentation is updated
- Todo.md reflects completion

### Project Completion Checklist:
- [ ] All todo.md tasks marked complete
- [ ] Project runs with single command
- [ ] No duplicate main() functions
- [ ] All files properly organized
- [ ] Documentation accurate and complete

## Summary

This workflow ensures consistent, systematic progress in learning Claude Code. Follow each step precisely, maintain accurate documentation, and track all progress in todo.md files. Use `#organize-wk` command to trigger workflow organization when needed.