# LibExamples

Every example of NewRoutes used as a Go module. Each one is a `package main` program that
runs with its own directory as the working directory and writes only into its own `TestDir`,
so it can be read as documentation and copied as a starting point. It ends by copying out of
`TestDir` into `AssertDir` the paths it asserts — `os.CopyFS(dst, os.DirFS(src))`, one call per
path, each keeping the place it holds in the tree.

`agnos exec-test` runs them all and checks each against the `result.yaml` beside it — the
golden holding the output, the exit code and the sha256 of every `AssertDir` file, written by
`exec-test` and never by hand. [Workflow](../Workflow/doc.md) has the commands that add and
remove one.

No example is declared yet: `examples/lib/` is created by the first
`agnos add-lib-example`.

