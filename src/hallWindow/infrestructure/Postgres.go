package infrestructure

import (
    "Multi/src/core"
    "Multi/src/hallWindow/domain"
    "Multi/src/hallWindow/domain/entities"
    "log"
)

type Postgres struct {
    conn *core.Conn_Postgres
}

func NewPostgres() domain.HallWindowRepository {
    conn := core.GetDBPool()
    if conn.Err != "" {
        log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
    }

    return &Postgres{conn: conn}
}

func (pg *Postgres) Create(data *entities.HallWindow) error {
    query := "INSERT INTO HallWindow (date, status) VALUES ($1, $2)"
    _, err := pg.conn.ExecutePreparedQuery(query, data.Date, data.Status)
    if err != nil {
        log.Printf("Error al insertar datos de HallWindow: %v", err)
        return err
    }
    return nil
}

func (pg *Postgres) GetAll() ([]entities.HallWindow, error) {
    query := "SELECT id, date, status FROM HallWindow"
    rows := pg.conn.FetchRows(query)
    defer rows.Close()

    var hallWindowData []entities.HallWindow

    for rows.Next() {
        var data entities.HallWindow
        if err := rows.Scan(&data.ID, &data.Date, &data.Status); err != nil {
            log.Printf("Error al escanear datos de HallWindow: %v", err)
            return nil, err
        }
        hallWindowData = append(hallWindowData, data)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return hallWindowData, nil
}

func (pg *Postgres) GetByID(id int) (*entities.HallWindow, error) {
    query := "SELECT id, date, status FROM HallWindow WHERE id = $1"
    rows := pg.conn.FetchRows(query, id)
    defer rows.Close()

    var data entities.HallWindow
    if rows.Next() {
        if err := rows.Scan(&data.ID, &data.Date, &data.Status); err != nil {
            log.Printf("Error al escanear datos de HallWindow: %v", err)
            return nil, err
        }
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &data, nil
}

func (pg *Postgres) UpdateStatus(id int, status int) error {
    query := "UPDATE HallWindow SET status = $1 WHERE id = $2"
    _, err := pg.conn.ExecutePreparedQuery(query, status, id)
    if err != nil {
        log.Printf("Error al actualizar el estado de HallWindow: %v", err)
        return err
    }
    return nil
}