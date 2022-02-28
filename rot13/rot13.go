package rot13

func Rotate(orig string) string {
	bs := []byte{}
	for _, ch := range []byte(orig) {
		bs = append(bs, rot13(ch))
	}
	return string(bs)
}

func rot13(b byte) byte {
	added := b + 13
	if 'a' <= b && b <= 'z' {
		if added > 'z' {
			added -= 26
		}
		return added
	} else {
		if added > 'Z' {
			added -= 26
		}
		return added
	}
}
