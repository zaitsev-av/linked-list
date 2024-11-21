# linked-list
- HasNext ✅  
Проверяет, есть ли следующий элемент в списке.
Возвращает true, если следующий узел существует, и false, если текущий узел последний.

- Next ✅
  Переходит к следующему узлу и возвращает его значение.
Если следующего узла нет, может вернуть ошибку или специальное значение, я возвращаю nil.

- HasPrevious (для двунаправленного списка) ❌ 
Проверяет, есть ли предыдущий узел.
Актуально только для двунаправленных списков.

- Previous (для двунаправленного списка) ❌ 
Возвращает значение предыдущего узла и перемещает итератор назад.

- Current ❌ 
Возвращает значение текущего узла без изменения положения итератора.

- Reset ❌ 
Сбрасывает итератор на начальное положение (первый элемент списка).
Удобно для повторного прохождения списка.

- Insert ❌ 
Вставляет новый элемент в список на текущую позицию итератора.

- Remove ❌ 
Удаляет текущий узел и корректно обновляет ссылки списка.

- SetValue ❌ 
Изменяет значение текущего узла.

- Seek (опционально) ❌ 
Перемещает итератор на определённую позицию, например, на узел с заданным индексом.