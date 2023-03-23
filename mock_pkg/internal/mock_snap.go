package internal

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/golang/mock/gomock"
)

var _ gomock.Matcher = (*mySnaps)(nil)

type mySnaps struct {
	t *testing.T
}

func EqSnap(t *testing.T) *mySnaps {
	return &mySnaps{t}
}

func (s mySnaps) Matches(data any) bool {
	// this might not be even needed
	s.t.Helper()
	snaps.MatchSnapshot(s.t, data)
	return true
}

func (s mySnaps) String() string {
	return "my-snaps"
}
