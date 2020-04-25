package delegateverify

type MyOrderModel struct {
	repo Repository
}

func NewMyOrderModel(repo Repository) OrderModel {
	return &MyOrderModel{repo}
}

func (m *MyOrderModel) Save(order Order) {
	// 不做檢查直接存
	// 只想練習當參數為函數時
	// 如何驗證 該函數的行為符合預期
	m.repo.Insert(order)
}

func (m *MyOrderModel) Delete(filter func(o Order) bool) {
	// 無實做
}
