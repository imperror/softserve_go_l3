package orders

import (
	"errors"
	"strings"
	"testing"
)

// mockOrderStore — згенерований мок (Завдання 2.3): мінімальна
// заміна OrderStore без реальної БД. Живе в _test.go, тому не
// потрапляє в продакшн-бінарник і не роздуває публічний API пакету.
//
// Мок записує кожен виклик (query + args) і вміє ін'єктувати
// помилку — цього достатньо, бо інтерфейс має рівно один метод.
type mockOrderStore struct {
	calls []execCall
	err   error
}

// execCall — один зафіксований виклик Exec.
type execCall struct {
	query string
	args  []any
}

// Compile-time перевірка: мок задовольняє контракт споживача.
var _ OrderStore = (*mockOrderStore)(nil)

func (m *mockOrderStore) Exec(query string, args ...any) error {
	m.calls = append(m.calls, execCall{query: query, args: args})
	return m.err
}

// execs повертає кількість викликів Exec.
func (m *mockOrderStore) execs() int { return len(m.calls) }

// TestPlaceOrder_ExecArgs перевіряє контракт із store: рівно один
// Exec, значення йдуть аргументами (не в SQL-рядок), у правильному
// порядку та з правильними типами.
func TestPlaceOrder_ExecArgs(t *testing.T) {
	tests := []struct {
		name    string
		orderID string
		amount  float64
	}{
		{"дробна сума", "order-1", 42.5},
		{"ціла сума", "order-2", 10},
		{"порожній ID все одно прокидається", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &mockOrderStore{}
			svc := NewOrderService(store)

			if err := svc.PlaceOrder(tt.orderID, tt.amount); err != nil {
				t.Fatalf("PlaceOrder: неочікувана помилка: %v", err)
			}
			if got := store.execs(); got != 1 {
				t.Fatalf("Exec викликано %d разів, очікувалось 1", got)
			}

			call := store.calls[0]
			if call.query == "" {
				t.Fatal("порожній запит (query)")
			}
			// Значення не повинні зшиватися в SQL — інакше SQL-ін'єкція.
			if tt.orderID != "" && strings.Contains(call.query, tt.orderID) {
				t.Errorf("query міг би містити SQL-ін'єкцію: %q", call.query)
			}
			if len(call.args) != 2 {
				t.Fatalf("args = %v, очікувалось 2 аргументи", call.args)
			}
			if call.args[0] != tt.orderID {
				t.Errorf("args[0] = %v, очікувалось %q", call.args[0], tt.orderID)
			}
			amt, ok := call.args[1].(float64)
			if !ok || amt != tt.amount {
				t.Errorf("args[1] = %v (%T), очікувалось %v (float64)", call.args[1], call.args[1], tt.amount)
			}
		})
	}
}

// TestPlaceOrder_ErrorIsWrapped перевіряє, що помилка від store не
// губиться й не підміняється новою: %w зберігає ланцюжок, тож
// споживач може порівняти її через errors.Is.
func TestPlaceOrder_ErrorIsWrapped(t *testing.T) {
	sentinel := errors.New("constraint violation")
	store := &mockOrderStore{err: sentinel}
	svc := NewOrderService(store)

	err := svc.PlaceOrder("order-boom", 7)
	if err == nil {
		t.Fatal("PlaceOrder мав повернути помилку")
	}
	if !errors.Is(err, sentinel) {
		t.Fatalf("помилка %v не загортає початкову %v", err, sentinel)
	}
	if !strings.Contains(err.Error(), "order-boom") {
		t.Errorf("у помилці бракує контексту (orderID): %v", err)
	}
}

// TestPlaceOrder_UsesOnlyInjectedStore перевіряє, що два сервіси не
// ділять стан: кожен звертається саме до свого injected store, а не
// до глобальної змінної.
func TestPlaceOrder_UsesOnlyInjectedStore(t *testing.T) {
	first, second := &mockOrderStore{}, &mockOrderStore{}

	for _, tc := range []struct {
		store   *mockOrderStore
		orderID string
	}{{first, "a"}, {second, "b"}} {
		if err := NewOrderService(tc.store).PlaceOrder(tc.orderID, 1); err != nil {
			t.Fatalf("PlaceOrder(%q): %v", tc.orderID, err)
		}
	}

	if first.execs() != 1 || second.execs() != 1 {
		t.Fatalf("кількість викликів: first=%d, second=%d, очікувалось 1 і 1", first.execs(), second.execs())
	}
	if first.calls[0].args[0] != "a" || second.calls[0].args[0] != "b" {
		t.Error("сервіс звертався не до свого store")
	}
}
