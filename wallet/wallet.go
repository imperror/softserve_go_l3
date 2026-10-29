// Package wallet реалізує Розділ 3, Крок 1 домашньої роботи:
// діагностику та виправлення класичної пастки Go — змішування
// приймача-значення і приймача-вказівника на одному типі.
package wallet

// SecureWallet — той самий тип, що й у прикладі з домашнього
// завдання. Обидва методи тепер мають приймач-вказівник: див.
// Balance() нижче.
type SecureWallet struct {
	balance float64
}

// Balance читає стан. Приймач — вказівник: золоте правило Go — якщо
// хоч один метод типу потребує *T, то й решта методів має бути на *T
// (консистентні набори методів + жодних прихованих копій структури).
func (w *SecureWallet) Balance() float64 {
	return w.balance
}

func (w *SecureWallet) Deposit(amt float64) {
	w.balance += amt
}

// ApplyDeposits додає amt до кожного гаманця у зрізі wallets.
//
// Це — саме той сценарій із заняття: wallets це []SecureWallet
// (зріз ЗНАЧЕНЬ, а не вказівників). Наївний `for _, w := range
// wallets { w.Deposit(amt) }` НЕ подіє, бо w — копія елемента
// циклу, тому тут індексуємо зріз напряму.
func ApplyDeposits(wallets []SecureWallet, amt float64) {
	// Індексуємо зріз напряму: wallets[i] — addressable, тому
	// wallets[i].Deposit(amt) компілюється у (&wallets[i]).Deposit(amt),
	// тобто приймач вказує на справжній елемент зрізу, а не на копію.
	for i := range wallets {
		wallets[i].Deposit(amt)
	}
}
