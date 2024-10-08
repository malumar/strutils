package strutils

import (
	"bufio"
	"io"
	"strings"
)

func SplitQuotedString(text string) (items []string) {

	var sb strings.Builder

	items = make([]string, 0)
	defer func() {
		if len(strings.TrimSpace(sb.String())) > 0 {
			items = append(items, sb.String())
		}
		if len(items) == 0 {
			items = nil
		}
	}()
	r := bufio.NewReader(strings.NewReader(text))
	openslash := false
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {

			switch c {
			case '\\':
				if !openslash {
					openslash = true
					if checkNext('"', r) {
						openslash = false
						sb.WriteString(`\"`)
						continue
					}
				}
			case '"':
				if !openslash {

					s, foundQuote := returnBeforeFoundQuote(r)
					if !foundQuote {
						s = sb.String() + `"` + s
						sb.Reset()
					}
					items = append(items, s)
					openslash = false
					continue
				}
				break
			case ' ':
				content := sb.String()
				if len(strings.TrimSpace(content)) > 0 {
					items = append(items, content)

				}
				sb.Reset()
				openslash = false
				continue
			default:
				openslash = false
			}
			sb.WriteRune(c)
		}
		openslash = false
	}

	return
}

func checkNext(expected rune, reader *bufio.Reader) bool {

	if r, _, err := reader.ReadRune(); err != nil {
		reader.UnreadRune()
		return false
	} else {
		if expected == r {
			return true
		} else {
			return false
		}
	}
}

func returnBeforeFoundQuote(reader *bufio.Reader) (string, bool) {
	var sb strings.Builder
	for {
		prevIsUnquote := false
		if r, _, err := reader.ReadRune(); err != nil {
			return sb.String(), false
		} else {
			switch r {
			case '"':
				if !prevIsUnquote {
					return sb.String(), true
				}
				break
			case '\\':
				prevIsUnquote = true
				break
			}

			sb.WriteRune(r)
		}
	}

	return sb.String(), false
}
