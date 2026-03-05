# Normalize — GitHub Project Issues

Use this document to create GitHub Milestones and Issues. Each **Milestone** section below maps to a GitHub Milestone. Each **Issue** within has a title, description, and acceptance criteria ready to paste into GitHub Issues.

---

## Milestone 1: Core Copy

**Description:** Backup and restore a single file by name using TOML config-driven path resolution, with Lip Gloss styled output from the start.
**Target:** 1–2 weekends

---

### Issue 1.1: Project Scaffolding

**Labels:** `setup`, `phase-1`

**Description:**
Initialize the Go module, set up the project directory structure, install dependencies (Cobra, BurntSushi/toml, Lip Gloss, Charm Log), and create a Makefile with build targets for both Linux (amd64) and macOS (arm64).

**Acceptance Criteria:**
- [ ] `go mod init github.com/<user>/normalize` runs successfully
- [ ] Directory structure matches the design doc (`cmd/`, `internal/config/`, `internal/engine/`, `internal/resolver/`, `internal/ui/`)
- [ ] `go build ./...` compiles with no errors
- [ ] Makefile has `build-linux`, `build-macos`, and `build` (current OS) targets
- [ ] Cobra root command exists and prints help text when run with no args
- [ ] `.gitignore` covers Go binaries and OS artifacts
- [ ] Dependencies installed: cobra, bubbletea, lipgloss, bubbles, charmbracelet/log, BurntSushi/toml

---

### Issue 1.2: Lip Gloss Theme & Output Helpers

**Labels:** `ui`, `phase-1`

**Description:**
Create the shared Lip Gloss theme in `internal/ui/theme.go` and styled output helpers in `internal/ui/output.go`. Define a consistent color palette and reusable styles for success, error, warning, info, and dimmed text. Set up Charm Log as the structured logger for verbose/debug output.

**Acceptance Criteria:**
- [ ] `theme.go` defines a `Theme` struct with Lip Gloss styles for: success (green), error (red), warning (yellow), info (blue), dimmed (gray), bold, and header
- [ ] `output.go` exposes helper functions: `Success(msg)`, `Error(msg)`, `Warn(msg)`, `Info(msg)` that use the theme
- [ ] Each output helper prints a styled prefix icon (e.g., ✓, ✗, ⚠, ℹ) followed by the message
- [ ] Charm Log is configured with styled level output (debug, info, warn, error)
- [ ] All helpers respect `NO_COLOR` environment variable (degrade to plain text)
- [ ] Visual test: a simple `main.go` that prints all styled output types looks correct in terminal

---

### Issue 1.3: TOML Config Parser

**Labels:** `core`, `phase-1`

**Description:**
Create the config package that reads and validates `normalize.toml`. Define Go structs that map to the TOML schema (target name, type, common flag, OS paths, repo location). Support `type = "file"` only for this issue.

**Acceptance Criteria:**
- [ ] `config.Load(path string)` reads a TOML file and returns a typed config struct
- [ ] Struct supports fields: `type`, `common`, `paths.linux`, `paths.macos`, `repo.location`
- [ ] Returns a clear error if the TOML file is missing or malformed
- [ ] Returns a clear error if a target has `type` other than `file` (for now)
- [ ] `config.GetTarget(name string)` retrieves a single target by its TOML key
- [ ] Error messages use the Lip Gloss styled output helpers
- [ ] Unit tests cover: valid config, missing file, malformed TOML, unknown target name

---

### Issue 1.4: OS Detection & Path Resolver

**Labels:** `core`, `phase-1`

**Description:**
Create the resolver package. It should auto-detect the current OS (`runtime.GOOS`), expand `~` to the user's home directory, and resolve the correct system path and repo path for a given target and OS.

**Acceptance Criteria:**
- [ ] `resolver.DetectOS()` returns `"linux"` or `"macos"`
- [ ] `resolver.ExpandPath(path string)` replaces `~` with `os.UserHomeDir()`
- [ ] `resolver.Resolve(target, os)` returns the expanded system path and repo path
- [ ] Returns an error if the target doesn't have a path for the requested OS
- [ ] Unit tests cover: Linux detection, macOS detection, tilde expansion, missing OS path

---

### Issue 1.5: Backup Command — Single File

**Labels:** `command`, `phase-1`

**Description:**
Implement the `backup` Cobra subcommand for `type = "file"` targets. It reads the target name as an argument, resolves paths using the current OS, and copies the file from the system path to the repo path. Creates parent directories in the repo if they don't exist.

**Acceptance Criteria:**
- [ ] `normalize backup git` copies `~/.gitconfig` to `<repo>/common/git/.gitconfig`
- [ ] Parent directories are created automatically if missing
- [ ] Prints Lip Gloss styled success message with source → destination
- [ ] Prints Lip Gloss styled error if the source file doesn't exist
- [ ] Prints Lip Gloss styled error if the target name isn't in the config
- [ ] File content is identical after copy (byte-for-byte)
- [ ] Unit/integration test using temp directories

---

### Issue 1.6: Restore Command — Single File

**Labels:** `command`, `phase-1`

**Description:**
Implement the `restore` Cobra subcommand for `type = "file"` targets. It reads the target name as an argument, resolves paths, and copies the file from the repo to the system path. Creates parent directories on the system if they don't exist.

**Acceptance Criteria:**
- [ ] `normalize restore git` copies `<repo>/common/git/.gitconfig` to `~/.gitconfig`
- [ ] Parent directories are created automatically if missing
- [ ] Prints Lip Gloss styled success message with source → destination
- [ ] Prints Lip Gloss styled error if the repo file doesn't exist
- [ ] File content is identical after copy (byte-for-byte)
- [ ] Unit/integration test using temp directories

---

### Issue 1.7: Normalize Tool Config

**Labels:** `core`, `phase-1`

**Description:**
Create a separate config loader for Normalize's own settings (`~/.config/normalize/config.toml`). This stores the repo path, network path, and backup cache settings. The backup and restore commands should read the repo path from this config instead of being hardcoded.

**Acceptance Criteria:**
- [ ] Reads `~/.config/normalize/config.toml`
- [ ] Exposes `repo.path`, `network.path`, `backup.cache_dir`, `backup.max_age_days`
- [ ] Falls back to sensible defaults if config file is missing (`~/dotfiles`, etc.)
- [ ] Backup and restore commands use `repo.path` from this config
- [ ] Unit tests cover: present config, missing config (defaults), partial config

---

## Milestone 2: Flags & Directories

**Description:** Add OS selection flags, --dry-run, the -a (all) flag, and recursive directory support.
**Target:** 1–2 weekends

---

### Issue 2.1: OS Selection Flags (-l, -m)

**Labels:** `cli`, `phase-2`

**Description:**
Add `-l` / `--linux` and `-m` / `--macos` persistent flags to the root command. When set, they override OS auto-detection. If neither is set, auto-detect. If both are set, return an error.

**Acceptance Criteria:**
- [ ] `-l` forces Linux path resolution regardless of current OS
- [ ] `-m` forces macOS path resolution regardless of current OS
- [ ] No flag = auto-detect via `runtime.GOOS`
- [ ] Both flags set simultaneously = clear styled error message
- [ ] Flags are available on all subcommands (backup, restore, list, diff, status)
- [ ] Unit tests cover: each flag individually, both flags, neither flag

---

### Issue 2.2: All Targets Flag (-a)

**Labels:** `cli`, `phase-2`

**Description:**
Add `-a` / `--all` flag to backup and restore commands. When set, the command operates on every target in the config that has a path for the selected OS (instead of requiring a target name argument).

**Acceptance Criteria:**
- [ ] `normalize backup -a` backs up all targets available for the current/selected OS
- [ ] `normalize restore -a` restores all targets available for the current/selected OS
- [ ] Targets without a path for the selected OS are skipped with a styled warning
- [ ] Lip Gloss styled summary printed at end: N backed up, N skipped, N failed
- [ ] Cannot be combined with a specific target name (error if both provided)

---

### Issue 2.3: Dry Run Mode (--dry-run)

**Labels:** `cli`, `phase-2`

**Description:**
Add `--dry-run` flag to backup and restore commands. When set, all file operations are printed to stdout but not executed. Output uses Lip Gloss dimmed styling to visually distinguish dry runs from real operations.

**Acceptance Criteria:**
- [ ] `normalize backup git --dry-run` prints styled: `[DRY RUN] Would copy ~/.gitconfig → <repo>/common/git/.gitconfig`
- [ ] Dry run output uses Lip Gloss dimmed/italic style to distinguish from real operations
- [ ] No files are created, modified, or deleted during dry run
- [ ] Works with `-a` flag (shows all operations that would occur)
- [ ] Works with directory targets (shows each file that would be copied)

---

### Issue 2.4: Directory Backup & Restore

**Labels:** `core`, `phase-2`

**Description:**
Extend the engine to handle `type = "directory"` targets. Backup recursively copies the source directory to the repo location. Restore copies from the repo to the system. Preserve directory structure and handle nested files.

**Acceptance Criteria:**
- [ ] `normalize backup nvim` recursively copies `~/.config/nvim/` to `<repo>/common/nvim/`
- [ ] Preserves subdirectory structure
- [ ] `normalize restore nvim` recursively copies from repo back to system path
- [ ] Handles empty subdirectories (skip or create, document choice)
- [ ] Handles symlinks (skip with warning, document choice)
- [ ] Prints styled count of files copied on success
- [ ] Integration tests with nested directory structures

---

### Issue 2.5: Verbose Output Flag (-v)

**Labels:** `cli`, `phase-2`

**Description:**
Add `-v` / `--verbose` flag. When set, Charm Log outputs detailed information for each file operation (source path, destination path, file size, whether it was created or overwritten).

**Acceptance Criteria:**
- [ ] Default output: one-line styled summary per target
- [ ] Verbose output: per-file Charm Log details within directory targets
- [ ] Verbose + dry-run shows what each file operation would do
- [ ] Verbose flag works across all subcommands

---

## Milestone 3: List, Diff & Status

**Description:** Add informational commands using Bubbles table components and Lip Gloss styling.
**Target:** 1 weekend

---

### Issue 3.1: List Command

**Labels:** `command`, `phase-3`

**Description:**
Implement the `list` subcommand using the Bubbles table component. Displays all registered targets from `normalize.toml` with their type, which OSes they support, and their system/repo paths in a styled, aligned table.

**Acceptance Criteria:**
- [ ] `normalize list` shows all targets in a Bubbles table with styled headers
- [ ] Each row shows: target name, type (file/directory/split), OS support (linux/macos/both), repo location
- [ ] `normalize list -l` filters to show only targets with Linux paths
- [ ] `normalize list -m` filters to show only targets with macOS paths
- [ ] Table uses Lip Gloss theme colors for headers and alternating row shading
- [ ] Graceful fallback to simple output if terminal is too narrow

---

### Issue 3.2: Diff Command

**Labels:** `command`, `phase-3`

**Description:**
Implement the `diff` subcommand. Compares the live system file(s) with the repo version and displays differences. Use SHA-256 checksums for quick comparison, with optional detailed diff output.

**Acceptance Criteria:**
- [ ] `normalize diff git -l` compares system `.gitconfig` with repo version
- [ ] If identical: prints Lip Gloss styled "in sync" message (green ✓)
- [ ] If different: prints a summary (file sizes, modification times) and optionally the diff
- [ ] For directory targets: lists files that are added, removed, or modified
- [ ] `--verbose` shows actual content diff (shell out to `diff` or use go-diff)
- [ ] Handles missing files gracefully (file only on system, file only in repo)

---

### Issue 3.3: Status Command

**Labels:** `command`, `phase-3`

**Description:**
Implement the `status` subcommand. Shows a summary overview of all targets' sync state for the selected OS using Lip Gloss styled indicators.

**Acceptance Criteria:**
- [ ] `normalize status` shows all targets for the current OS with sync state
- [ ] States: ✓ in-sync (green), ✗ modified (yellow), ? missing (red) — all Lip Gloss styled
- [ ] Targets displayed using Bubbles table or Lip Gloss aligned columns
- [ ] Summary line at bottom: "12 targets: 8 synced, 3 modified, 1 missing" (styled)
- [ ] Respects `-l` / `-m` OS flags

---

## Milestone 4: Split Files & NAS Fallback

**Description:** Handle split/concat files with delimiter-based OS sections, add --network fallback, and introduce Bubbles progress components.
**Target:** 2 weekends

---

### Issue 4.1: Split File Backup

**Labels:** `core`, `phase-4`

**Description:**
Implement backup logic for `type = "split"` targets. Read the live system file, find the delimiter, write everything above the delimiter to `repo.base`, and everything from the delimiter down to `repo.<os>`.

**Acceptance Criteria:**
- [ ] Reads the delimiter from the target's TOML config
- [ ] Correctly splits content at the delimiter line
- [ ] Writes base content to `repo.base` path
- [ ] Writes platform content (including delimiter line) to `repo.<os>` path
- [ ] If no delimiter found: writes entire file to `repo.base` and prints styled warning
- [ ] Creates parent directories as needed
- [ ] Unit tests with various delimiter positions (top, middle, bottom, missing)

---

### Issue 4.2: Split File Restore (Concat)

**Labels:** `core`, `phase-4`

**Description:**
Implement restore logic for `type = "split"` targets. Read `repo.base` and `repo.<os>`, concatenate them (base first), and write the combined result to the system path.

**Acceptance Criteria:**
- [ ] Reads base file and platform file from repo
- [ ] Concatenates: base content + newline + platform content
- [ ] Writes combined result to the system path
- [ ] If platform file missing: restores only base with styled warning
- [ ] If base file missing: returns styled error (base is required)
- [ ] Resulting file matches what backup would produce from the original
- [ ] Round-trip test: backup then restore produces identical file

---

### Issue 4.3: Network Fallback Flag (--network)

**Labels:** `cli`, `phase-4`

**Description:**
Add `--network` flag to backup and restore commands. When set, the copy destination/source switches from the git repo path to the NAS path (defined in `~/.config/normalize/config.toml`).

**Acceptance Criteria:**
- [ ] `--network` redirects all file operations to `network.path` instead of `repo.path`
- [ ] Same directory structure is used on the NAS as in the repo
- [ ] If NAS path is not accessible (not mounted, permission denied): clear styled error with troubleshooting hint
- [ ] If a normal backup/restore fails due to path errors, the error message suggests: "Try again with --network if your repo is unavailable"
- [ ] Works with all target types (file, directory, split)
- [ ] Works with -a, --dry-run, and other flags

---

### Issue 4.4: Bubbles Progress Components

**Labels:** `ui`, `phase-4`

**Description:**
Create reusable Bubble Tea models for a spinner (single operations) and a progress bar (bulk operations). Use these during long-running operations like bulk backup/restore and network transfers.

**Acceptance Criteria:**
- [ ] Spinner model wraps a Bubbles spinner with Lip Gloss theme colors
- [ ] Progress bar model wraps a Bubbles progress bar with theme colors
- [ ] Spinner displays during single-target operations that take >500ms
- [ ] Progress bar displays during `-a` operations showing N/total targets complete
- [ ] Both components clear cleanly when the operation finishes
- [ ] Falls back to simple text output when stdout is not a terminal (piped output)

---

### Issue 4.5: Pre-Restore Backup Cache

**Labels:** `safety`, `phase-4`

**Description:**
Before any restore overwrites a live system file, save a timestamped copy to `~/.normalize/backups/`. Implement cache pruning to remove backups older than `max_age_days`.

**Acceptance Criteria:**
- [ ] Before overwriting, copies existing system file to `~/.normalize/backups/<target>/<timestamp>.<filename>`
- [ ] Timestamp format: `YYYYMMDD-HHMMSS`
- [ ] Backup cache directory created automatically if it doesn't exist
- [ ] Pruning runs on each restore: deletes backups older than configured max age
- [ ] `--dry-run` shows what would be backed up without doing it
- [ ] Works for files and directories (directory targets get zipped or copied as tree)

---

## Milestone 5: Bootstrap Init

**Description:** One-command full environment restore using Bubble Tea for interactive confirmation and real-time progress.
**Target:** 1 weekend

---

### Issue 5.1: Bubble Tea Confirmation Model

**Labels:** `ui`, `phase-5`

**Description:**
Create a Bubble Tea interactive confirmation model for the init command. Before restoring, display a styled preview of all targets that will be restored and prompt for y/n confirmation. Add `--yes` flag to bypass for scripted usage.

**Acceptance Criteria:**
- [ ] Bubble Tea model shows a Lip Gloss styled list of all targets to be restored
- [ ] Each target shows: name, type, destination path
- [ ] User can confirm (y/enter) or cancel (n/esc/q)
- [ ] Cancel exits cleanly with no changes
- [ ] `--yes` flag bypasses the interactive prompt entirely
- [ ] Falls back to simple y/n prompt when stdout is not a terminal

---

### Issue 5.2: Bubble Tea Multi-Select Target Picker

**Labels:** `ui`, `phase-5`

**Description:**
Create an optional Bubble Tea multi-select model that lets users pick a subset of targets to restore during init, instead of restoring everything.

**Acceptance Criteria:**
- [ ] `normalize init -l --pick` launches an interactive multi-select list
- [ ] All targets for the selected OS are shown with checkboxes
- [ ] Arrow keys navigate, space toggles selection, enter confirms
- [ ] Selected targets are highlighted with Lip Gloss theme accent color
- [ ] After selection, flows into the confirmation model (Issue 5.1)
- [ ] Without `--pick`, all targets are selected by default (existing behavior)

---

### Issue 5.3: Init Command & Progress View

**Labels:** `command`, `phase-5`

**Description:**
Implement the `init` subcommand. Restores all (or selected) targets for the selected OS. Uses a Bubble Tea progress view showing each target's restore status in real time.

**Acceptance Criteria:**
- [ ] `normalize init -l` restores all targets with Linux paths
- [ ] `normalize init -m` restores all targets with macOS paths
- [ ] No OS flag = auto-detect current OS
- [ ] Bubble Tea progress view shows: current target name, progress bar (N/total), per-target status (✓ done, ⟳ in progress, — pending)
- [ ] `--dry-run` shows full plan without executing (no interactive progress)

---

### Issue 5.4: Target Ordering

**Labels:** `core`, `phase-5`

**Description:**
Add an optional `order` field to target definitions in `normalize.toml`. The init command should restore targets in ascending order. Targets without an order field are restored last.

**Acceptance Criteria:**
- [ ] `order = 1` in TOML config defines restore priority (lower = earlier)
- [ ] Init command sorts targets by order before restoring
- [ ] Targets without `order` default to `999` (restored last)
- [ ] Ordering is displayed in the Bubble Tea progress view and during dry-run
- [ ] Unit test verifies ordering logic

---

### Issue 5.5: Init Summary Report

**Labels:** `output`, `phase-5`

**Description:**
After init completes, print a comprehensive Lip Gloss styled summary report showing what happened to each target.

**Acceptance Criteria:**
- [ ] Summary grouped by outcome: restored, skipped, failed
- [ ] Each entry shows: target name, type, source path, destination path
- [ ] Failed entries include the error reason (styled red)
- [ ] Skipped entries include the reason (styled yellow, e.g., "no macOS path defined")
- [ ] Total counts at the bottom in a styled box
- [ ] Option to write summary to a log file (`--log` flag or automatic)

---

### Issue 5.6: End-to-End Testing

**Labels:** `testing`, `phase-5`

**Description:**
Create comprehensive end-to-end tests that simulate a full backup → fresh state → init → verify cycle using temporary directories.

**Acceptance Criteria:**
- [ ] Test creates a mock system with dotfiles in temp dirs
- [ ] Runs backup -a to populate a temp repo
- [ ] Clears the mock system directories
- [ ] Runs init to restore everything
- [ ] Verifies all files match originals (checksum comparison)
- [ ] Tests both file and directory targets
- [ ] Tests at least one split file target
- [ ] Tests run in CI (no real system paths touched)

---

## Backlog (Future)

These are tracked but not scheduled. Create these as issues with a `backlog` label.

- **Git auto-commit:** Optionally run `git add . && git commit -m "normalize backup"` after backup operations
- **Package manifest:** Track installed packages (pacman, brew) alongside dotfiles
- **Encrypted secrets:** Support for encrypting sensitive files (SSH keys, tokens) with age
- **Conflict resolution:** Handle cases where both system and repo have diverged since last sync
- **Watch mode:** Use fsnotify to auto-backup on file changes
- **Shell completions:** Generate bash/zsh/fish completions via Cobra
- **Interactive diff viewer:** Bubble Tea side-by-side diff model for reviewing changes before backup/restore
- **Glamour integration:** Render markdown help docs and changelogs beautifully in-terminal
- **Wish SSH server:** SSH-accessible Normalize TUI hosted on the NAS for remote management
