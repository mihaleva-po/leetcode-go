# LeetCode на Go

Решения задач с LeetCode в процессе изучения Go.

Каждая папка — отдельная задача: решение (`main.go`) и юнит-тесты (`main_test.go`).

## Задачи

| # | Задача | Сложность | Темы |
|---|--------|-----------|------|
| 1 | [Two Sum](./001-two-sum) | Easy | массивы, хеш-map |
| 2 | [Find if Path Exists in Graph](./002-find-if-path-exists-in-graph) | Easy | графы, DFS |
| 3 | [Sum of Compatible Numbers in Range I](./003-sum-of-compatible-numbers-in-range-i) | Easy | битовые операции, перебор |
| 4 | [Valid Sudoku](./004-valid-sudoku) | Medium | массивы, хеш-map |
| 5 | [Group Anagrams](./005-group-anagrams) | Medium | строки, хеш-map, сортировка |
| 6 | [Sort Colors](./006-sort-colors) | Medium | два указателя, in-place сортировка |
| 7 | [Car Fleet](./007-car-fleet) | Medium | стек, сортировка, математика |
| 8 | [Sort an Array](./008-sort-an-array) | Medium | сортировка, quicksort, рекурсия |
| 9 | [Valid Binary Strings with Cost Limit](./009-valid-binary-strings-with-cost-limit) | Medium | битовая генерация, перебор |

## Как запустить

Тесты для конкретной задачи:
```bash
cd 0001-two-sum
go test -v ./...
```

Тесты для всех задач сразу (из корня репозитория):
```bash
go test -v ./...
```

## Стек

Go 1.26+
