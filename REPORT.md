# Звіт до домашньої роботи · Урок 3

## 1. Посилання на код

https://github.com/imperror/softserve_go_l3

## 2. Промпт з Розділу 2

Точний текст промпту зберігається в
[`prompts/interface-extraction-prompt.md`](./prompts/interface-extraction-prompt.md).

## 3. Міксовані приймачі та інтерфейси на стороні споживача

Дайте відповідь своїми словами (1-2 абзаци) на питання: *«Чому
міксування value та pointer ресиверів є небезпечним і як ШІ допомагає
проектувати інтерфейси на стороні споживача?»*

Value receiver - працює на копії , а pointer - на оригіналі, через це вони технічно працюють на різних обєктах , через це T та *T мають різні набори методів, Т - лише методи з value ресівером, тому будь який інтерфейс що має хоч один pointer метод - задовільняє лише *T. ШІ діючи в рамках існуючого проєкту може виявити такі проблеми та запропонувати вирішення, може робити mock тести для виявлення дефектів. 

## 4. Тест і мок, згенеровані ШІ (Розділ 3, Крок 2)

Окрім наданих автотестів, я попросив ШІ-агента згенерувати власний юніт-тест
і мок для `OrderService.PlaceOrder` — вони збережені окремим файлом
[`orders/orders_ai_test.go`](./orders/orders_ai_test.go): мок
`mockOrderStore` з фіксацією викликів `Exec` та compile-time перевірка
`var _ OrderStore = (*mockOrderStore)(nil)`. Згенеровані тести:
`TestPlaceOrder_ExecArgs`, `TestPlaceOrder_ErrorIsWrapped`,
`TestPlaceOrder_UsesOnlyInjectedStore`.
Перевірка: `go test ./orders/... -v -cover` — 6 тестів PASS, покриття 100 %.

## 5. Скріншоти / відео виконання

https://docs.google.com/document/d/1jh-IA8gnvkNGGJNtadgczcvh74YLdXpXOFXeAwuvYWY/edit?usp=sharing