package repository

type WhereOp string

const (
	WhereOpEq  WhereOp = "EQ"
	WhereOpLt  WhereOp = "LT"
	WhereOpLte WhereOp = "LTE"
	WhereOpGt  WhereOp = "GT"
	WhereOpGte WhereOp = "GTE"
	WhereOpIn  WhereOp = "IN"
	WhereOpBtw WhereOp = "BTW"
	WhereOpRgx WhereOp = "REG"

	WhereOpNot WhereOp = "NOT"
	WhereOpAnd WhereOp = "AND"
	WhereOpOr  WhereOp = "OR"
)

type OrderDir string

const (
	OrderDirAsc  OrderDir = "ASC"
	OrderDirDesc OrderDir = "DESC"
)
