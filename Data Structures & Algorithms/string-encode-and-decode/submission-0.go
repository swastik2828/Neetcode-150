
type Solution struct{}

func (s *Solution) Encode(strs []string) string {
		var result strings.Builder

		for _, str := range strs {
			result.WriteString(strconv.Itoa(len(str)))
			result.WriteByte('#')
			result.WriteString(str)
		}

	return result.String()
}

func (s *Solution) Decode(encoded string) []string {
	var result[]string
	for i:=0;i<len(encoded);{
		j := i
		for encoded[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		start := j+1
		result = append(result, encoded[start:start+length])
		i = start + length
	}
	return result
}
