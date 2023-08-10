# logfile

Easily write all Golang "log" in a file

## Usage

```go
import "github.com/l3dlp/logfile"
```
[...]
```go
func main() {
  logFile := logfile.Use("chemindufichier.log")
  if logFile != nil {
    defer logFile.Close()
  }
  log.Println("blablabla")
}
```

## License

This project is under MIT license - see [LICENSE file](LICENSE) for all details.

