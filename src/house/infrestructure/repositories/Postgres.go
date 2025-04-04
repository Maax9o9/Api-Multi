package repositories

import (
	"database/sql"
	"errors"
	"time"

	"Multi/src/core"
	"Multi/src/house/domain"
	"Multi/src/house/domain/entities"
)

type PostgresHouseRepository struct {
	db *core.Conn_Postgres
}

func NewPostgresHouseRepository(db *core.Conn_Postgres) domain.HouseRepository {
	return &PostgresHouseRepository{
		db: db,
	}
}

func (r *PostgresHouseRepository) Save(house *entities.HouseProfile) (int, error) {
	query := `
        INSERT INTO houses (owner_id, ubication_gps, house_name, image, device_code, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING house_id
    `

	if house.CreatedAt == 0 {
		house.CreatedAt = time.Now().Unix()
	}

	var houseID int
	err := r.db.FetchRow(
		query,
		house.OwnerID,
		house.UbicationGps,
		house.HouseName,
		house.Image,
		house.DeviceCode,
		house.CreatedAt,
	).Scan(&houseID)

	if err != nil {
		return 0, err
	}

	return houseID, nil
}

func (r *PostgresHouseRepository) GetByID(id int) (*entities.HouseProfile, error) {
	query := `
        SELECT house_id, owner_id, ubication_gps, house_name, image, device_code, created_at
        FROM houses
        WHERE house_id = $1
    `

	var house entities.HouseProfile
	err := r.db.FetchRow(query, id).Scan(
		&house.HouseID,
		&house.OwnerID,
		&house.UbicationGps,
		&house.HouseName,
		&house.Image,
		&house.DeviceCode,
		&house.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("casa no encontrada")
		}
		return nil, err
	}

	return &house, nil
}

func (r *PostgresHouseRepository) GetByDeviceCode(deviceCode string) (*entities.HouseProfile, error) {
	query := `
        SELECT house_id, owner_id, ubication_gps, house_name, image, device_code, created_at
        FROM houses
        WHERE device_code = $1
    `

	var house entities.HouseProfile
	err := r.db.FetchRow(query, deviceCode).Scan(
		&house.HouseID,
		&house.OwnerID,
		&house.UbicationGps,
		&house.HouseName,
		&house.Image,
		&house.DeviceCode,
		&house.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("casa no encontrada")
		}
		return nil, err
	}

	return &house, nil
}

func (r *PostgresHouseRepository) GetHousesByUserID(userID int) ([]*entities.HouseProfile, error) {
	query := `
        SELECT h.house_id, h.owner_id, h.ubication_gps, h.house_name, h.image, h.device_code, h.created_at
        FROM houses h
        LEFT JOIN house_members m ON h.house_id = m.house_id
        WHERE h.owner_id = $1 OR m.user_id = $1
    `

	rows := r.db.FetchRows(query, userID)
	defer rows.Close()

	var houses []*entities.HouseProfile
	for rows.Next() {
		var house entities.HouseProfile
		if err := rows.Scan(
			&house.HouseID,
			&house.OwnerID,
			&house.UbicationGps,
			&house.HouseName,
			&house.Image,
			&house.DeviceCode,
			&house.CreatedAt,
		); err != nil {
			return nil, err
		}

		houses = append(houses, &house)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return houses, nil
}

func (r *PostgresHouseRepository) Update(house *entities.HouseProfile) error {
	query := `
        UPDATE houses
        SET ubication_gps = $1, house_name = $2, image = $3, device_code = $4
        WHERE house_id = $5
    `

	result, err := r.db.ExecutePreparedQuery(
		query,
		house.UbicationGps,
		house.HouseName,
		house.Image,
		house.DeviceCode,
		house.HouseID,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("casa no encontrada")
	}

	return nil
}

func (r *PostgresHouseRepository) Delete(id int) error {
	query := "DELETE FROM houses WHERE house_id = $1"

	result, err := r.db.ExecutePreparedQuery(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("casa no encontrada")
	}

	return nil
}
