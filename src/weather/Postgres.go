package infrestructure

import (
    "Multi/src/core"
    "Multi/src/weather/domain"
    "Multi/src/weather/domain/entities"
    "log"
)

type Postgres struct {
    conn *core.Conn_Postgres
}

func NewPostgres() domain.WeatherRepository {
    conn := core.GetDBPool()
    if conn.Err != "" {
        log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
    }

    return &Postgres{conn: conn}
}

func (pg *Postgres) Create(data *entities.SensorDataWeather) error {
    query := "INSERT INTO weather_data (house_id, date, heat, damp) VALUES ($1, $2, $3, $4)"
    _, err := pg.conn.ExecutePreparedQuery(query, data.HouseID, data.Date, data.Heat, data.Damp)
    if err != nil {
        log.Printf("Error al insertar datos meteorológicos: %v", err)
        return err
    }
    return nil
}

func (pg *Postgres) GetAll() ([]entities.SensorDataWeather, error) {
    query := "SELECT weather_id, house_id, date, heat, damp FROM weather_data"
    rows := pg.conn.FetchRows(query)
    defer rows.Close()

    var weatherData []entities.SensorDataWeather

    for rows.Next() {
        var data entities.SensorDataWeather
        if err := rows.Scan(&data.WeatherID, &data.HouseID, &data.Date, &data.Heat, &data.Damp); err != nil {
            log.Printf("Error al escanear datos meteorológicos: %v", err)
            return nil, err
        }
        weatherData = append(weatherData, data)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return weatherData, nil
}

func (pg *Postgres) GetByID(id int) (*entities.SensorDataWeather, error) {
    query := "SELECT weather_id, house_id, date, heat, damp FROM weather_data WHERE weather_id = $1"
    rows := pg.conn.FetchRows(query, id)
    defer rows.Close()

    var data entities.SensorDataWeather
    if rows.Next() {
        if err := rows.Scan(&data.WeatherID, &data.HouseID, &data.Date, &data.Heat, &data.Damp); err != nil {
            log.Printf("Error al escanear datos meteorológicos: %v", err)
            return nil, err
        }
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &data, nil
}