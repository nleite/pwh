package main

import "fmt"
import "os"
import "bufio"
import "io"
import "strings"
import "regexp"

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println(`pipe me some text
	command | pwh 
for more info - https://github.com/nleite/pwh#usage`)
		return
	}

	processPipe()
}

func processPipe() {
	reader := bufio.NewReader(os.Stdin)

	var out []string

	for {
		// Probably need to see how to make it work for windows.
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
			return
		}
		//TODO: replace second attribute with flag defined value
		processLine(line, "*")
		out = append(out, string(line))
	}
}

// processLine checks if there are any passwords in the line argument, and
// replaces its characters with obf string.
func processLine(line string, obf string) {

	// check if there are any uri matching
	line, found := matchURI(line, obf)
	if !found {
		line, _ = matchConfig(line, obf)
	}
	fmt.Printf("%s", line)
}

// Apply regexp URI expression - looks for a hint of uri placeholder.
func matchURI(line string, obf string) (string, bool) {
	r := regexp.MustCompile(`(?P<uri>://(?P<userinfo>[-\w;/?:]+)@)`)
	matched := r.FindStringSubmatch(line)
	if len(matched) < 3 {
		return line, false
	}
	userInfo := matched[2]
	parts := strings.Split(userInfo, ":")
	if len(parts) > 1 {
		replacer := strings.Repeat(obf, len(parts[1]))
		return strings.Replace(line, parts[1], replacer, 1), true
	}
	return line, false
}

// Looks out to match passwords defined in config files
func matchConfig(line string, obf string) (string, bool) {
	r := regexp.MustCompile(`^((?i).*PASS?.+(:|=))([ A-z0-9"']+)`)
	matched := r.FindStringSubmatch(line)
	if len(matched) < 4 {
		return line, false
	}
	pwd := r.FindStringSubmatch(line)[3]
	if len(pwd) > 0 {
		replacer := strings.Repeat(obf, len(pwd))
		return strings.Replace(line, pwd, replacer, 1), true
	}
	return line, false
}
