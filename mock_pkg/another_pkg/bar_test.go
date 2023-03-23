package anotherpkg

import (
	"testing"

	"go-test/mock_pkg/internal"

	"github.com/golang/mock/gomock"
)

func TestBar(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockBar(ctrl)

	// Asserts that the first and only call to Bar() is passed 99.
	// Anything else will fail.
	m.
		EXPECT().
		Foo(internal.EqSnap(t)).
		Return(101)

	SUT(m)
}
