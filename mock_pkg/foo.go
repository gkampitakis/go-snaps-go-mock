package mockpkg

type Foo interface {
	Bar(x int) int
}

func SUT(f Foo) {
	f.Bar(55)
}
