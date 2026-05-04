# Basics of Coding Go — Project Context

## Purpose

Multi-language comparative study of programming syntax, language simplicity,
lines of code required, and compile/run performance. Go is one of several
languages implemented against the same set of example programs, enabling
direct side-by-side comparison.

## Licensing Summary

This project contains two tiers of content with different licenses:

- **CC BY 3.0** — lesson files and code examples adapted from "Go by Example"
  by Mark McGranaghan (https://github.com/mmcgrana/gobyexample).
  License: http://creativecommons.org/licenses/by/3.0/

- **CC0 1.0** — original contributions by Jon Marcum (project structure,
  README, comparative-study additions, and any lessons not derived from
  Go by Example). See LICENSE.

Project-level attribution is covered in NOTICE and the README Attribution/
License sections. Per-file attribution is covered by the footer that appears
at the bottom of every lesson `.md` file.

When adding a new lesson adapted from Go by Example, include this footer:

```
___
###### This work and the accompanying code was originally from Mark McGranaghan at [https://github.com/mmcgrana/gobyexample](https://github.com/mmcgrana/gobyexample) and licensed under a Creative Commons Attribution 3.0 Unported License [http://creativecommons.org/licenses/by/3.0/](http://creativecommons.org/licenses/by/3.0/). It has been used to provide an example base for multiple languages to provide a basis of comparitive programming language study for syntax, language simplicity, number of lines of code and operations required to perform the same task, as well as compile and run speed combined.
```

## Project Structure

```
BasicsOfCodingGo/
├── CLAUDE.md          — this file; canonical project context for Claude sessions
├── LICENSE            — CC0 (applies to Jon Marcum's original contributions)
├── NOTICE             — attribution notice for CC BY 3.0 derived content
├── README.md          — project overview, attribution section, license table
└── ##_topic-name/
    ├── topic-name.go  — runnable Go source for the lesson
    └── topic-name.md  — lesson explanation with per-file attribution footer
```

Lessons are numbered with a two-digit prefix (e.g., `01_hello-world`). Numbers
are not strictly contiguous — gaps exist where topics were skipped or reserved
for future additions.

## .gitignore

The project `.gitignore` covers:

```gitignore
# Compiled binaries produced by go build (lessons 64-66)
*.exe
command-line-arguments
command-line-flags
command-line-subcommands

# Temporary files created by lesson examples (lessons 58-60)
tmp/

# Go test and coverage artifacts
*.test
*.out
```

- `*.exe` and the three named binaries are produced when a reader runs `go build` in lessons 64–66.
- `tmp/` is the working directory expected by lessons 58 (reading-files), 59 (writing-files), and 60 (line-filters). It must exist at runtime but should not be committed.
- `*.test` / `*.out` are standard Go toolchain artifacts.

## Notes for Future Claude Sessions

- The root `LICENSE` file is CC0 but does **not** cover the Go by Example
  derived content. Always refer to NOTICE and per-file footers for the full
  picture; do not treat CC0 as applying to the whole repo.
- The project has a `go.mod` file (or individual files use `package main`).
  Check before adding new files that depend on modules.
- Lesson `.md` files use a consistent format: narrative explanation, code
  block with output, horizontal rule, attribution footer. Follow this pattern
  when adding lessons.
- Lessons 58–60 (reading-files, writing-files, line-filters) use `./tmp/` for
  file I/O. The `tmp/` directory is gitignored and must exist at runtime; for
  lesson 58, `tmp/dat.txt` must be pre-populated (content: `hello\ngo\n`).
- The `go-run.ps1` runner uses `wasmtime run --dir ".::/"` (host `.` → guest
  `/`). This works around a Go 1.25 WASI bug where `preparePath` strips the
  leading `.` from relative paths, causing `path_open` to receive FD -1.
