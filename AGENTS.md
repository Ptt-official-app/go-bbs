# Repository Guidelines

## Project Overview
`github.com/Ptt-official-app/go-bbs`: Go library for reading and writing the on-disk data structures of Taiwanese BBS systems (primarily pttbbs / CurrentPTT / OpenPTT; formosabbs is partial). It parses fixed-size binary records (`.PASSWDS`, `.BRD`, `.DIR`, `.fav`, `logins.bad`, `logins.recent`, USERINFO) and the shared-memory cache. `go.mod` declares `go 1.14`. Only dependencies: `golang-jwt/jwt/v4`, `golang.org/x/sys`, `golang.org/x/text`.

## Architecture & Data Flow
- Root package `bbs` defines driver-agnostic interfaces (`UserRecord`, `BoardRecord`, `ArticleRecord`, `Connector`, `WriteBoardConnector`, `WriteArticleConnector`, `UserArticleConnector`, `UserCommentConnector`, `UserDraftConnector`) and the `DB` wrapper in `bbs.go`.
- A driver package registers itself in `init()` via `bbs.Register("pttbbs", &Connector{})`; callers import it with a blank import and call `bbs.Open("pttbbs", "file://<bbshome>")`.
- `DB` methods ask the connector for a file path (`Get*Path`), then call a `Read*File` / write method that parses or marshals the binary records. `bbs.UnimplementedConnector` supplies defaults for unsupported methods.
- `pttbbs/cache.go` + `cache/` read the PTT shared memory (SHM) either through `mmap` of a dump file or System V `shmget`/`shmat`; offsets come from `cachePos`/`MemoryMappingSetting`.

## Key Directories
- `./` (package `bbs`): interfaces and `DB` (`bbs.go`), `Cstr`, Big5 helpers (`big5.go`), ANSI helpers (`ansi.go`), user article/comment/draft/favorite record helpers.
- `pttbbs/`: pttbbs driver (`Connector`, `passwd.go`, `board.go`, `file.go`, `fav.go`, `userinfo.go`, `bad_logins.go`, `login_recent.go`, `path.go`, `cache.go`, `pttbbs_write_*_connector.go`), `internal/fnv`, `cmd/` (SHM demo), `testcase/` (fixtures).
- `formosabbs/`: FormosaBBS file header reader (`file.go`) with `testcase/`.
- `cache/`: `Cache` interface, `Mmap` and System V `SHM` with per-OS files (`*_unix.go`, `*_linux.go`, `*_darwin.go`, `*_windows.go`).
- `crypt/`: PTT password crypt (DES-based `bbscrypt.go`, `crypt.go`).
- `filelock/`: file lock API; currently a stub (`Lock`/`Unlock` return nil, `IsLock` returns false), so record writes are not actually protected.
- `serverlet/`: sample HTTP server built on the library (`go run ./serverlet`, listens on `:8083`).
- `cmd/bbstool/`: CLI (`--driver`, `--bbshome`, subcommands such as `show_board_list`, `add_board`, `show_user_list`, `show_board_article`).

## Development Commands
```
go build ./...                    # build everything (CI step)
go test -v ./...                  # run all tests (CI step)
golangci-lint run                 # lint (CI uses golangci-lint v2.0)
gofmt -l .                        # the only formatter enabled in .golangci.yml
go run ./cmd/bbstool --bbshome=<path> show_board_list   # CLI example from cmd/bbstool/main.go
```
Branches: open PRs against `development` (`master` is admin-only), see `FOR_DEVELOPERS.md`; `CONTRIBUTING.md` asks for a linked issue number in each PR (PR template: `.github/PULL_REQUEST_TEMPLATE.md`).

## Code Conventions & Common Patterns
- Binary layouts are described with `PosOf<Struct><Field>` offset constants plus size constants in `pttbbs/const.go`; decode with `encoding/binary` little endian (`Unmarshal*` / `New*WithByte`) and encode with `Marshal*`. Match the C structs in https://github.com/ptt/pttbbs when changing offsets.
- Fixed-length C strings are converted with `bbs.CstrToString` / `newStringFormCString`; Big5-UAO text uses `big5uaoToUTF8String` / `utf8ToBig5UAOString` (`pttbbs/encoding.go`).
- File readers are `Open<Thing>File(filename)` returning slices; writers `Append*`/`Update*`/`Remove*` open the file, call `filelock.Lock`, then write. Always `defer f.Close()`.
- Path builders live in `pttbbs/path.go` (`Get*Path(workDirectory, ...)`) and index `userID[0]` / `boardID[0]` for the bucket directory; validate IDs before calling.
- Platform-specific code is split by file suffix and `//go:build` tags (keep both `//go:build` and `// +build` lines as existing files do).
- Source files carry the Apache 2.0 header from the original authors; comments are a mix of Traditional Chinese and English, follow the surrounding file.

## Important Files
`bbs.go`, `pttbbs/pttbbs.go`, `pttbbs/const.go`, `pttbbs/path.go`, `pttbbs/passwd.go`, `pttbbs/board.go`, `pttbbs/file.go`, `pttbbs/cache.go`, `cache/cache.go`, `.golangci.yml`, `.github/workflows/tests.yml`, `FOR_DEVELOPERS.md`, `CONTRIBUTING.md`.

## Runtime/Tooling Preferences
- Plain Go modules, no vendor directory, no Makefile.
- No external services needed; tests use files under `testcase/` directories and temp files.
- `.golangci.yml` (v2 format) enables no linters, only the `gofmt` formatter.
- CI (`.github/workflows/tests.yml`) runs lint, `go test -v ./...` and `go build -v ./...` on macOS, Ubuntu and Windows; the `cache` tests are restricted to unix-like systems by build tags.

## Testing & QA
- Standard `testing` package, table-driven tests, `<source>_test.go` beside the source; no third-party test framework.
- Fixtures: `pttbbs/testcase/` (`passwd`, `board`, `boards`, `fav`, `file`, `bad_logins`, `logins_recent`, `userinfo`) and `formosabbs/testcase/`.
- Some write tests modify fixtures in place (e.g. `pttbbs/testcase/boards/t/test/M.1606672292.A.B23` shows as modified after `go test ./pttbbs`); run `git checkout -- pttbbs/testcase` before committing.
- Add a regression test when fixing parser bugs, preferably with `t.TempDir()` files rather than editing fixtures.
- `pttbbs/path_benchmark_test.go` holds benchmarks (`go test -bench . ./pttbbs`).
