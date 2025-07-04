# Claude Code Learning Workflow

This document outlines the systematic workflow for learning and documenting Claude Code functionality, including file management, URL tracking, and documentation processes.

## Overview

The workflow consists of multiple interconnected files that work together to create a comprehensive learning and documentation system for Claude Code.

## File Structure and Interactions

```
├── workflow.md                    # This file - workflow documentation
├── planning.md                    # Core findings and learning content
├── claude-code-links.md          # Active/unvisited documentation links
├── visited-urls.md               # Tracking file for searched URLs
├── questions-for-objectives.md   # Question management system
├── research.md                   # User profile and insights
├── todo.md                       # Learning plan and milestones
└── CLAUDE.md                     # Project memory for Claude Code
```

## Core Workflow Process

### 1. URL Research and Documentation Workflow

#### Step 1: URL Discovery
When searching documentation online:
- **Action**: Use WebFetch to search URLs
- **Requirement**: Extract ALL URLs found in the documentation
- **Output**: List of new URLs discovered

#### Step 2: URL Management
- **Add New URLs**: Place discovered URLs in `claude-code-links.md` (categorized)
- **Remove Searched URLs**: Move searched URLs from `claude-code-links.md` to `visited-urls.md`
- **Avoid Duplicates**: Check `visited-urls.md` before adding to `claude-code-links.md`
- **Rule**: Never store URLs more than once across the system

#### Step 3: Content Integration
- **Primary Target**: Update `planning.md` with all findings
- **Secondary Updates**: Update other relevant files based on content discovered

### 2. Question Management Workflow

#### File Structure: `questions-for-objectives.md`
```markdown
# Questions to Understand Your Claude Code Learning Objectives

## Questions
1. [New unanswered questions in numbered list]
2. [Additional follow-up questions]

## Answered Questions
1. [Original question text]
   - (A) [User's answer indented on next line]
2. [Next answered question]
   - (A) [User's answer indented on next line]
```

#### Process Flow
1. **New Questions**: Add to "Questions" section with numbering
2. **When Answered**: Move from "Questions" to "Answered Questions"
3. **Format Answers**: Use `- (A)` prefix and indent on next line
4. **Research Integration**: Extract insights to update `research.md`
5. **Follow-up Generation**: Create new questions based on answers

### 3. Research and Insights Workflow

#### Target File: `research.md`
- **Purpose**: Store user profile, preferences, and learning insights
- **Source**: Answers from `questions-for-objectives.md`
- **Content**: User profile summary, pain points, learning preferences, technical interests
- **Updates**: Add new insights whenever questions are answered

#### Integration Points
- Influences todo list priorities in `todo.md`
- Guides focus areas in `planning.md`
- Helps determine which URLs to prioritize

### 4. Planning Documentation Workflow

#### File Purpose: `planning.md`
Central repository for all Claude Code knowledge and findings.

#### Content Organization
1. **Key Capabilities**: Core Claude Code functions
2. **System Requirements**: Technical prerequisites  
3. **Installation**: Setup instructions
4. **Authentication**: Login methods
5. **Commands**: Basic and advanced command examples
6. **Feature Deep Dives**: Detailed explanations of major features
7. **Best Practices**: Tips and recommendations

#### Update Process
- **Always Update**: When searching any documentation URL
- **Content Type**: Comprehensive details, examples, commands, warnings
- **Structure**: Maintain consistent markdown hierarchy
- **Integration**: Cross-reference with other files when relevant

### 5. Link Management System

#### Active Links File: `claude-code-links.md`
**Purpose**: Store unvisited/unsearched URLs only
**Categories**:
- Core Documentation Links
- External Links  
- Development Tools
- Claude Code Resources
- GitHub Integration Resources
- Cloud Provider Integration
- Additional Links

**Management Rules**:
- Only contains URLs not yet searched
- Remove URLs immediately after searching
- Add new URLs discovered during searches
- Maintain categorical organization

#### Visited Links File: `visited-urls.md`
**Purpose**: Track all searched URLs to prevent duplicates
**Categories**:
- Claude Code Documentation URLs
- External Links
- Development Tools
- [Additional categories as needed]

**Management Rules**:
- Add URLs immediately after searching
- Never remove URLs (permanent record)
- Check before adding new URLs to `claude-code-links.md`
- Maintain chronological or categorical order

### 6. Learning Plan Workflow

#### Target File: `todo.md`
**Purpose**: Structured learning plan based on user objectives
**Content**: Phase-based learning with specific tasks and milestones
**Updates**: Based on `research.md` insights and `planning.md` discoveries

## Workflow Execution Instructions

### When Searching Documentation URLs

1. **Prepare for Search**
   ```
   TodoWrite: Create tasks for URL searches
   ```

2. **Execute Search**
   ```
   WebFetch: Search each URL with comprehensive prompts
   Extract: All URLs, features, examples, commands, warnings
   ```

3. **Update Files (in order)**
   ```
   a. visited-urls.md: Add searched URLs
   b. claude-code-links.md: Remove searched URLs, add discovered URLs
   c. planning.md: Add comprehensive findings
   d. Other files: As relevant to content
   ```

4. **Complete Tasks**
   ```
   TodoWrite: Mark tasks as completed
   ```

### When Receiving Question Answers

1. **Update Questions File**
   ```
   Move answered questions to "Answered Questions" section
   Format with (A) prefix and indentation
   ```

2. **Extract Insights**
   ```
   Analyze answers for user preferences, pain points, goals
   Update research.md with new insights
   ```

3. **Generate Follow-ups**
   ```
   Create new questions based on answers
   Add to "Questions" section
   ```

4. **Update Learning Plan**
   ```
   Modify todo.md based on new insights
   Prioritize relevant learning paths
   ```

## File Interaction Rules

### Dependencies
- `research.md` ← `questions-for-objectives.md` (answers drive insights)
- `todo.md` ← `research.md` (insights drive learning plan)
- `planning.md` ← All URL searches (findings accumulate)
- `visited-urls.md` ← All searches (permanent tracking)
- `claude-code-links.md` ← All searches (dynamic URL management)

### Update Triggers
- **URL Search**: Always update `planning.md`, `visited-urls.md`, `claude-code-links.md`
- **Question Answered**: Always update `research.md`, `questions-for-objectives.md`
- **New Insights**: Update `todo.md` if learning priorities change
- **New URLs Discovered**: Always add to `claude-code-links.md` if not duplicate

### Quality Standards
- **Comprehensive Documentation**: Don't summarize; include full details
- **Consistent Formatting**: Maintain markdown structure across files
- **No Duplicates**: Strict URL deduplication across the system
- **Cross-references**: Link related information between files
- **Version Control**: Track changes and updates systematically

## Success Metrics

### Workflow Effectiveness
- [ ] All discovered URLs properly categorized and tracked
- [ ] No duplicate URLs across the system
- [ ] Comprehensive documentation accumulation in `planning.md`
- [ ] User insights driving personalized learning in `todo.md`
- [ ] Questions leading to actionable research updates

### Documentation Quality
- [ ] Each URL search produces significant `planning.md` updates
- [ ] New features documented with examples and commands
- [ ] User-specific learning path maintained and updated
- [ ] Cross-file relationships maintained consistently

This workflow ensures systematic learning progression while maintaining comprehensive documentation and avoiding duplicate work.