//-----------------------------------------------------------------------------
/*

 */
//-----------------------------------------------------------------------------

package sdf

//-----------------------------------------------------------------------------

type Triangle struct {
	V [3]V3
}

//-----------------------------------------------------------------------------

func NewTriangle(a, b, c V3) *Triangle {
	t := Triangle{}
	t.V[0] = a
	t.V[1] = b
	t.V[2] = c
	return &t
}

//-----------------------------------------------------------------------------

func (t *Triangle) Normal() V3 {
	e1 := t.V[1].Sub(t.V[0])
	e2 := t.V[2].Sub(t.V[0])
	return e1.Cross(e2).Normalize()
}

//-----------------------------------------------------------------------------
