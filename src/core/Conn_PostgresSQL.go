package core

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "sync"
)

type Conn_Postgres struct {
    Db  *sql.DB
    Err string
}

func (c *Conn_Postgres) ExecuteQuery(query string, id int) (any, any) {
    panic("unimplemented")
}

var (
    instance *Conn_Postgres
    once     sync.Once
)

func GetDBPool() *Conn_Postgres {
    once.Do(func() {
    	dsn := "user=max password= dbname=Domotica host=localhost port=5432 sslmode=disable"
        db, err := sql.Open("postgres", dsn)
        if err != nil {
            log.Fatalf("Error al conectar con PostgreSQL: %v", err)
        }

        if err = db.Ping(); err != nil {
            log.Fatalf("Error al hacer ping a PostgreSQL: %v", err)
        }

        instance = &Conn_Postgres{
            Db:  db,
            Err: "",
        }
    })

    return instance
}

func (c *Conn_Postgres) ExecutePreparedQuery(query string, args ...interface{}) (sql.Result, error) {
    stmt, err := c.Db.Prepare(query)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    return stmt.Exec(args...)
}

func (c *Conn_Postgres) FetchRows(query string, args ...interface{}) *sql.Rows {
    rows, err := c.Db.Query(query, args...)
    if err != nil {
        log.Fatalf("Error al ejecutar la consulta: %v", err)
    }
    return rows
}