package vehicle

type IElement interface {
	Accept(v IVisitor)
	TypeName() string
}
