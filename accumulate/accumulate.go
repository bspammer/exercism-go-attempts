package accumulate

const testVersion = 1

func Accumulate(list []string, f func(string) string) []string {
	output := make([]string, len(list))
	for i, item := range list {
		output[i] = f(item)
	}
	return output
}
