package delegateverify

type Repository interface {
	IsExist(o Order) bool
	Insert(o Order)
}
