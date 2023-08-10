# logfile

Easily write all Golang "log" in a file

## Usage

```go
import "github.com/l3dlp/logfile"
```
[...]
```go
func main() {
  logFile := logfile.LogFile("chemindufichier.log")
  defer logFile.Close()
  log.Println("blablabla")
}
```

## License

Ce projet est sous licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.

