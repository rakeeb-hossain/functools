package functools

type Summable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Sum[S Summable](slice []S) S {
	var res S
	for _, v := range slice {
		res += v
	}
	return res
}
