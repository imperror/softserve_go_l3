# Промпт для виокремлення інтерфейсу (Розділ 2)

Вставте нижче точний текст промпту, який ви розробили для Розділу 2
домашньої роботи. Промпт має явно містити:

- **роль** ШІ (наприклад, "Ти — досвідчений Go-розробник, що
  спеціалізується на чистій архітектурі...");
- **контекст** — наведений у завданні код `OrderService`;
- **завдання** — виокремити мінімальний інтерфейс `OrderStore`,
  переписати `OrderService` на dependency injection, згенерувати мок;
- **обмеження** — маленький інтерфейс (приказка Роба Пайка),
  конкретний стиль коду, тощо;
- **бажаний формат виводу** (наприклад, "поверни лише Go-код у трьох
  блоках: інтерфейс, сервіс, мок").



## Мій промпт

You're an experienced Go Software developer, specialyzing on code review , clean architectre, please take a look at ./orders and dedicate minimal interface 'OrderStore' , rewrite OrderService into a dependency injection, generate mock. Do the minimal required changes, create minimal effective diff required too meet task goal. Idiomatic go code. 