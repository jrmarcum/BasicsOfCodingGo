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

Attribution for Go by Example derived content is provided centrally in
README.md and NOTICE — do **not** add a per-file attribution footer to
lesson `.md` files.

## Project Structure

```
BasicsOfCodingGo/
├── CLAUDE.md          — this file; canonical project context for Claude sessions
├── LICENSE            — CC0 (applies to Jon Marcum's original contributions)
├── NOTICE             — attribution notice for CC BY 3.0 derived content
├── README.md          — project overview, attribution section, license table
└── ##_topic-name/
    ├── topic-name.go  — runnable Go source for the lesson
    └── topic-name.md  — lesson explanation (run command + expected output)
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
  derived content. Always refer to NOTICE and README for the full picture;
  do not treat CC0 as applying to the whole repo.
- The project has a `go.mod` file (or individual files use `package main`).
  Check before adding new files that depend on modules.
- Lesson `.md` files use a consistent format: optional description line,
  `___` divider, `##### Run Command:`, run command in backticks, blank line,
  `##### Results:`, each output line in backticks. No per-file footer.
- Lessons 58–60 (reading-files, writing-files, line-filters) use `./tmp/` for
  file I/O. The `tmp/` directory is gitignored and must exist at runtime; for
  lesson 58, `tmp/dat.txt` must be pre-populated (content: `hello\ngo\n`).
- **Go 1.25 WASI file I/O bug:** Any WASI runner for this project must pass
  `--dir ".::/"` to wasmtime (maps host `.` → guest `/`). Root cause:
  `preparePath` calls `joinPath(".", "hello.txt")` which strips the leading
  `.`, yielding `"hello.txt"`; the preopen-matching loop then checks
  `HasPrefix("hello.txt", ".")` → `false`, so `path_open` is called with
  FD -1 → EBADF ("Bad file number"). Mapping to `/` instead makes the preopen
  name `/`, so `HasPrefix("/hello.txt", "/")` matches correctly. If file I/O
  fails with "Bad file number" in any lesson, this is the cause. Affects all
  lessons that do file I/O (58 reading-files, 59 writing-files, 60
  line-filters). Any new WASI runner must use `--dir ".::/"` for Go 1.25+
  with wasmtime.
