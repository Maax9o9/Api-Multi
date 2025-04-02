package infrestructure

import (
    "Multi/src/core"
    "Multi/src/interruptors/movement/domain"
    "Multi/src/interruptors/movement/domain/entities"
    "log"
)

type Postgres struct {
    conn *core.Conn_Postgres
}

func NewPostgres() domain.MovementRepository {
    conn := core.GetDBPool()
    if conn.Err != "" {
        log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
    }

    return &Postgres{conn: conn}
}

func (pg *Postgres) Create(data *entities.MotionSensor) error {
    query := "INSERT INTO MotionSensor (house_id, createdAt, status) VALUES ($1, $2, $3)"
    _, err := pg.conn.ExecutePreparedQuery(query, data.HouseID, data.CreatedAt, data.Status)
    if err != nil {
        log.Printf("Error al insertar datos de movimiento: %v", err)
        return err
    }
    return nil
}

func (pg *Postgres) GetAll() ([]entities.MotionSensor, error) {
    query := "SELECT id, house_id, createdAt, status FROM MotionSensor"
    rows := pg.conn.FetchRows(query)
    defer rows.Close()

    var motionData []entities.MotionSensor

    for rows.Next() {
        var data entities.MotionSensor
        if err := rows.Scan(&data.ID, &data.HouseID, &data.CreatedAt, &data.Status); err != nil {
            log.Printf("Error al escanear datos de movimiento: %v", err)
            return nil, err
        }
        motionData = append(motionData, data)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return motionData, nil
}

func (pg *Postgres) GetByID(id int) (*entities.MotionSensor, error) {
    query := "SELECT id, house_id, createdAt, status FROM MotionSensor WHERE id = $1"
    rows := pg.conn.FetchRows(query, id)
    defer rows.Close()

    var data entities.MotionSensor
    if rows.Next() {
        if err := rows.Scan(&data.ID, &data.HouseID, &data.CreatedAt, &data.Status); err != nil {
            log.Printf("Error al escanear datos de movimiento: %v", err)
            return nil, err
        }
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &data, nil
}