# CliExamples

Every example of the teste cli. Each one is a shell session that runs with its own
directory as the working directory and writes only into its own `TestDir`, so it can be read
as documentation and copied line by line. The script types `teste`, which `exec-test`
resolves to the code in this tree. It ends by copying out of `TestDir` into `AssertDir` the
paths it asserts — `mkdir -p AssertDir/<path>` then `cp -R TestDir/<path>/. AssertDir/<path>/`,
each keeping the place it holds in the tree.

`agnos exec-test` runs them all and checks each against the `result.yaml` beside it — the
golden holding the output, the exit code and the sha256 of every `AssertDir` file, written by
`exec-test` and never by hand. [Workflow](../Workflow/doc.md) has the commands that add and
remove one; the lib side is [LibExamples](../LibExamples/doc.md).

No example is declared yet: `examples/cli/` is created by the first
`agnos add-cli-example`.

