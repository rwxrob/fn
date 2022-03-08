package fn

// Reverse reverses the order of any slice passed to it.
func Reverse[T any](list []T) chan T {
	ret := make(chan T)
	go func() {
		for i, _ := range list {
			ret <- list[len(list)-1-i]
		}
		close(ret)
	}()
	return ret
}
