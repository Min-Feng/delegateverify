package delegateverify

type OrderController struct {
	orderModel OrderModel
}

func NewOrderController(model OrderModel) *OrderController {
	return &OrderController{
		orderModel: model,
	}
}

func (c *OrderController) Save(o Order) {
	c.orderModel.Save(o)
}

func (c *OrderController) DeleteAmountMoreThan100() {
	filter := func(o Order) bool { return o.Amount > 100 }
	c.orderModel.Delete(filter)
}
