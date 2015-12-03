package multibar

func chunkInt64(v int64) chan int {
	var maxInt int
	maxInt = 1<<31 - 1
	maxIntCast := int64(maxInt)

	chunks := make(chan int)

	go func() {
		q, r := v/maxIntCast, v%maxIntCast
		for i := int64(0); i < q; i += 1 {
			chunks <- maxInt
		}

		if r > 0 {
			chunks <- int(r)
		}

		close(chunks)
	}()

	return chunks
}
