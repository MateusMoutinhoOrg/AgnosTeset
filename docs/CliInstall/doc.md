# CliInstall

`backoffice` is a single static binary: no runtime, no dependencies. Pick your platform,
paste the block, done. Go 1.25+ is needed only to build it from source.

**macOS (Apple Silicon)**

```bash
curl -sL https://github.com/MateusMoutinhoOrg/AgnosTeset/releases/latest/download/macarm64.bin -o backoffice && chmod +x backoffice && sudo mv backoffice /usr/local/bin/ && backoffice version
```

**macOS (Intel)**

```bash
curl -sL https://github.com/MateusMoutinhoOrg/AgnosTeset/releases/latest/download/mac86.bin -o backoffice && chmod +x backoffice && sudo mv backoffice /usr/local/bin/ && backoffice version
```

**Linux (amd64)**

```bash
curl -sL https://github.com/MateusMoutinhoOrg/AgnosTeset/releases/latest/download/linux86.out -o backoffice && chmod +x backoffice && sudo mv backoffice /usr/local/bin/ && backoffice version
```

**Linux (arm64)**

```bash
curl -sL https://github.com/MateusMoutinhoOrg/AgnosTeset/releases/latest/download/linuxarm64.out -o backoffice && chmod +x backoffice && sudo mv backoffice /usr/local/bin/ && backoffice version
```

**Linux (32-bit)**

```bash
curl -sL https://github.com/MateusMoutinhoOrg/AgnosTeset/releases/latest/download/linuxi32.out -o backoffice && chmod +x backoffice && sudo mv backoffice /usr/local/bin/ && backoffice version
```

**Windows (64-bit)** — PowerShell:

```powershell
$dir="$HOME\.local\bin"; New-Item -ItemType Directory -Force -Path $dir | Out-Null; curl.exe -sL https://github.com/MateusMoutinhoOrg/AgnosTeset/releases/latest/download/windows86.exe -o "$dir\backoffice.exe"; [Environment]::SetEnvironmentVariable('PATH', [Environment]::GetEnvironmentVariable('PATH','User') + ";$dir", 'User')
```

**Windows (32-bit)** — PowerShell:

```powershell
$dir="$HOME\.local\bin"; New-Item -ItemType Directory -Force -Path $dir | Out-Null; curl.exe -sL https://github.com/MateusMoutinhoOrg/AgnosTeset/releases/latest/download/windowsi32.exe -o "$dir\backoffice.exe"; [Environment]::SetEnvironmentVariable('PATH', [Environment]::GetEnvironmentVariable('PATH','User') + ";$dir", 'User')
```

**From a checkout** — needs Go 1.25+:

```bash
go build -o backoffice ./cmd/main && sudo mv backoffice /usr/local/bin/ && backoffice version
```

The released binaries are the ones `agnos compile --target all` builds and `agnos publish`
uploads. `backoffice version` prints the `version` of `AgnosConfig/project.yaml`,
`backoffice help` every command — each one is listed in [Commands](../Commands/doc.md).
