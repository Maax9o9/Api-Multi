package repositories

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"Multi/src/core"
	"Multi/src/house/domain"
	"Multi/src/house/domain/entities"
)

type PostgresHouseMemberRepository struct {
	db *core.Conn_Postgres
}

func NewPostgresHouseMemberRepository(db *core.Conn_Postgres) domain.HouseMemberRepository {
	return &PostgresHouseMemberRepository{
		db: db,
	}
}

// En el método SaveMember:
func (r *PostgresHouseMemberRepository) SaveMember(member *entities.HouseMember) error {
	existingMember, err := r.GetByUserAndHouse(member.UserID, member.HouseID)
	if err == nil && existingMember != nil {
		return nil
	}

	// Si el error es diferente de "no encontrado", entonces es un error real
	if err != nil && err != sql.ErrNoRows && !strings.Contains(err.Error(), "no encontrado") {
		return err
	}

	query := `
        INSERT INTO house_members (house_id, user_id, role, joined_at)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (house_id, user_id) DO NOTHING
    `

	if member.JoinedAt == 0 {
		member.JoinedAt = time.Now().Unix()
	}

	_, err = r.db.ExecutePreparedQuery(
		query,
		member.HouseID,
		member.UserID,
		member.Role,
		member.JoinedAt,
	)

	return err
}

func (r *PostgresHouseMemberRepository) GetByUserID(userID int) ([]*entities.HouseMember, error) {
	query := `
        SELECT house_id, user_id, role, joined_at
        FROM house_members
        WHERE user_id = $1
    `

	rows := r.db.FetchRows(query, userID)
	defer rows.Close()

	var members []*entities.HouseMember
	for rows.Next() {
		var member entities.HouseMember
		if err := rows.Scan(
			&member.HouseID,
			&member.UserID,
			&member.Role,
			&member.JoinedAt,
		); err != nil {
			return nil, err
		}

		members = append(members, &member)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}

func (r *PostgresHouseMemberRepository) GetByUserAndHouse(userID, houseID int) (*entities.HouseMember, error) {
	query := `
        SELECT house_id, user_id, role, joined_at
        FROM house_members
        WHERE user_id = $1 AND house_id = $2
    `

	var member entities.HouseMember
	err := r.db.FetchRow(query, userID, houseID).Scan(
		&member.HouseID,
		&member.UserID,
		&member.Role,
		&member.JoinedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("miembro no encontrado")
		}
		return nil, err
	}

	return &member, nil
}

func (r *PostgresHouseMemberRepository) GetHouseMembers(houseID int) ([]*entities.HouseMember, error) {
	query := `
        SELECT house_id, user_id, role, joined_at
        FROM house_members
        WHERE house_id = $1
    `

	rows := r.db.FetchRows(query, houseID)
	defer rows.Close()

	var members []*entities.HouseMember
	for rows.Next() {
		var member entities.HouseMember
		if err := rows.Scan(
			&member.HouseID,
			&member.UserID,
			&member.Role,
			&member.JoinedAt,
		); err != nil {
			return nil, err
		}

		members = append(members, &member)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}

func (r *PostgresHouseMemberRepository) DeleteMember(userID, houseID int) error {
	query := `
        DELETE FROM house_members
        WHERE user_id = $1 AND house_id = $2
    `

	result, err := r.db.ExecutePreparedQuery(query, userID, houseID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("miembro no encontrado")
	}

	return nil
}
