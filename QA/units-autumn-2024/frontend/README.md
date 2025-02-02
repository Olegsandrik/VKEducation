# Юнит тесты на TypeSciprt

**Перед установкой убедитесь, что у вас установлен npm**

**Установить проект:**

- форкаем репозиторий
- клонируем и переходим в папку проекта `cd units-spring-2024`
- проверяем версию ноды `node -v`. подходят 16, 18 версии, если ниже - лучше обновите.
- выполняем `npm ci`

**Запустить проект:**

- `npm start`

**Запустить тесты:**

- прогон тестов `npm test`
- запустить тесты в режиме отслеживания изменений `npm run test:watch`
- открыть html отчет с покрытием `npm run open:coverage`
- обновить снапшоты `npm run test:update-snapshots`
- линтер `npm run lint`
- линтер с фиксом ошибок `npm run lint --fix`

**Документация по тестам:**

https://create-react-app.dev/docs/running-tests

_jest:_

https://jestjs.io/docs/ru/getting-started

https://jestjs.io/docs/ru/expect

_jest моки:_

https://jestjs.io/docs/ru/bypassing-module-mocks

_react testing library_

https://testing-library.com/docs/react-testing-library/example-intro

https://www.leighhalliday.com/introduction-react-testing-library

**Условия сдачи практики:**

1. Написать юнит-тесты на каждую функцию в utils
    - используйте [test.each](https://jestjs.io/docs/en/api#testeachtablename-fn-timeout) для одинаковых выходных значений
2. Написать юнит тесты на главную страницу (pages/MainPage), подумайте как решить проблему с непостоянством текущего времени
3. Написать юнит тесты на хуки в папке [hooks](./src/hooks)
4. Написать юнит тесты на компонент в [ProductCard](./src/components/ProductCard) 
5. Добиться покрытия не менее 90% по всем модулям
