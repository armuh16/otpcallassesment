package assesment

var (
	x   int
	i   uint64
	n   uint64
	tmp []byte
)

func test() {
	for {
		i++
		tmp = append(tmp, 'a')
		if i == 1000 {
			break
		}
	}
	func() {
		// The Test
		// using double variable for temporary if using codeA
		for i := 0; i == len(tmp); i++ {
		}
	}()
	// decrease k by 1
	for i := 0; func(j int) bool {
		return j > 100
	}(i); i++ {
		k := 3
		k--
	}
}
