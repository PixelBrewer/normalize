# normalize

Go CLI for backing up and restoring dotfiles/configs across Linux and macOS.

## Module

`github.com/PixelBrewer/normalize` — Go 1.26.1

## Commands

```sh
go build ./...     # compile
go run .           # run locally
go install .       # install to $GOPATH/bin
```

No Makefile, CI, or test suite yet.

## Architecture

- `main.go` — entry point (TODO: wire up Cobra root command)
- `cmd/` — Cobra subcommands: backup, restore, init, list, diff, status (mostly stubs)
- `internal/config/` — TOML config loader (stub, hardcoded `/home/alex/` path — do not hardcode paths)
- `internal/engine/` — backup/restore/split logic (stubs)
- `internal/resolver/` — OS detection and path expansion (stub)
- `internal/ui/` — Lip Gloss styled output, Charm Log logger (stubs)

## Planned Dependencies (not yet installed)

cobra, bubbletea, lipgloss, bubbles, charmbracelet/log, BurntSushi/toml

Run `go get` for each when implementing commands.

## Design & Issue Tracking

`docs/Normalize-GitHub-Issues.md` has full milestone/issue specs with acceptance criteria. Reference this before implementing features.

## Current State

Most files are `// TODO: Implement...` stubs. Do not assume any package is functional. The `backup.go` stub has a compile error (`bufio.scann()` → `bufio.NewScanner()`).

## Style

- Use Lip Gloss for all terminal output (theme in `internal/ui/theme.go`)
- Use Charm Log for structured/verbose logging
- Respect `NO_COLOR` env var
- Unit tests expected per issue (temp dir patterns)
