func isAnagram(s string, t string) bool {
	countS := make(map[string]int)
	countT := make(map[string]int)
	for _,c := range s {
		countS[string(c)]++
	}
	for _, c := range t{
		countT[string(c)]++
	}

	if len(countS) != len(countT){
		return false
	}
	
	for k,v := range countS{
		if countT[k] != v{
			return false
		}
	}
	return true
}
