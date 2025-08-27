
<h3 align="center">filebin-cli</h3>


<p align="center">
    A command-line client for filebin.net - a simple, temporary file sharing service.
    <br>
    <br>
    <a href="##installation">Install</a>
    ·
    <a href="##usage">Usage</a>
    ·
    <a href="##license">License</a>
    ·
    <a href="##contributing">Contributing</a>
    ·
    <a href="https://kavin.is-a.dev/blog/soon">Blog</a>
</p>


> [!WARNING]
> filebin-cli is not affiliated with [filebin.net](https://filebin.net) <br>
> What you upload or do with this tool is entirely your responsibility. You must comply with [filebin.net's Terms of Service](https://filebin.net/terms)

## Installation

### Download the latest release bundle

Check the [releases page](https://github.com/kavin81/filebin-cli/releases).


### Compiling from source code


```bash
# download the source code
git clone https://github.com/kavin81/filebin-cli.git
cd filebin-cli

# build the binary
go build -o filebin
mv filebin /usr/local/bin/

# cleanup
cd ..
rm -rf filebin-cli
```

## Usage

### Quick Start

```bash
# Upload a file to a bin
filebin push my-bin document.pdf

# Download files from a bin
filebin pull my-bin

# Show bin information
filebin show my-bin

# Delete a bin
filebin delete my-bin
```

### Command Reference

```
Usage:
  filebin [command]

Available Commands:
  delete      Remove bin or file from bin
  lock        Lock bin to prevent uploads
  pull        Download bin or file from bin
  push        Upload file to bin
  show        Show information about bin
  help        Help about any command

Global Flags:
  -h, --help      help for filebin
      --icons     Set icon display mode (experimental)
      --json      Enable JSON output
  -v, --verbose   Enable verbose output

Use "filebin [command] --help" for more information about a command.
```

---

## Commands

### `push` - Upload Files

Upload files to a filebin for sharing. If the bin doesn't exist, it will be created automatically.

```bash
# Basic upload
filebin push team-docs presentation.pptx
filebin push temp-share ./archive.zip

# Upload with custom filename
filebin push public-files image.jpg --filename "hero-image.jpg"
```

| Flag | Description |
|------|-------------|
| `--client-id` | Custom client identifier |
| `--filename` | Custom filename for the uploaded file |

**Aliases:** `upload`, `put`

---

### `pull` - Download Files

Download files from a filebin. Downloads entire bin as archive if no file is specified.

```bash
# Download entire bin as archive
filebin pull team-docs

# Download specific file
filebin pull temp-share/archive.zip
filebin pull public-files presentation.pptx

# Download to specific directory
filebin pull team-docs --output ./downloads
```

| Flag | Description | Default |
|------|-------------|---------|
| `-o, --output` | Output directory | Current directory |

**Aliases:** `get`, `fetch`

---

### `show` - Display Bin Information

Display detailed information about a bin including files, metadata, file sizes, upload dates, and bin status.

```bash
# Show bin information
filebin show team-docs
filebin show temp-share

# Show as JSON
filebin show public-files --json
```

**Aliases:** `info`

---

### `delete` - Remove Bins or Files

Remove a bin or specific file from a filebin. The bin will be permanently deleted if no file is specified.

```bash
# Delete entire bin
filebin delete team-docs

# Delete specific file from bin
filebin delete temp-share/archive.zip
filebin delete public-files presentation.pptx
```

**Aliases:** `remove`, `del`, `rm`

---

### `lock` - Lock a Bin

Lock a bin to prevent further file uploads. **This action is irreversible once applied.**

```bash
# Lock a bin
filebin lock team-docs
filebin lock temp-share
```

---

## Global Options

| Flag | Description |
|------|-------------|
| `-h, --help` | Show help information |
| `--icons` | Set icon display mode (experimental) |
| `--json` | Enable JSON output |
| `-v, --verbose` | Enable verbose output |


## License
- This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

contributors are welcome! Contributions of all sizes are appreciated

- Open an issue.
- Package it for your favorite distribution or package manager
- Share it with a friend!
- Open a pull request
