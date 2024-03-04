package test

func Test(msg []byte) []byte {
	val := string(msg) + "output"
	return []byte(val)
}
