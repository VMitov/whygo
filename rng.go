type Rand struct {
	...
}

// Returns random byte
func (rand *Rand) GetByte() byte {
	...
}

// Implements io.Reader
func (rand *Rand) Read(p []byte) (n int, err error) {
	for x := 0; x < len(p); x++ {
		p[x] = rand.GetByte()
	}

	return len(p), nil
}
