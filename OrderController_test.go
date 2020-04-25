package delegateverify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 模仿 91單元測試課程
// 由於沒有 go 的程式碼, 自己另外寫
// 刪除多餘的部份
// 只專注 class OrderController 的 deleteAmountMoreThan100 method
// 如何驗證 這個 Lambda 函數 o -> o.getAmount() > 100 是否寫錯
//
// 故意將 filter 函數改為 > 101
// 出現以下訊息
//
// Running tool: go test -timeout 30s delegateverify -v -gcflags=-l
// === RUN   TestOrderController_DeleteAmountMoreThan100
// --- FAIL: TestOrderController_DeleteAmountMoreThan100 (0.00s)
//     OrderController_test.go:36:
//         	Error Trace:	OrderController_test.go:36
//         	Error:      	Not equal:
//         	            	expected: true
//         	            	actual  : false
//         	Test:       	TestOrderController_DeleteAmountMoreThan100
//         	Messages:   	Amount > 臨界值(100), 應該被刪除
// FAIL
// FAIL	delegateverify	0.003s
// FAIL
// Error: Tests failed.
func TestOrderController_DeleteAmountMoreThan100(t *testing.T) {
	type setup struct {
		msg   string
		model *TestOrderModel
	}
	setups := []setup{
		{
			"Amount < 臨界值(100), 不應該被刪除",
			&TestOrderModel{order: Order{ID: 1, Amount: 99}, expect: false},
		},
		{
			"刪除價格的臨界值應該 100",
			&TestOrderModel{order: Order{ID: 1, Amount: 100}, expect: false},
		},
		{
			"Amount > 臨界值(100), 應該被刪除",
			&TestOrderModel{order: Order{ID: 1, Amount: 101}, expect: true},
		},
	}
	for _, s := range setups {
		controller := NewOrderController(s.model)
		controller.DeleteAmountMoreThan100()
		assert.Equal(t, s.model.expect, s.model.actual, s.msg)
	}
}

type TestOrderModel struct {
	order  Order
	expect bool
	actual bool
}

func (m *TestOrderModel) Save(order Order) {
	// 內部不實做, 單純為了滿足 interface
}

func (m *TestOrderModel) Delete(filter func(o Order) bool) {
	// 為了驗證 OrderController 的 DeleteAmountMoreThan100 Method
	m.actual = filter(m.order)
}
