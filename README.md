## Usage

### Installation
```bash
go build && go install
```

### Get Help
To list all available commands and options:
```bash
kctl --help
```

### Example Command
Display services sorted by memory usage:
```bash
kctl top service --sort-by memory --sort-by memory
```

### Shortened Version
```bash
kctl top service --s memory -H
```
### For Live Update
```bash
watch -n 5 --no-title kctl top service -s cpu -H
```
