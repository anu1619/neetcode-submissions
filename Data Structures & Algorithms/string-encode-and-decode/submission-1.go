type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	out := ""
	for i := 0; i < len(strs); i++ {
		out += fmt.Sprintf("%s@#@", strs[i])
	}
	return out
}

func (s *Solution) Decode(encoded string) []string {
	out := make([]string, 0)
	prevIdx := 0
	for i := 0; i < len(encoded)-2; i++ {
		if encoded[i+1] == '#' && encoded[i+2] == '@' && encoded[i] == '@' {
			out = append(out, encoded[prevIdx:i])
			prevIdx = i + 3
		}

	}
	return out
}