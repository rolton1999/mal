package types

type BaseType interface{}

type List struct {
	Elements []BaseType
}

type Scalar struct {
	BaseType
}

type Atom struct {
	Scalar
}

type Integer struct {
	Val int
}
type String struct {
	Val string
}
