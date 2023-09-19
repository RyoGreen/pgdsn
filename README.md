# pgdsn

A simple Go package for PostgreSQL generating DSN strings.

## Instalation:

```
go get github.com/RyoGreen/pgdsn
```

## Usage Example
```
package main

import (
	"database/sql"
	"log"

	"github.com/RyoGreen/pgdsn"
	_ "github.com/lib/pq"
)

func main() {
	c := &pgdsn.Config{
		User:     "testuser",
		DbName:   "testdb",
		Password: "testpass",
		Host:     "localhost",
		Port:     5432,
		SslMode:  pgdsn.Disable,
	}
	db, err := sql.Open("postgres", c.FormatDSN())
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Println(err)
		return
	}
}
```

## SSL Modes
This package supports the following SSL modes for PostgreSQL connections:

- disable: No SSL encryption is used.
- require: SSL encryption is required.
- verify-ca: SSL encryption with certificate verification.
- verify-full: SSL encryption with full certificate verification.

## License
MIT
