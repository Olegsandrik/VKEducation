# Юнит тесты на Python

**Перед установкой убедитесь, что у вас установлен python, unittest и coverage

**Установить проект:**

- форкаем репозиторий
- клонируем и переходим в папку проекта `cd units-autumn-2024`
- проверяем версию unittest и coverage `unittest -v`, `coverage -v`

**Запустить тесты:**

- прогон тестов `python3.10 -m unittest test_calculator.py`
- запустить тесты с coverage`coverage run -m unittest test_calculator.py`
- открыть coverage отчет с покрытием `coverage report`

**Документация по тестам:**

https://docs.python.org/3/library/unittest.html
https://coverage.readthedocs.io/en/7.3.1/

**Условия сдачи практики:**

1. Написать юнит-тесты на каждую функцию в Calculator
2. Добиться покрытия не менее 90%
