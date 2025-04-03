package infrestructure

import (
    "Multi/src/core"
    "Multi/src/interruptors/window/domain"
    "Multi/src/interruptors/window/domain/entities"
    "log"
)

type Postgres struct {
    conn *core.Conn_Postgres
}

func NewPostgres() domain.WindowRepository {
    conn := core.GetDBPool()
    if conn.Err != "" {
        log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
    }

    return &Postgres{conn: conn}
}

func (pg *Postgres) Create(data *entities.WindowSensor) error {
    query := "INSERT INTO WindowSensor (createdAt, status) VALUES ($1, $2)"
    _, err := pg.conn.ExecutePreparedQuery(query, data.CreatedAt, data.Status)
    if err != nil {
        log.Printf("Error al insertar datos de ventana: %v", err)
        return err
    }
    return nil
}

func (pg *Postgres) GetAll() ([]entities.WindowSensor, error) {
    query := "SELECT id, createdAt, status FROM WindowSensor"
    rows := pg.conn.FetchRows(query)
    defer rows.Close()

    var windowData []entities.WindowSensor

    for rows.Next() {
        var data entities.WindowSensor
        if err := rows.Scan(&data.ID, &data.CreatedAt, &data.Status); err != nil {
            log.Printf("Error al escanear datos de ventana: %v", err)
            return nil, err
        }
        windowData = append(windowData, data)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return windowData, nil
}

func (pg *Postgres) GetByID(id int) (*entities.WindowSensor, error) {
    query := "SELECT id, createdAt, status FROM WindowSensor WHERE id = $1"
    rows := pg.conn.FetchRows(query, id)
    defer rows.Close()

    var data entities.WindowSensor
    if rows.Next() {
        if err := rows.Scan(&data.ID, &data.CreatedAt, &data.Status); err != nil {
            log.Printf("Error al escanear datos de ventana: %v", err)
            return nil, err
        }
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &data, nil
}