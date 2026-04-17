type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var final string
	for i:=0; i<len(strs); i++{
		final =  final + strconv.Itoa(len(strs[i])) + "#" + strs[i]
	}

	return final
}

func (s *Solution) Decode(encoded string) []string {
	ans := []string{}
	start :=0
	end :=0
	for end<len(encoded){
		end = start
		for encoded[end] != '#' {
			end++
		}

		lenght , _ := strconv.Atoi(encoded[start:end])

		start = end +1
		end = start + lenght
		ans = append(ans , encoded[start:end])
		start = end
	}
	return ans
}
