package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	prefix := ""
	var r byte
	idx := 0 
	
	if len(strs) == 1 {
        return string(strs[0])
	} else {
for {
	for i, str := range strs {
			if len(str) == 0 {
					return ""
			}
			if idx >= len(str) {
			  return prefix
            } else if i == 0 {
                r = str[idx]
            } else if i == len(strs)-1 && r == str[idx] {
					prefix += string(r)
					idx++
			} else if r == str[idx] {
							i++
			} else {
					return prefix         
			}
			}
	}
	return prefix

	}	
}




	




 