# Planning

## Key Capabilities
Based on the overview, Claude Code can:
- Edit files and fix bugs across codebases
- Answer questions about code architecture
- Execute and fix tests
- Search git history
- Resolve merge conflicts
- Create commits and pull requests
- Browse documentation via web search

## Overview - Key Concepts
- Claude Code is an "agentic coding tool that lives in your terminal"
- Helps developers code faster through natural language commands
- Integrates directly with development environments
- Direct API connection to Anthropic
- No complex server setup required
- Maintains project context awareness
- Performs real coding operations

## System Requirements
- **Operating Systems**: 
  - macOS 10.15+
  - Ubuntu 20.04+/Debian 10+
  - Windows via WSL
- **Hardware**: Minimum 4GB RAM
- **Software**:
  - Node.js 18+
  - git 2.23+ (optional)
  - GitHub or GitLab CLI (optional)
- **Network**: Active internet connection

## Installation
```bash
npm install -g @anthropic-ai/claude-code
```
**Note**: Do NOT use `sudo npm install -g`

### WSL Troubleshooting
If npm/node detection issues occur:
```bash
npm config set os linux
# Or install with:
npm install -g @anthropic-ai/claude-code --force --no-os-check
```

## Authentication Options
1. **Anthropic Console** (default)
   - Requires active billing at console.anthropic.com
2. **Claude App** 
   - Requires Pro or Max plan
3. **Enterprise platforms**
   - Configure with Amazon Bedrock or Google Vertex AI

## Initial Setup
1. Navigate to project directory
2. Run `claude`
3. Complete authentication process
4. Run initial command like `> summarize this project`
5. Generate project guide with `/init`

## Running Claude Code
```bash
# Start interactive mode
claude

# Run a one-time task
claude "task description"

# Run a quick query
claude -p "query"

# Create a Git commit
claude commit
```

## Basic Commands
- `/help` - Show available commands
- `exit` or `Ctrl+C` - Exit Claude Code
- `Tab` - Command completion
- `↑` - Command history

## Common Workflows
- **Understand project**: "what does this project do?"
- **Code changes**: "add a hello world function"
- **Git operations**: "what files have I changed?"
- **Debugging**: "fix the login bug where users see a blank screen"

## Pro Tips
- Be specific with requests
- Break complex tasks into steps
- Use Tab for command completion
- Press ↑ for command history
- Supported shells: Bash, Zsh, Fish

## Getting Help
- **In Claude Code**: Type `/help`
- **Documentation**: Browse guides at docs.anthropic.com
- **Community**: Join [Anthropic Discord](https://www.anthropic.com/discord)
- **Report issues**: https://github.com/anthropics/claude-code/issues

## Recommended Next Steps
- Explore CLI reference
- Customize settings
- Learn advanced techniques
- Review detailed setup guide
- Familiarize yourself with CLI commands and settings

## Additional Commands and Features (from recent search)

### Command Examples
- `claude`: Start interactive mode with REPL session
- `claude "task"`: Run a one-time task
- `claude -p "query"`: Run one-off query (quick mode)
- `claude commit`: Create a Git commit with AI assistance
- `/clear`: Clear conversation history
- `/config`: Theme matching configuration
- `/vim`: Enable vim mode support
- `claude config set`: Configure notifications

### Configuration Options
- Theme matching via `/config`
- Notification setup via `claude config set`
- Vim mode support with `/vim`
- Supported shells: Bash, Zsh, Fish

### New Insights from Documentation
- Claude Code operates as an interactive REPL session
- Provides comprehensive development tools in terminal
- Compatible with enterprise platforms (Amazon Bedrock, Google Vertex AI)
- Claude Code is described as an "AI pair programmer"
- Supports natural language commands for all operations
- Direct API connection ensures security (no intermediate servers)
- Project context awareness helps maintain code consistency

### DevContainer Support
- Claude Code provides DevContainer configuration at https://github.com/anthropics/claude-code/tree/main/.devcontainer
- Useful for consistent development environments

### First-Time Setup Workflow
1. Install Claude Code via npm
2. Navigate to your project directory
3. Run `claude` to start
4. Complete authentication (Console, App, or Enterprise)
5. Use `/init` to generate CLAUDE.md for your project
6. Commit the generated CLAUDE.md file
7. Start with simple commands to explore the codebase

### Key Workflows Demonstrated
- Understanding project structure: "what does this project do?"
- Adding functions: "add a hello world function"
- Git branch management: "create a new feature branch"
- Code refactoring: "refactor this function to be more efficient"
- Writing tests: "write unit tests for this module"
- Updating documentation: "update the README with usage examples"

### WSL-Specific Notes
- Node.js must be installed via Linux package manager (not Windows)
- Use `npm config set os linux` for proper configuration
- Alternative installation: `npm install -g @anthropic-ai/claude-code --force --no-os-check`

## Key Features Deep Dive

### Memory Management
Claude Code uses three types of memory:
1. **Project Memory** (`./CLAUDE.md`): Team-shared instructions for the project
2. **User Memory** (`~/.claude/CLAUDE.md`): Personal preferences across all projects
3. **Import System**: Use `@path/to/import` syntax to include additional files

#### Memory Management Commands
- `/memory`: Direct edit memory
- `#`: Quick add to memory (start input with #)
- `/init`: Initialize project memory

#### Memory Best Practices
- Be specific in instructions
- Use structured markdown with headings and bullet points
- Document frequently used commands and code style preferences
- Include architectural patterns and design decisions
- Keep memories concise and clear
- Maximum import depth: 5 hops

### Extended Thinking Feature
- **Basic thinking**: Triggered by "think"
- **Deeper thinking**: Triggered by "think more", "think harder"
- Most valuable for: architectural changes, debugging, feature planning
- Thinking process displayed as italic gray text above responses

### Image Workflows
Three methods to add images:
1. **Drag and drop** into Claude Code window
2. **Copy and paste** with Ctrl+V
3. **Provide image path**: "/path/to/your/image.png"

### Resume and Continue Features
- `--continue`: Automatically resumes most recent conversation
- `--resume`: Interactive conversation picker for multiple sessions

### IDE Integrations

#### Supported IDEs
- **Visual Studio Code** (including Cursor and Windsurf forks)
- **JetBrains IDEs**: PyCharm, WebStorm, IntelliJ, GoLand

#### Integration Features
- **Quick launch**: `Cmd+Esc` (Mac) or `Ctrl+Esc` (Windows/Linux)
- **Diff viewing** in IDE
- **Automatic context sharing** for selected text/tabs
- **File reference shortcuts**: `Cmd+Option+K` (Mac) or `Alt+Ctrl+K` (Linux/Windows)
- **Automatic diagnostic error sharing**

#### Setup Instructions
- **VS Code**: Open terminal → Run `claude` (auto-installs extension)
- **JetBrains**: Install Claude Code plugin from marketplace → Restart IDE
- **Configuration**: Use `/config` command, set diff tool to `auto`

## Advanced Features Deep Dive

### Model Context Protocol (MCP)
MCP enables LLMs to access external tools and data sources through an open protocol.

#### MCP Server Types
1. **Stdio Servers**: Direct process communication
2. **SSE Servers**: Server-Sent Events
3. **HTTP Servers**: REST API communication

#### MCP Server Scopes
- **Local scope**: Project-specific, private servers
- **Project scope**: Shared team configurations  
- **User scope**: Cross-project servers

#### MCP Usage Examples
- Database queries: "describe the schema of our users table"
- Resource references: "@github:issue://123", "@docs:file://api/authentication"
- Slash commands: "/mcp__github__list_prs", "/mcp__jira__create_issue"

#### MCP Setup
```bash
claude mcp add my-server -e API_KEY=123 -- /path/to/server arg1 arg2
```

**Security Warning**: Use third-party MCP servers at your own risk. Ensure you trust the servers.

### GitHub Actions Integration

#### Setup Requirements
- Install Claude GitHub App
- Add `ANTHROPIC_API_KEY` to repository secrets
- Copy workflow file to `.github/workflows/`

#### Use Cases
- Turn issues into PRs automatically
- Get implementation help for features
- Fix bugs quickly with AI assistance
- Automated code reviews
- Custom code implementations

#### Security Best Practices
- Never commit API keys directly to repository
- Use GitHub Secrets for sensitive data
- Limit action permissions appropriately
- Review AI suggestions before merging

### Claude Code SDK

#### Available SDKs
- **TypeScript**: `npm install @anthropic-ai/claude-code`
- **Python**: `pip install claude-code-sdk`

#### SDK Capabilities
- Run Claude Code as subprocess
- Build AI-powered coding assistants
- Multi-turn conversation support
- Custom system prompts
- Non-interactive and interactive modes
- MCP integration for extending tools

#### Key SDK Methods
- `query()`: Primary interaction method
- Supports streaming and non-streaming responses
- Configurable with max turns, system prompts
- Resume conversation capability
- JSON and text output formats

#### Authentication Options
- Anthropic API key
- Amazon Bedrock
- Google Vertex AI

#### Unique SDK Features
- Explicit tool permission management
- Custom permission prompt tools
- Flexible input/output formats
- Programmatic code generation and analysis

### Git Worktrees for Parallel Sessions
Use Git worktrees to run multiple Claude Code sessions simultaneously on different branches without conflicts.

### Unix-Style Utility Integration
Claude Code integrates with standard Unix utilities and supports various output formats for scripting and automation.

## Reference and Support Deep Dive

### Extended Thinking (Advanced)
Extended thinking allows Claude to work through complex problems step-by-step, improving performance on difficult tasks.

#### Key Characteristics
- **Minimum thinking budget**: 1024 tokens
- **Language**: Performs best in English, though final outputs can be in any supported language
- **Use cases**: Complex STEM problems, constraint optimization, structured thinking frameworks

#### Best Practices for Extended Thinking
1. **Start with high-level, general instructions**
2. **Allow Claude creative problem-solving space**
3. **Use multishot prompting with XML tags**
4. **Break complex instructions into numbered steps**
5. **Explicitly ask Claude to verify and test its work**

#### Specific Triggers
- "think thoroughly"
- "Consider multiple approaches" 
- "Show complete reasoning"
- Explicit request for step-by-step analysis

#### Performance Optimization
- Start with minimum thinking budget (1024 tokens)
- Incrementally increase based on task complexity
- For workloads over 32K tokens, use batch processing

#### Important Limitations
- Avoid manually changing Claude's thinking output
- Do not pass thinking text back as user input
- Prefilling extended thinking is not allowed
- Don't push Claude to output more tokens unnecessarily

### CLI Reference

#### Core CLI Commands
```bash
# Interactive modes
claude                           # Start interactive REPL
claude "query"                   # Start REPL with initial prompt
claude -p "query"               # Query via SDK, then exit
cat file | claude -p "query"    # Process piped content

# Session management
claude -c                       # Continue most recent conversation  
claude -r "<session-id>" "query" # Resume session by ID

# Utility commands
claude update                   # Update to latest version
claude mcp                      # Configure Model Context Protocol servers
```

#### Key CLI Flags
- `--add-dir`: Add working directories
- `--allowedTools`: Specify allowed tools
- `--print/-p`: Print response without interactive mode
- `--output-format`: Specify output format
- `--verbose`: Enable detailed logging
- `--model`: Set model for session
- `--permission-mode`: Set permission mode
- `--continue`: Load most recent conversation

### Slash Commands System

#### Built-in Slash Commands
```bash
# Core functionality
/add-dir      # Add additional working directories
/clear        # Clear conversation history
/help         # Get usage help
/init         # Initialize project with CLAUDE.md guide

# Configuration and management
/config       # View/modify configuration
/memory       # Edit CLAUDE.md memory files
/model        # Select or change AI model
/permissions  # View/update permissions

# Development tools
/review       # Request code review
/pr_comments  # View pull request comments
/mcp          # Manage MCP server connections

# System utilities
/cost         # Show token usage statistics
/doctor       # Check Claude Code installation health
/status       # View account and system statuses
/terminal-setup # Install key binding for newlines
/vim          # Enter vim mode

# Account management
/login        # Switch Anthropic accounts
/logout       # Sign out of Anthropic account
/bug          # Report bugs to Anthropic

# Advanced features
/compact      # Compact conversation with optional focus
```

#### Custom Slash Commands
- **Location**: Project (`.claude/commands/`) or Personal (`~/.claude/commands/`)
- **Syntax**: `/<prefix>:<command-name> [arguments]`
- **Prefixes**: `project` (project-specific), `user` (personal)
- **Features**: Supports bash execution, file references, arguments via `$ARGUMENTS`

#### MCP Slash Commands
- **Format**: `/mcp__<server-name>__<prompt-name> [arguments]`
- **Discovery**: Dynamically discovered from connected MCP servers

### Interactive Mode

#### Keyboard Shortcuts
```bash
# General controls
Ctrl+C        # Cancel current input/generation
Ctrl+D        # Exit Claude Code session
Ctrl+L        # Clear terminal screen
↑/↓           # Navigate command history
Esc + Esc     # Edit previous message

# Multiline input
\             # Quick escape (all terminals)
Option+Enter  # macOS default
Shift+Enter   # After /terminal-setup
```

#### Vim Mode Navigation
```bash
# Movement
h/j/k/l       # left/down/up/right

# Word navigation  
w             # next word
e             # end of word
b             # previous word

# Line navigation
0             # start of line
$             # end of line
^             # first non-blank character
```

#### Session Management
- **Command history**: Stored per working directory
- **History clearing**: Use `/clear` command
- **Reverse search**: `Ctrl+R` for history search

#### Quick Commands
- `#` at start: Memory shortcut
- `/` at start: Slash command

### Settings and Configuration

#### Configuration File Locations
```bash
# User settings
~/.claude/settings.json

# Project settings (shared)
.claude/settings.json

# Project settings (personal)  
.claude/settings.local.json

# Enterprise settings
# macOS: /Library/Application Support/ClaudeCode/managed-settings.json
# Linux/Windows: /etc/claude-code/managed-settings.json
```

#### Key Settings Options
- `apiKeyHelper`: Custom script for authentication
- `cleanupPeriodDays`: Chat transcript retention period
- `env`: Environment variables for sessions
- `includeCoAuthoredBy`: Toggle Claude byline in commits
- `permissions`: Control tool access rules

#### Permission Configuration
```json
{
  "permissions": {
    "allow": ["ToolName(pattern:*)"],
    "deny": ["DangerousTool(*)"],
    "additionalDirectories": ["/safe/path"],
    "defaultMode": "allow"
  }
}
```

#### Environment Variables
- `ANTHROPIC_API_KEY`: API authentication
- `CLAUDE_CODE_USE_BEDROCK`: Enable AWS Bedrock
- `CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC`: Disable non-critical network
- `DISABLE_TELEMETRY`: Opt-out of usage tracking

#### Available Tools for Claude
- Read/Write file operations
- Bash command execution
- Web searching
- Code editing
- Notebook manipulation

### Hooks System

#### Hook Types and Triggers
- **PreToolUse**: Runs before tool calls
- **PostToolUse**: Runs after successful tool completion
- **Notification**: Triggered on Claude Code notifications
- **Stop**: Runs when main agent finishes responding
- **SubagentStop**: Runs when subagent (Task) finishes

#### Hook Configuration Structure
```json
{
  "hooks": {
    "EventName": [
      {
        "matcher": "ToolPattern",
        "hooks": [
          {
            "type": "command",
            "command": "your-command-here"
          }
        ]
      }
    ]
  }
}
```

#### Example Use Cases
1. **Logging**: Record shell commands
2. **Formatting**: Automatic code formatting
3. **Notifications**: Custom alerts
4. **Permissions**: Custom access control
5. **Feedback**: Code convention enforcement

#### Security Best Practices for Hooks
- Hooks execute with full user permissions
- Always validate and sanitize inputs
- Use absolute paths for scripts
- Avoid sensitive file access
- Test in safe environments first
- Quote shell variables properly
- Block path traversal attempts

#### Hook Features
- Supports both exit code and JSON output
- Provides granular control over tool execution
- Can block or approve specific actions
- Integrates with MCP tools
- Configurable matching for tool interactions