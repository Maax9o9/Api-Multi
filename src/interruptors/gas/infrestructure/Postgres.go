package infrestructure

import (
    "Multi/src/core"
    "Multi/src/interruptors/gas/domain"
    "Multi/src/interruptors/gas/domain/entities"
    "log"
)

type Postgres struct {
    conn *core.Conn_Postgres
}

func NewPostgres() domain.GasRepository {
    conn := core.GetDBPool()
    if conn.Err != "" {
        log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
    }

    return &Postgres{conn: conn}
}

func (pg *Postgres) Create(data *entities.GasSensor) error {
    query := "INSERT INTO GasSensor (createdAt, status, gasLevel) VALUES ($1, $2, $3)"
    _, err := pg.conn.ExecutePreparedQuery(query, data.CreatedAt, data.Status, data.GasLevel)
    if err != nil {
        log.Printf("Error al insertar datos de gas: %v", err)
        return err
    }
    return nil
}

func (pg *Postgres) GetAll() ([]entities.GasSensor, error) {
    query := "SELECT id, createdAt, status, gasLevel FROM GasSensor"
    rows := pg.conn.FetchRows(query)
    defer rows.Close()

    var gasData []entities.GasSensor

    for rows.Next() {
        var data entities.GasSensor
        if err := rows.Scan(&data.ID, &data.CreatedAt, &data.Status, &data.GasLevel); err != nil {
            log.Printf("Error al escanear datos de gas: %v", err)
            return nil, err
        }
        gasData = append(gasData, data)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return gasData, nil
}

func (pg *Postgres) GetByID(id int) (*entities.GasSensor, error) {
    query := "SELECT id, createdAt, status, gasLevel FROM GasSensor WHERE id = $1"
    rows := pg.conn.FetchRows(query, id)
    defer rows.Close()

    var data entities.GasSensor
    if rows.Next() {
        if err := rows.Scan(&data.ID, &data.CreatedAt, &data.Status, &data.GasLevel); err != nil {
            log.Printf("Error al escanear datos de gas: %v", err)
            return nil, err
        }
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &data, nil
}