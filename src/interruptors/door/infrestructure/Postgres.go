package infrestructure

import (
    "Multi/src/core"
    "Multi/src/interruptors/door/domain"
    "Multi/src/interruptors/door/domain/entities"
    "log"
)

type Postgres struct {
    conn *core.Conn_Postgres
}

func NewPostgres() domain.DoorRepository {
    conn := core.GetDBPool()
    if conn.Err != "" {
        log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
    }

    return &Postgres{conn: conn}
}

func (pg *Postgres) Create(data *entities.DoorData) error {
    query := "INSERT INTO DoorSensor (house_id, createdAt, status) VALUES ($1, $2, $3)"
    _, err := pg.conn.ExecutePreparedQuery(query, data.HouseID, data.CreatedAt, data.Status)
    if err != nil {
        log.Printf("Error al insertar datos de Door: %v", err)
        return err
    }
    return nil
}

func (pg *Postgres) GetAll() ([]entities.DoorData, error) {
    query := "SELECT id, house_id, createdAt, status FROM DoorSensor"
    rows := pg.conn.FetchRows(query)
    defer rows.Close()

    var doorData []entities.DoorData

    for rows.Next() {
        var data entities.DoorData
        if err := rows.Scan(&data.ID, &data.HouseID, &data.CreatedAt, &data.Status); err != nil {
            log.Printf("Error al escanear datos de Door: %v", err)
            return nil, err
        }
        doorData = append(doorData, data)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return doorData, nil
}

func (pg *Postgres) GetByID(id int) (*entities.DoorData, error) {
    query := "SELECT id, house_id, createdAt, status FROM DoorSensor WHERE id = $1"
    rows := pg.conn.FetchRows(query, id)
    defer rows.Close()

    var data entities.DoorData
    if rows.Next() {
        if err := rows.Scan(&data.ID, &data.HouseID, &data.CreatedAt, &data.Status); err != nil {
            log.Printf("Error al escanear datos de Door: %v", err)
            return nil, err
        }
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &data, nil
}