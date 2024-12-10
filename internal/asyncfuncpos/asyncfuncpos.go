package asyncfuncpos

import (
	"go/token"
	gotypes "go/types"
)

type AsyncFuncPos struct {
	Scope   *gotypes.Scope
	CallPos token.Pos
}

type Poses []*AsyncFuncPos

func (p Poses) Len() int {
	return len(p)
}

func (p Poses) Less(i, j int) bool {
	return p[i].Scope.Pos() < p[j].Scope.Pos()
}

func (p Poses) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Poses) Contains(pos token.Pos) (*AsyncFuncPos, bool) {
	if len(p) == 0 {
		return nil, false
	}
	mid := len(p) / 2

	if p[mid].Scope.Contains(pos) {
		return p[0], true
	}

	left := p[:mid]
	if l := len(left); l > 0 && left[0].Scope.Pos() <= pos && left[l-1].Scope.End() >= pos {
		return left.Contains(pos)
	}

	right := p[mid+1:]
	if l := len(right); l > 0 && right[0].Scope.Pos() <= pos && right[l-1].Scope.End() >= pos {
		return right.Contains(pos)
	}

	return nil, false
}
