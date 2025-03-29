package infrestructure

import (
    "Multi/src/core"
    "Multi/src/user/domain"
    "Multi/src/user/domain/entities"
    "log"
)

type Postgres struct {
    conn *core.Conn_Postgres
}

func NewPostgres() domain.UserRepository {
    conn := core.GetDBPool()
    if conn.Err != "" {
        log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
    }

    return &Postgres{conn: conn}
}

func (pg *Postgres) Create(user *entities.User) error {
    query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3)"
    _, err := pg.conn.ExecutePreparedQuery(query, user.Username, user.Password, user.Email)
    if err != nil {
        log.Printf("Error al insertar usuario: %v", err)
        return err
    }
    return nil
}

func (pg *Postgres) GetAll() ([]entities.User, error) {
    query := "SELECT id, username, password, email FROM users"
    rows := pg.conn.FetchRows(query)
    defer rows.Close()

    var users []entities.User

    for rows.Next() {
        var user entities.User
        if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
            log.Printf("Error al escanear usuario: %v", err)
            return nil, err
        }
        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

func (pg *Postgres) GetByID(id int) (*entities.User, error) {
    query := "SELECT id, username, password, email FROM users WHERE id = $1"
    rows := pg.conn.FetchRows(query, id)
    defer rows.Close()

    var user entities.User
    if rows.Next() {
        if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
            log.Printf("Error al escanear usuario: %v", err)
            return nil, err
        }
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &user, nil
}
func (pg *Postgres) GetByUsername(username string) (*entities.User, error) {
    query := "SELECT id, username, password, email FROM users WHERE username = $1"
    rows := pg.conn.FetchRows(query, username)
    defer rows.Close()

    var user entities.User
    if rows.Next() {
        if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
            log.Printf("Error al escanear usuario: %v", err)
            return nil, err
        }
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &user, nil
}