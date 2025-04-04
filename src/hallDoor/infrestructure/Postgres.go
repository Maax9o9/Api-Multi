package infrestructure

import (
    "Multi/src/core"
    "Multi/src/hallDoor/domain"
    "Multi/src/hallDoor/domain/entities"
    "log"
)

type Postgres struct {
    conn *core.Conn_Postgres
}

func NewPostgres() domain.HallDoorRepository {
    conn := core.GetDBPool()
    if conn.Err != "" {
        log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
    }

    return &Postgres{conn: conn}
}

func (pg *Postgres) Create(data *entities.HallDoor) error {
    query := "INSERT INTO HallDoor (date, status) VALUES ($1, $2)"
    _, err := pg.conn.ExecutePreparedQuery(query, data.Date, data.Status)
    if err != nil {
        log.Printf("Error al insertar datos de HallDoor: %v", err)
        return err
    }
    return nil
}

func (pg *Postgres) GetAll() ([]entities.HallDoor, error) {
    query := "SELECT id, date, status FROM HallDoor"
    rows := pg.conn.FetchRows(query)
    defer rows.Close()

    var hallDoorData []entities.HallDoor

    for rows.Next() {
        var data entities.HallDoor
        if err := rows.Scan(&data.ID, &data.Date, &data.Status); err != nil {
            log.Printf("Error al escanear datos de HallDoor: %v", err)
            return nil, err
        }
        hallDoorData = append(hallDoorData, data)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return hallDoorData, nil
}

func (pg *Postgres) GetByID(id int) (*entities.HallDoor, error) {
    query := "SELECT id, date, status FROM HallDoor WHERE id = $1"
    rows := pg.conn.FetchRows(query, id)
    defer rows.Close()

    var data entities.HallDoor
    if rows.Next() {
        if err := rows.Scan(&data.ID, &data.Date, &data.Status); err != nil {
            log.Printf("Error al escanear datos de HallDoor: %v", err)
            return nil, err
        }
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &data, nil
}

func (pg *Postgres) UpdateStatus(id int, status int) error {
    query := "UPDATE HallDoor SET status = $1 WHERE id = $2"
    _, err := pg.conn.ExecutePreparedQuery(query, status, id)
    if err != nil {
        log.Printf("Error al actualizar el estado de HallDoor: %v", err)
        return err
    }
    return nil
}