package infrestructure

import (
    "Multi/src/core"
    "Multi/src/interruptors/lightOutside/domain"
    "Multi/src/interruptors/lightOutside/domain/entities"
    "log"
)

type Postgres struct {
    conn *core.Conn_Postgres
}

func NewPostgres() domain.LightOutsideRepository {
    conn := core.GetDBPool()
    if conn.Err != "" {
        log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
    }

    return &Postgres{conn: conn}
}

func (pg *Postgres) Create(data *entities.LightOutsideData) error {
    query := "INSERT INTO LedControl (createdAt, status) VALUES ($1, $2)"
    _, err := pg.conn.ExecutePreparedQuery(query, data.CreatedAt, data.Status)
    if err != nil {
        log.Printf("Error al insertar datos de Light: %v", err)
        return err
    }
    return nil
}

func (pg *Postgres) GetAll() ([]entities.LightOutsideData, error) {
    query := "SELECT id, createdAt, status FROM LedControl"
    rows := pg.conn.FetchRows(query)
    defer rows.Close()

    var lightData []entities.LightOutsideData

    for rows.Next() {
        var data entities.LightOutsideData
        if err := rows.Scan(&data.ID, &data.CreatedAt, &data.Status); err != nil {
            log.Printf("Error al escanear datos de Light: %v", err)
            return nil, err
        }
        lightData = append(lightData, data)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return lightData, nil
}

func (pg *Postgres) GetByID(id int) (*entities.LightOutsideData, error) {
    query := "SELECT id, createdAt, status FROM LedControl WHERE id = $1"
    rows := pg.conn.FetchRows(query, id)
    defer rows.Close()

    var data entities.LightOutsideData
    if rows.Next() {
        if err := rows.Scan(&data.ID, &data.CreatedAt, &data.Status); err != nil {
            log.Printf("Error al escanear datos de Light: %v", err)
            return nil, err
        }
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &data, nil
}