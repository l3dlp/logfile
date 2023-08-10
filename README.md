# logfile

Listen up, you filthy troglodytes inhabiting the cyber-scarred wastelands of code. You're tired of those log messages littering your console like the droppings of a thousand ill-mannered pigeons? Got a hankering for a solution that's not as tedious as your day-to-day existence? Enter logfile.

This `logfile` package writes all your Golang "log" into a file, like some sort of obedient scribe tirelessly scribbling the rantings of a madman. In a world filled with bloated garbage, this is a tool that makes sense.

## Usage
Oh, you want to know how to use it? Well, isn't that precious? Fine, I'll hold your hand like the delicate child you are:

```go
import "github.com/l3dlp/logfile"
// [...]
func main() {
  logFile := logfile.Use("chemindufichier.log")
  if logFile != nil {
    defer logFile.Close()
  }
  log.Println("blablabla")
}
```

See that? Even you can do it. It's as simple as buying off a politician or ignoring the nagging sense that your life is a repetitive cycle of despair. You just Use it, and BAM! You've got a logfile, sleek and beautiful, like the truth in a world full of lies.

## License
[MIT license](LICENSE). It's free and unburdened by moral constraint, like a good systems library should be.

---

Now, run along, kids. Dive into the world of logfile and embrace the brutal, unfiltered reality it offers. Or don't. Either way, the world's still turning, and I've still got deadlines to meet.

