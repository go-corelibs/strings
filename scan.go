package strings

// Scan is a text scanner which looks for unquoted and unescaped `sep`
func Scan(src, sep string) (before, after string, found bool) {

	s := struct {
		quote  rune
		quoted bool
	}{}

	runes := []rune(src)
	total, size := len(runes), len(sep)

	for idx := 0; idx < total; idx++ {

		r := runes[idx]

		if r == '\\' {
			// next character is escaped, skip
			idx += 1
			continue
		} else if IsQuote(r) {
			// this character is a double, single or backtick quotation detected
			if s.quoted {
				// scanning within a quoted string
				if s.quote == r {
					// this character is the ending quotation
					s.quote = 0
					s.quoted = false
				}
				// nothing to do with quoted contents
				continue
			}
			// this character is a starting quotation
			s.quote = r
			s.quoted = true
			continue
		} else if s.quoted {
			// nothing to do with quoted contents
			continue
		} else if remainder := total - idx; size > remainder {
			// early out, not enough characters for sep matching
			break
		} else if found = src[idx:idx+size] == sep; found {
			// sep match found, Scan complete
			before = src[:idx]
			after = src[idx+size:]
			return
		}

	}

	return src, "", false
}
