# Popular Go Frameworks and Libraries

This document provides a curated selection of popular and commonly used Go frameworks and libraries. It is not an exhaustive list, as the Go ecosystem is constantly evolving. This list aims to provide a starting point for exploring the rich landscape of Go tools.

| Category            | Library/Framework   | Description                                               | Repository                                                                   |
|---------------------|---------------------|-----------------------------------------------------------|------------------------------------------------------------------------------|
| **Web Frameworks**  | Gin                 | High-performance, minimalist web framework.               | [github.com/gin-gonic/gin](github.com/gin-gonic/gin)                         |
|                     | Echo                | Fast, extensible, and HTTP/2-ready.                       | [github.com/labstack/echo](github.com/labstack/echo)                         |
|                     | Fiber               | Express.js-inspired, fast and friendly.                   | [github.com/gofiber/fiber](github.com/gofiber/fiber)                         |
|                     | Chi                 | Lightweight, idiomatic HTTP router.                       | [github.com/go-chi/chi](github.com/go-chi/chi)                               |
|                     | Iris                | Fast, easy-to-learn, and efficient.                       | [github.com/kataras/iris](github.com/kataras/iris)                           |
|                     | Buffalo             | Helps you to quickly get a web project started.           | [github.com/gobuffalo/buffalo](github.com/gobuffalo/buffalo)                 |
|                     | Revel               | A high-productivity web framework.                        | [github.com/revel/revel](github.com/revel/revel)                             |
| **Databases/ORMs**  | GORM                | Popular ORM, supports various databases.                  | [github.com/go-gorm/gorm](github.com/go-gorm/gorm)                           |
|                     | sqlx                | Extensions to the standard `database/sql` package.        | [github.com/jmoiron/sqlx](github.com/jmoiron/sqlx)                           |
|                     | pgx                 | Pure Go PostgreSQL driver.                                | [github.com/jackc/pgx](github.com/jackc/pgx)                                 |
|                     | go-sql-driver/mysql | MySQL driver.                                             | [github.com/go-sql-driver/mysql](github.com/go-sql-driver/mysql)             |
|                     | mattn/go-sqlite3    | SQLite driver.                                            | [github.com/mattn/go-sqlite3](github.com/mattn/go-sqlite3)                   |
| **Microservices**   | Go kit              | Toolkit for microservices.                                | [github.com/go-kit/kit](github.com/go-kit/kit)                               |
|                     | Micro               | Platform for microservices.                               | [github.com/micro/micro](github.com/micro/micro)                             |
|                     | go-zero             | A web and microservices framework.                        | [github.com/zeromicro/go-zero](github.com/zeromicro/go-zero)                 |
| **CLI Development** | Cobra               | Library for creating powerful CLIs.                       | [github.com/spf13/cobra](github.com/spf13/cobra)                             |
|                     | urfave/cli          | Another popular CLI library.                              | [github.com/urfave/cli](github.com/urfave/cli)                               |
|                     | pflag               | POSIX/GNU style flag parsing.                             | [github.com/spf13/pflag](github.com/spf13/pflag)                             |
| **Env Variables**   | Viper               | Handles configuration files and environment variables.    | [github.com/spf13/viper](github.com/spf13/viper)                             |
|                     | godotenv            | Loads environment variables from `.env` files.            | [github.com/joho/godotenv](github.com/joho/godotenv)                         |
|                     | envconfig           | Reads configuration from environment variables.           | [github.com/kelseyhightower/envconfig](github.com/kelseyhightower/envconfig) |
| **Testing**         | testify             | Suite of testing tools.                                   | [github.com/stretchr/testify](github.com/stretchr/testify)                   |
|                     | goconvey            | Web UI for Go testing.                                    | [github.com/smartystreets/goconvey](github.com/smartystreets/goconvey)       |
|                     | ginkgo              | BDD-style testing framework.                              | [github.com/onsi/ginkgo](github.com/onsi/ginkgo)                             |
|                     | gomock              | Mocking framework.                                        | [github.com/golang/mock](github.com/golang/mock)                             |
| **Logging**         | logrus              | Structured logger.                                        | [github.com/sirupsen/logrus](github.com/sirupsen/logrus)                     |
|                     | zap                 | Fast, structured, leveled logger.                         | [github.com/uber-go/zap](github.com/uber-go/zap)                             |
|                     | zerolog             | Zero-allocation JSON logger.                              | [github.com/rs/zerolog](github.com/rs/zerolog)                               |
|                     | slog                | (Standard library) Structured logging.                    | (Part of the standard library)                                               |
| **Caching**         | go-cache            | In-memory key:value store/cache.                          | [github.com/patrickmn/go-cache](github.com/patrickmn/go-cache)               |
|                     | ristretto           | High-performance concurrent cache.                        | [github.com/dgraph-io/ristretto](github.com/dgraph-io/ristretto)             |
|                     | groupcache          | Caching and cache-population library.                     | [github.com/golang/groupcache](github.com/golang/groupcache)                 |
|                     | bigcache            | Efficient cache for gigabytes of data.                    | [github.com/allegro/bigcache](github.com/allegro/bigcache)                   |
| **Authentication**  | jwt-go              | JWT (JSON Web Token) library.                             | [github.com/golang-jwt/jwt](github.com/golang-jwt/jwt)                       |
|                     | casbin              | Authorization library.                                    | [github.com/casbin/casbin](github.com/casbin/casbin)                         |
|                     | scrypt              | Password hashing.                                         | [github.com/dchest/scrypt](github.com/dchest/scrypt)                         |
|                     | bcrypt              | Another password hashing library.                         | [github.com/go-crypt/bcrypt](github.com/go-crypt/bcrypt)                     |
| **Validation**      | validator           | Struct validation.                                        | [github.com/go-playground/validator](github.com/go-playground/validator)     |
|                     | ozzo-validation     | Another validation library.                               | [github.com/go-ozzo/ozzo-validation](github.com/go-ozzo/ozzo-validation)     |
| **Task Queues**     | queue               | A simple, in-memory queue.                                | [github.com/go-queue/queue](github.com/go-queue/queue)                       |
|                     | asynq               | A simple, reliable, and efficient distributed task queue. | [github.com/hibiken/asynq](github.com/hibiken/asynq)                         |