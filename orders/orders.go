// Package orders реалізує Розділ 2 домашньої роботи: рефакторинг
// OrderService через виокремлення мінімального інтерфейсу OrderStore
// та впровадження залежності через конструктор.
//
// Саме цей пакет ви й попросите ШІ-асистента рефакторити (Розділ 2)
// та покрити тестами (Розділ 3, Крок 2). Файл orders_test.go вже
// містить автотести з фейковим store — використайте їх, щоб
// перевірити свою реалізацію локально, а власний промпт і згенерований
// ШІ тест додайте окремо (наприклад, у orders_ai_test.go) для звіту.
package orders

import "fmt"

// OrderStore — мінімальний інтерфейс для залежності від БД
// (Завдання 2.1). Замінює прямий *sql.DB, дотримуючись приказки
// Роба Пайка: "чим більший інтерфейс, тим слабша абстракція".
type OrderStore interface {
	Exec(query string, args ...any) error
}

// OrderService — Завдання 2.2: сервіс приймає залежність через
// конструктор (dependency injection), а не створює її сам.
type OrderService struct {
	store OrderStore
}

// NewOrderService — конструктор із впровадженням залежності.
func NewOrderService(store OrderStore) *OrderService {
	return &OrderService{store: store}
}

// placeOrderSQL — параметризований запит: значення передаються як
// args, а не зшиваються в рядок (захист від SQL-ін'єкції).
const placeOrderSQL = "INSERT INTO orders (id, amount) VALUES ($1, $2)"

// PlaceOrder виконує вставку замовлення через injected store
// (Завдання 2.2). Сервіс нічого не знає про *sql.DB — він знає лише
// про один метод OrderStore.Exec.
func (s *OrderService) PlaceOrder(orderID string, amount float64) error {
	if err := s.store.Exec(placeOrderSQL, orderID, amount); err != nil {
		return fmt.Errorf("place order %q: %w", orderID, err)
	}
	return nil
}
