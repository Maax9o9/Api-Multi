package infrestructure

import (
    "Multi/src/core"
    "Multi/src/house/domain"
    "Multi/src/house/domain/entities"
    "log"
)

type Postgres struct {
    conn *core.Conn_Postgres
}

func NewPostgres() domain.HouseRepository {
    conn := core.GetDBPool()
    if conn.Err != "" {
        log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
    }

    return &Postgres{conn: conn}
}

func (pg *Postgres) Create(house *entities.HouseProfile) error {
    query := "INSERT INTO houses (user_id, ubication_gps, house_name, image) VALUES ($1, $2, $3, $4)"
    _, err := pg.conn.ExecutePreparedQuery(query, house.UserID, house.UbicationGps, house.HouseName, house.Image)
    if err != nil {
        log.Printf("Error al insertar casa: %v", err)
        return err
    }
    return nil
}

func (pg *Postgres) GetAll() ([]entities.HouseProfile, error) {
    query := "SELECT house_id, user_id, ubication_gps, house_name, image FROM houses"
    rows := pg.conn.FetchRows(query)
    defer rows.Close()

    var houses []entities.HouseProfile

    for rows.Next() {
        var house entities.HouseProfile
        if err := rows.Scan(&house.HouseID, &house.UserID, &house.UbicationGps, &house.HouseName, &house.Image); err != nil {
            log.Printf("Error al escanear casa: %v", err)
            return nil, err
        }
        houses = append(houses, house)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return houses, nil
}

func (pg *Postgres) GetByID(id int) (*entities.HouseProfile, error) {
    query := "SELECT house_id, user_id, ubication_gps, house_name, image FROM houses WHERE house_id = $1"
    rows := pg.conn.FetchRows(query, id)
    defer rows.Close()

    var house entities.HouseProfile
    if rows.Next() {
        if err := rows.Scan(&house.HouseID, &house.UserID, &house.UbicationGps, &house.HouseName, &house.Image); err != nil {
            log.Printf("Error al escanear casa: %v", err)
            return nil, err
        }
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &house, nil
}

func (pg *Postgres) Update(house *entities.HouseProfile) error {
    query := "UPDATE houses SET user_id = $1, ubication_gps = $2, house_name = $3, image = $4 WHERE house_id = $5"
    _, err := pg.conn.ExecutePreparedQuery(query, house.UserID, house.UbicationGps, house.HouseName, house.Image, house.HouseID)
    if err != nil {
        log.Printf("Error al actualizar casa: %v", err)
        return err
    }
    return nil
}