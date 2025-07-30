package main

import (
	"log"
	"maps"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"sort"
	"strings"
)

var (
	gCmdWords = []string{
		"set",
		"setlocal",
		"map",
		"nmap",
		"vmap",
		"cmap",
		"cmd",
		"quit",
		"up",
		"half-up",
		"page-up",
		"scroll-up",
		"down",
		"half-down",
		"page-down",
		"scroll-down",
		"updir",
		"open",
		"jump-next",
		"jump-prev",
		"top",
		"bottom",
		"high",
		"middle",
		"low",
		"toggle",
		"invert",
		"unselect",
		"glob-select",
		"glob-unselect",
		"calcdirsize",
		"clearmaps",
		"copy",
		"cut",
		"paste",
		"clear",
		"sync",
		"draw",
		"redraw",
		"load",
		"reload",
		"echo",
		"echomsg",
		"echoerr",
		"cd",
		"select",
		"delete",
		"rename",
		"source",
		"push",
		"read",
		"shell",
		"shell-pipe",
		"shell-wait",
		"shell-async",
		"find",
		"find-back",
		"find-next",
		"find-prev",
		"search",
		"search-back",
		"search-next",
		"search-prev",
		"filter",
		"setfilter",
		"mark-save",
		"mark-load",
		"mark-remove",
		"tag",
		"tag-toggle",
		"addcustominfo",
		"tty-write",
		"cmd-escape",
		"cmd-complete",
		"cmd-menu-complete",
		"cmd-menu-complete-back",
		"cmd-menu-accept",
		"cmd-enter",
		"cmd-interrupt",
		"cmd-history-next",
		"cmd-history-prev",
		"cmd-left",
		"cmd-right",
		"cmd-home",
		"cmd-end",
		"cmd-delete",
		"cmd-delete-back",
		"cmd-delete-home",
		"cmd-delete-end",
		"cmd-delete-unix-word",
		"cmd-yank",
		"cmd-transpose",
		"cmd-transpose-word",
		"cmd-word",
		"cmd-word-back",
		"cmd-delete-word",
		"cmd-delete-word-back",
		"cmd-capitalize-word",
		"cmd-uppercase-word",
		"cmd-lowercase-word",
		"visual",
		"visual-accept",
		"visual-unselect",
		"visual-discard",
		"visual-change",
	}

	gOptWords      = getOptWords(gOpts)
	gLocalOptWords = getLocalOptWords(gLocalOpts)
)

func getOptWords(opts any) (optWords []string) {
	t := reflect.TypeOf(opts)
	for i := range t.NumField() {
		field := t.Field(i)
		switch field.Type.Kind() {
		case reflect.Map:
			continue
		case reflect.Bool:
			name := field.Name
			optWords = append(optWords, name, "no"+name, name+"!")
		default:
			optWords = append(optWords, field.Name)
		}
	}
	sort.Strings(optWords)
	return
}

func getLocalOptWords(localOpts any) (localOptWords []string) {
	t := reflect.TypeOf(localOpts)
	for i := range t.NumField() {
		field := t.Field(i)
		name := field.Name
		if field.Type.Kind() != reflect.Map {
			continue
		}
		if field.Type.Elem().Kind() == reflect.Bool {
			localOptWords = append(localOptWords, name, "no"+name, name+"!")
		} else {
			localOptWords = append(localOptWords, name)
		}
	}
	sort.Strings(localOptWords)
	return
}

func commonPrefix(s1, s2 string) string {
	r1 := []rune(s1)
	r2 := []rune(s2)

	i := 0
	for ; i < len(r1) && i < len(r2); i++ {
		if r1[i] != r2[i] {
			break
		}
	}

	return string(r1[:i])
}

type compMatch struct {
	name   string // display name in completion menu
	result string // result when cycling through completion menu
}

func matchWord2(s string, words []string) (matches []compMatch, result string) {
	for _, w := range words {
		if !strings.HasPrefix(w, s) {
			continue
		}

		matches = append(matches, compMatch{w, w})
		if len(matches) == 1 {
			result = w
		} else {
			result = commonPrefix(result, w)
		}
	}

	switch len(matches) {
	case 0:
		result = s
	case 1:
		result = result + " "
	}
	return
}

func matchCmd2(s string) (matches []compMatch, result string) {
	words := append(gCmdWords, slices.Collect(maps.Keys(gOpts.cmds))...)
	slices.Sort(words)
	matches, result = matchWord2(s, slices.Compact(words))
	return
}

func matchFile2(s string) (matches []compMatch, result string) {
	dir, file := filepath.Split(unescape(replaceTilde(s)))

	d := dir
	if dir == "" {
		d = "."
	}
	files, err := os.ReadDir(d)
	if err != nil {
		log.Printf("reading directory: %s", err)
		result = s
		return
	}

	var longest string

	for _, f := range files {
		if !strings.HasPrefix(strings.ToLower(f.Name()), strings.ToLower(file)) {
			continue
		}

		name := f.Name()
		if f.IsDir() {
			name += string(filepath.Separator)
		} else if f.Type()&os.ModeSymlink != 0 {
			if stat, err := os.Stat(filepath.Join(d, f.Name())); err == nil && stat.IsDir() {
				name += string(filepath.Separator)
			}
		}
		result := escape(dir + name)
		matches = append(matches, compMatch{name, result})

		if len(matches) == 1 {
			longest = result
		} else {
			longest = commonPrefix(strings.ToLower(longest), strings.ToLower(result))
		}
	}

	switch len(matches) {
	case 0:
		result = s
	case 1:
		result = longest
		if !strings.HasSuffix(result, string(filepath.Separator)) {
			result += " "
		}
	default:
		result = longest
	}
	return
}

func matchExec2(s string) (matches []compMatch, result string) {
	var words []string
	for _, p := range strings.Split(envPath, string(filepath.ListSeparator)) {
		files, err := os.ReadDir(p)
		if err != nil {
			if !os.IsNotExist(err) {
				log.Printf("reading path: %s", err)
			}
			continue
		}

		for _, f := range files {
			if !strings.HasPrefix(f.Name(), s) {
				continue
			}

			finfo, err := f.Info()
			if err != nil {
				log.Printf("getting file information: %s", err)
				continue
			}

			if !finfo.Mode().IsRegular() || !isExecutable(finfo) {
				continue
			}

			words = append(words, f.Name())
		}
	}

	slices.Sort(words)
	matches, result = matchWord2(s, slices.Compact(words))
	return
}

func matchSearch2(s string) (matches []compMatch, result string) {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Printf("reading directory: %s", err)
		result = s
		return
	}

	for _, f := range files {
		if !strings.HasPrefix(strings.ToLower(f.Name()), strings.ToLower(s)) {
			continue
		}

		matches = append(matches, compMatch{f.Name(), f.Name()})
		if len(matches) == 1 {
			result = f.Name()
		} else {
			result = commonPrefix(strings.ToLower(result), strings.ToLower(f.Name()))
		}
	}

	if len(matches) == 0 {
		result = s
	}
	return
}

func completeCmd2(acc []rune) (matches []compMatch, result string) {
	s := string(acc)
	f := tokenize(s)

	if len(f) == 1 {
		matches, result = matchCmd2(s)
		return
	}

	switch f[0] {
	case "set":
		if len(f) == 2 {
			matches, result = matchWord2(f[1], gOptWords)
			break
		}
		if len(f) != 3 {
			break
		}
		switch f[1] {
		case "filtermethod", "searchmethod":
			matches, result = matchWord2(f[2], []string{"glob", "regex", "text"})
		case "selmode":
			matches, result = matchWord2(f[2], []string{"all", "dir"})
		case "sortby":
			matches, result = matchWord2(f[2], []string{"atime", "btime", "ctime", "custom", "ext", "name", "natural", "size", "time"})
		default:
			if slices.Contains(gOptWords, f[1]+"!") {
				matches, result = matchWord2(f[2], []string{"false", "true"})
			}
		}
	case "setlocal":
		if len(f) == 2 {
			matches, result = matchFile2(f[1])
			break
		}
		if len(f) == 3 {
			matches, result = matchWord2(f[2], gLocalOptWords)
			break
		}
		if len(f) != 4 {
			break
		}
		switch f[2] {
		case "sortby":
			matches, result = matchWord2(f[3], []string{"atime", "btime", "ctime", "custom", "ext", "name", "natural", "size", "time"})
		default:
			if slices.Contains(gLocalOptWords, f[2]+"!") {
				matches, result = matchWord2(f[3], []string{"false", "true"})
			}
		}
	case "map", "nmap", "vmap", "cmap":
		if len(f) == 3 {
			matches, result = matchCmd2(f[2])
		}
	case "cmd":
	case "toggle":
		matches, result = matchFile2(f[len(f)-1])
	case "cd", "select", "source":
		if len(f) == 2 {
			matches, result = matchFile2(f[1])
		}
	default:
		if !slices.Contains(gCmdWords, f[0]) {
			matches, result = matchFile2(f[len(f)-1])
		}
	}

	f[len(f)-1] = result
	result = strings.Join(f, " ")
	return
}

func completeShell2(acc []rune) (matches []compMatch, result string) {
	f := tokenize(string(acc))

	switch len(f) {
	case 1:
		matches, result = matchExec2(f[0])
	default:
		matches, result = matchFile2(f[len(f)-1])
	}

	f[len(f)-1] = result
	result = strings.Join(f, " ")
	return
}

func completeSearch2(acc []rune) (matches []compMatch, result string) {
	matches, result = matchSearch2(string(acc))
	return
}

////////////////////////////////////////////////////////////////////////////////

func matchLongest(s1, s2 []rune) []rune {
	i := 0
	for ; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			break
		}
	}
	return s1[:i]
}

func matchWord(s string, words []string) (matches []string, longest []rune) {
	for _, w := range words {
		if !strings.HasPrefix(w, s) {
			continue
		}

		matches = append(matches, w)
		if len(longest) != 0 {
			longest = matchLongest(longest, []rune(w))
		} else if s != "" {
			longest = []rune(w + " ")
		}
	}

	if len(longest) == 0 {
		longest = []rune(s)
	}

	return
}

func matchExec(s string) (matches []string, longest []rune) {
	var words []string

	paths := strings.Split(envPath, string(filepath.ListSeparator))

	for _, p := range paths {
		if _, err := os.Stat(p); os.IsNotExist(err) {
			continue
		}

		files, err := os.ReadDir(p)
		if err != nil {
			log.Printf("reading path: %s", err)
		}

		for _, f := range files {
			if !strings.HasPrefix(f.Name(), s) {
				continue
			}

			finfo, err := f.Info()
			if err != nil {
				log.Printf("getting file information: %s", err)
				continue
			}

			if !finfo.Mode().IsRegular() || !isExecutable(finfo) {
				continue
			}

			words = append(words, finfo.Name())
		}
	}

	if len(words) > 0 {
		sort.Strings(words)

		uniq := words[:1]
		for i := 1; i < len(words); i++ {
			if words[i] != words[i-1] {
				uniq = append(uniq, words[i])
			}
		}
		words = uniq
	}

	return matchWord(s, words)
}

func matchFile(s string) (matches []string, longest []rune) {
	dir := replaceTilde(s)

	if !filepath.IsAbs(dir) {
		wd, err := os.Getwd()
		if err != nil {
			log.Printf("getting current directory: %s", err)
		} else {
			dir = wd + string(filepath.Separator) + dir
		}
	}

	dir = filepath.Dir(unescape(dir))

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("reading directory: %s", err)
	}

	for _, f := range files {
		name := filepath.Join(dir, f.Name())
		f, err := os.Lstat(name)
		if err != nil {
			log.Printf("getting file information: %s", err)
			continue
		}

		name = strings.ToLower(escape(f.Name()))
		_, last := filepath.Split(s)
		if !strings.HasPrefix(name, strings.ToLower(last)) {
			continue
		}

		name = f.Name()
		if isRoot(s) || filepath.Base(s) != s {
			name = filepath.Join(filepath.Dir(unescape(s)), f.Name())
		}
		name = escape(replaceTilde(name))

		item := f.Name()
		if f.Mode().IsDir() {
			item += escape(string(filepath.Separator))
		}
		matches = append(matches, item)

		if longest == nil {
			if f.Mode().IsRegular() {
				longest = []rune(name + " ")
			} else {
				longest = []rune(name + escape(string(filepath.Separator)))
			}
		} else {
			longest = matchLongest(longest, []rune(name))
		}
	}

	if len(longest) < len([]rune(s)) {
		longest = []rune(s)
	}

	return
}

func matchCmd(s string) (matches []string, longest []rune) {
	words := make([]string, 0, len(gCmdWords)+len(gOpts.cmds))
	words = append(words, gCmdWords...)
	for c := range gOpts.cmds {
		words = append(words, c)
	}
	sort.Strings(words)
	j := 0
	for i := 1; i < len(words); i++ {
		if words[j] == words[i] {
			continue
		}
		j++
		words[i], words[j] = words[j], words[i]
	}
	words = words[:j+1]
	matches, longest = matchWord(s, words)
	return
}

func completeCmd(acc []rune) (matches []string, longestAcc []rune) {
	s := string(acc)
	f := tokenize(s)

	if len(f) == 1 {
		matches, longestAcc = matchCmd(s)
		return
	}

	longest := []rune(f[len(f)-1])

	switch f[0] {
	case "set":
		if len(f) == 2 {
			matches, longest = matchWord(f[1], gOptWords)
			break
		}
		if len(f) != 3 {
			break
		}
		switch f[1] {
		case "filtermethod", "searchmethod":
			matches, longest = matchWord(f[2], []string{"glob", "regex", "text"})
		case "selmode":
			matches, longest = matchWord(f[2], []string{"all", "dir"})
		case "sortby":
			matches, longest = matchWord(f[2], []string{"atime", "btime", "ctime", "custom", "ext", "name", "natural", "size", "time"})
		default:
			if slices.Contains(gOptWords, f[1]+"!") {
				matches, longest = matchWord(f[2], []string{"false", "true"})
			}
		}
	case "setlocal":
		if len(f) == 2 {
			matches, longest = matchFile(f[1])
			break
		}
		if len(f) == 3 {
			matches, longest = matchWord(f[2], gLocalOptWords)
			break
		}
		if len(f) != 4 {
			break
		}
		switch f[2] {
		case "sortby":
			matches, longest = matchWord(f[3], []string{"atime", "btime", "ctime", "custom", "ext", "name", "natural", "size", "time"})
		default:
			if slices.Contains(gLocalOptWords, f[2]+"!") {
				matches, longest = matchWord(f[3], []string{"false", "true"})
			}
		}
	case "map", "nmap", "vmap", "cmap":
		if len(f) == 3 {
			matches, longest = matchCmd(f[2])
		}
	case "cmd":
	case "toggle":
		matches, longest = matchFile(f[len(f)-1])
	case "cd", "select", "source":
		if len(f) == 2 {
			matches, longest = matchFile(f[1])
		}
	default:
		if !slices.Contains(gCmdWords, f[0]) {
			matches, longest = matchFile(f[len(f)-1])
		}
	}

	longestAcc = append(acc[:len(acc)-len([]rune(f[len(f)-1]))], longest...)
	return
}

func completeFile(acc []rune) (matches []string, longestAcc []rune) {
	s := string(acc)

	wd, err := os.Getwd()
	if err != nil {
		log.Printf("getting current directory: %s", err)
	}

	files, err := os.ReadDir(wd)
	if err != nil {
		log.Printf("reading directory: %s", err)
	}

	for _, f := range files {
		if !strings.HasPrefix(strings.ToLower(f.Name()), strings.ToLower(s)) {
			continue
		}

		matches = append(matches, f.Name())

		if longestAcc == nil {
			longestAcc = []rune(f.Name())
		} else {
			longestAcc = matchLongest(longestAcc, []rune(f.Name()))
		}
	}

	if len(longestAcc) < len(acc) {
		longestAcc = acc
	}

	return
}

func completeShell(acc []rune) (matches []string, longestAcc []rune) {
	s := string(acc)
	f := tokenize(s)

	var longest []rune

	switch len(f) {
	case 1:
		matches, longestAcc = matchExec(s)
	default:
		matches, longest = matchFile(f[len(f)-1])
		longestAcc = append(acc[:len(acc)-len([]rune(f[len(f)-1]))], longest...)
	}

	return
}
