package infrestructure

import (
    "Multi/src/core"
    "Multi/src/incidencies/domain"
    "Multi/src/incidencies/domain/entities"
    "log"
)

type Postgres struct {
    conn *core.Conn_Postgres
}

func NewPostgres() domain.IncidenciesRepository {
    conn := core.GetDBPool()
    if conn.Err != "" {
        log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
    }

    return &Postgres{conn: conn}
}

func (pg *Postgres) Increment(typeNotification string) (*entities.Incidency, error) {
    query := `
        UPDATE incidencies
        SET incidencies_amount = incidencies_amount + 1
        WHERE type_notification = $1
        RETURNING type_notification, incidencies_amount
    `
    row := pg.conn.FetchRows(query, typeNotification)

    var incidency entities.Incidency
    if err := row.Scan(&incidency.TypeNotification, &incidency.IncidenciesAmount); err != nil {
        log.Printf("Error al incrementar la incidencia: %v", err)
        return nil, err
    }

    return &incidency, nil
}

func (pg *Postgres) GetAll() ([]entities.Incidency, error) {
    query := "SELECT type_notification, incidencies_amount FROM incidencies"
    rows := pg.conn.FetchRows(query)
    defer rows.Close()

    var incidencies []entities.Incidency

    for rows.Next() {
        var incidency entities.Incidency
        if err := rows.Scan(&incidency.TypeNotification, &incidency.IncidenciesAmount); err != nil {
            log.Printf("Error al escanear incidencia: %v", err)
            return nil, err
        }
        incidencies = append(incidencies, incidency)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return incidencies, nil
}

func (pg *Postgres) GetByType(typeNotification string) (*entities.Incidency, error) {
    query := "SELECT type_notification, incidencies_amount FROM incidencies WHERE type_notification = $1"
    row := pg.conn.FetchRows(query, typeNotification)

    var incidency entities.Incidency
    if err := row.Scan(&incidency.TypeNotification, &incidency.IncidenciesAmount); err != nil {
        log.Printf("Error al obtener la incidencia: %v", err)
        return nil, err
    }

    return &incidency, nil
}