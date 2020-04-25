package delegateverify

type OrderModel interface {
	Save(o Order)
	Delete(func(o Order) bool)
}
