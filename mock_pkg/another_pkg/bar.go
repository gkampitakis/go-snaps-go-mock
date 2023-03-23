package anotherpkg

type Bar interface {
	Foo(x int) int
}

func SUT(f Bar) {
	f.Foo(99)
}
