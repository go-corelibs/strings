package strings

// Scan is a text scanner which looks for unquoted and unescaped `sep`
func Scan(src, sep string) (before, after string, found bool) {

	s := struct {
		quote   rune
		quoted  bool
		escaped bool
	}{}

	runes := []rune(src)
	total, size := len(runes), len(sep)

	for idx := 0; idx < total; idx++ {

		r := runes[idx]
		if r == '\\' {
			s.escaped = true
			continue
		}

		if IsQuote(r) {
			if s.quoted {
				if s.quote == r {
					s.quote = 0
					s.quoted = false
				}
			} else {
				s.quote = r
				s.quoted = true
			}
			continue
		} else if s.quoted {
			continue
		}

		if remainder := total - idx; size > remainder {
			break
		}

		if s.escaped {
			s.escaped = false
			continue
		}

		if found = src[idx:idx+size] == sep; found {
			before = src[:idx]
			after = src[idx+size:]
			return
		}

	}

	return src, "", false
}
