# Популярні Go фреймворки та бібліотеки

Цей документ містить підбірку популярних та часто використовуваних Go фреймворків та бібліотек. Це не вичерпний список, оскільки екосистема Go постійно розвивається. Цей список має на меті стати відправною точкою для дослідження багатого ландшафту інструментів Go.

| Категорія           | Бібліотека/Фреймворк | Опис                                                   | Репозиторій                                                                  |
|---------------------|----------------------|--------------------------------------------------------|------------------------------------------------------------------------------|
| **Веб-фреймворки**  | Gin                  | Високопродуктивний, мінімалістичний веб-фреймворк.     | [github.com/gin-gonic/gin](github.com/gin-gonic/gin)                         |
|                     | Echo                 | Швидкий, розширюваний та готовий до HTTP/2.            | [github.com/labstack/echo](github.com/labstack/echo)                         |
|                     | Fiber                | Натхненний Express.js, швидкий та дружній.             | [github.com/gofiber/fiber](github.com/gofiber/fiber)                         |
|                     | Chi                  | Легкий, ідіоматичний HTTP-маршрутизатор.               | [github.com/go-chi/chi](github.com/go-chi/chi)                               |
|                     | Iris                 | Швидкий, легкий у вивченні та ефективний.              | [github.com/kataras/iris](github.com/kataras/iris)                           |
|                     | Buffalo              | Допомагає швидко розпочати веб-проєкт.                 | [github.com/gobuffalo/buffalo](github.com/gobuffalo/buffalo)                 |
|                     | Revel                | Високопродуктивний веб-фреймворк.                      | [github.com/revel/revel](github.com/revel/revel)                             |
| **Бази даних/ORM**  | GORM                 | Популярний ORM, підтримує різні бази даних.            | [github.com/go-gorm/gorm](github.com/go-gorm/gorm)                           |
|                     | sqlx                 | Розширення для стандартного пакета `database/sql`.     | [github.com/jmoiron/sqlx](github.com/jmoiron/sqlx)                           |
|                     | pgx                  | Чистий драйвер PostgreSQL на Go.                       | [github.com/jackc/pgx](github.com/jackc/pgx)                                 |
|                     | go-sql-driver/mysql  | Драйвер MySQL.                                         | [github.com/go-sql-driver/mysql](github.com/go-sql-driver/mysql)             |
|                     | mattn/go-sqlite3     | Драйвер SQLite.                                        | [github.com/mattn/go-sqlite3](github.com/mattn/go-sqlite3)                   |
| **Мікросервіси**    | Go kit               | Інструментарій для мікросервісів.                      | [github.com/go-kit/kit](github.com/go-kit/kit)                               |
|                     | Micro                | Платформа для мікросервісів.                           | [github.com/micro/micro](github.com/micro/micro)                             |
|                     | go-zero              | Веб-фреймворк та фреймворк для мікросервісів.          | [github.com/zeromicro/go-zero](github.com/zeromicro/go-zero)                 |
| **Розробка CLI**    | Cobra                | Бібліотека для створення потужних CLI.                 | [github.com/spf13/cobra](github.com/spf13/cobra)                             |
|                     | urfave/cli           | Ще одна популярна бібліотека CLI.                      | [github.com/urfave/cli](github.com/urfave/cli)                               |
|                     | pflag                | Розбір флагов у стилі POSIX/GNU.                       | [github.com/spf13/pflag](github.com/spf13/pflag)                             |
| **Змінні оточення** | Viper                | Обробка файлів конфігурації та змінних оточення.       | [github.com/spf13/viper](github.com/spf13/viper)                             |
|                     | godotenv             | Завантаження змінних оточення з файлів `.env`.         | [github.com/joho/godotenv](github.com/joho/godotenv)                         |
|                     | envconfig            | Читання конфігурації зі змінних оточення.              | [github.com/kelseyhightower/envconfig](github.com/kelseyhightower/envconfig) |
| **Тестування**      | testify              | Набір інструментів для тестування.                     | [github.com/stretchr/testify](github.com/stretchr/testify)                   |
|                     | goconvey             | Веб-інтерфейс для тестування на Go.                    | [github.com/smartystreets/goconvey](github.com/smartystreets/goconvey)       |
|                     | ginkgo               | Фреймворк для тестування в стилі BDD.                  | [github.com/onsi/ginkgo](github.com/onsi/ginkgo)                             |
|                     | gomock               | Фреймворк для створення моків.                         | [github.com/golang/mock](github.com/golang/mock)                             |
| **Ведення журналу** | logrus               | Структурований логер.                                  | [github.com/sirupsen/logrus](github.com/sirupsen/logrus)                     |
|                     | zap                  | Швидкий, структурований, рівневий логер.               | [github.com/uber-go/zap](github.com/uber-go/zap)                             |
|                     | zerolog              | Логер JSON з нульовим виділенням пам'яті.              | [github.com/rs/zerolog](github.com/rs/zerolog)                               |
|                     | slog                 | (Стандартна бібліотека) Структуроване ведення журналу. | (Частина стандартної бібліотеки)                                             |
| **Кешування**       | go-cache             | Сховище/кеш ключів-значень в пам'яті.                  | [github.com/patrickmn/go-cache](github.com/patrickmn/go-cache)               |
|                     | ristretto            | Високопродуктивний конкурентний кеш.                   | [github.com/dgraph-io/ristretto](github.com/dgraph-io/ristretto)             |
|                     | groupcache           | Бібліотека для кешування та заповнення кешу.           | [github.com/golang/groupcache](github.com/golang/groupcache)                 |
|                     | bigcache             | Ефективний кеш для гігабайтів даних.                   | [github.com/allegro/bigcache](github.com/allegro/bigcache)                   |
| **Аутентифікація**  | jwt-go               | Бібліотека JWT (JSON Web Token).                       | [github.com/golang-jwt/jwt](github.com/golang-jwt/jwt)                       |
|                     | casbin               | Бібліотека авторизації.                                | [github.com/casbin/casbin](github.com/casbin/casbin)                         |
|                     | scrypt               | Хешування паролів.                                     | [github.com/dchest/scrypt](github.com/dchest/scrypt)                         |
|                     | bcrypt               | Ще одна бібліотека для хешування паролів.              | [github.com/go-crypt/bcrypt](github.com/go-crypt/bcrypt)                     |
| **Валідація**       | validator            | Перевірка структур.                                    | [github.com/go-playground/validator](github.com/go-playground/validator)     |
|                     | ozzo-validation      | Ще одна бібліотека валідації.                          | [github.com/go-ozzo/ozzo-validation](github.com/go-ozzo/ozzo-validation)     |
| **Черги завдань**   | queue                | Проста черга в пам'яті.                                | [github.com/go-queue/queue](                                                 |