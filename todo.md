# Claude Code Learning TODO List

## Phase 1: Foundation (Week 1-2)
- [ ] Read all Getting Started documentation (Overview, Setup, Quickstart)
- [✅] Install Claude Code and complete authentication
- [ ] Run first "Hello World" command in Claude Code
- [✅] Complete `/init` command on a sample Golang project
- [✅] Learn all basic commands (`/help`, navigation, exit)
- [✅] Practice `/config` command
- [✅] Practice `/memory` command
- [ ] Practice basic file editing with Claude Code
- [ ] Explore CLI reference documentation

## Phase 1: Immediate Learning Tasks
- [✅] Create a simple Go hello world program with Claude
- [✅] Generate a basic Go struct for a Person with name, age, and email fields
- [✅] Ask Claude to explain current project files
- [ ] Create tic-tac-toe Bubble Tea project in separate folder
- [ ] Review and understand the tic-tac-toe project structure
- [ ] Implement the tic-tac-toe game following the plan.md

## Simple Program Organization Tasks
- [✅] Create go.mod file for the simple program
- [✅] Refactor code to have only one main() function in main.go
- [✅] Move Person struct to separate file or integrate into main program
- [✅] Test program with `go run .` command
- [✅] Fix any compilation errors that occur
- [✅] Ensure program runs successfully without errors
- [✅] Document the final program structure

## Fix Multiple main() Functions Issue - Detailed Plan
- [✅] Remove main() function from hello.go, keep only helper functions if any
- [✅] Remove main() function from person.go, keep Person struct and methods
- [✅] Update main.go to import and use functionality from other files
- [✅] Test each step with `go run .` to ensure no compilation errors
- [✅] Verify program runs with complete functionality from all files
- [✅] Update todo.md to mark each step as completed

## Phase 2: Core Golang Development (Week 2-3)
- [ ] Create a simple Golang calculator with Claude Code (simplified from financial)
- [ ] Use Claude Code to generate Go structs and interfaces
- [ ] Practice refactoring existing Go code with Claude
- [ ] Learn to use Claude for Go testing and debugging
- [ ] Explore how Claude handles Go modules and dependencies
- [ ] Generate API schemas (OpenAPI/Swagger) with Claude

## Phase 3: Research and Exploration Features (Week 3-4)
- [ ] Learn to use Claude Code for code exploration in large codebases
- [ ] Practice asking Claude to explain complex code sections
- [ ] Use Claude to research Go best practices
- [ ] Explore alternative implementations for common patterns
- [ ] Learn effective prompting to get Claude to ask helpful questions

## Phase 4: Multi-Language Development (Week 4-5)
- [ ] Create a web interface using JavaScript/HTML with Claude
- [ ] Start learning TypeScript basics through Claude Code
- [ ] Build a Python data analysis script with Claude's help
- [ ] Practice switching between languages in a single project
- [ ] Create a full-stack financial dashboard (Go backend, TS frontend)

## Phase 5: MCP (Model Context Protocol) Learning (Week 5-8)

### Week 1: MCP Fundamentals
- [ ] Research MCP documentation and understand basic concepts
- [ ] Learn what "data access" and "workflow automation" mean in MCP context
- [ ] Understand the three MCP server types: Stdio, HTTP, SSE
- [ ] Set up MCP development environment
- [ ] Install and test first existing MCP server
- [ ] Practice basic MCP slash commands in Claude Code

### Week 2: Practical MCP Usage  
- [ ] Set up file system MCP server for local file access
- [ ] Use Claude to analyze codebase through MCP file access
- [ ] Install and configure git MCP server
- [ ] Practice git operations through Claude using MCP
- [ ] Learn MCP resource references (@file://, @git:// syntax)
- [ ] Document MCP setup and usage patterns

### Week 3: Building Custom MCP Tools
- [ ] Create first custom stdio MCP server in Go
- [ ] Implement basic file operations through custom MCP server
- [ ] Build Golang-specific development helpers via MCP
- [ ] Add code formatting and linting integration to MCP server
- [ ] Test and debug custom MCP server integration
- [ ] Create documentation for custom MCP tools

### Week 4: Real-World MCP Applications  
- [ ] Build project automation tools using MCP
- [ ] Create development workflow automation via MCP
- [ ] Implement error handling and logging in MCP server
- [ ] Add configuration management to MCP tools
- [ ] Create usage examples and tutorials for MCP setup
- [ ] Plan ongoing MCP tool expansion and maintenance

## Phase 6: Advanced Features (Week 9-10)
- [ ] Explore Claude Code SDK documentation
- [ ] Build a custom integration using the SDK
- [ ] Learn all available slash commands
- [ ] Master Claude Code settings and configuration
- [ ] Explore real-time data streaming with SSE servers
- [ ] Create comprehensive MCP + SDK workflow

## Phase 7: Specialized Applications (Week 11+)
- [ ] Build a simple project management system with Claude
- [ ] Create development productivity tools
- [ ] Implement project templates and scaffolding
- [ ] Develop code analysis and metrics tools
- [ ] Create comprehensive documentation with Claude

## Ongoing Learning Tasks
- [ ] Document effective prompts for common tasks
- [ ] Create a personal Claude Code cheat sheet
- [ ] Build a library of reusable prompts for development
- [ ] Join Anthropic Discord and engage with community
- [ ] Track and document any issues for GitHub feedback
- [ ] Maintain and expand custom MCP servers
- [ ] Create tutorials for MCP setup and usage

## Questions to Answer Through Practice
- [ ] What are Claude Code's limitations with Golang?
- [ ] How effective is Claude at generating domain-specific code?
- [ ] What prompting strategies work best for refactoring?
- [ ] How can Claude best help with schema generation?
- [ ] What questions does Claude ask that are most helpful?
- [ ] What MCP tools provide the most development value?
- [ ] How can MCP servers best integrate with existing workflows?

## MCP-Specific Success Metrics
- [ ] Can explain MCP concepts to others clearly
- [ ] Can set up existing MCP servers independently
- [ ] Can build custom MCP servers for specific needs
- [ ] Can troubleshoot MCP connection and configuration issues
- [ ] Can integrate MCP tools into daily development workflow
- [ ] Understand when to use MCP vs built-in Claude Code features

## General Success Metrics
- [ ] Can generate complete Go modules without manual fixes
- [ ] Can refactor code 5x faster than manual approach
- [ ] Can explore new codebases efficiently with Claude
- [ ] Understand when to use Claude vs manual coding
- [ ] Master all Claude Code features documented
- [ ] Successfully implement MCP-enhanced development workflow