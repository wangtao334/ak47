package util

func Find(vars map[string]string, k string) string {
	n := len(k)
	if n > 3 && k[0] == '$' && k[1] == '{' && k[n-1] == '}' {
		if v, ok := vars[k[2:len(k)-1]]; ok {
			return v
		}
	}
	return k
}
