package infrestructure

import (
    "Multi/src/core"
    "Multi/src/notification/domain"
    "Multi/src/notification/domain/entities"
    "log"
)

type Postgres struct {
    conn *core.Conn_Postgres
}

func NewPostgres() domain.NotificationRepository {
    conn := core.GetDBPool()
    if conn.Err != "" {
        log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
    }

    return &Postgres{conn: conn}
}

func (pg *Postgres) Create(notification *entities.Notification) error {
    query := "INSERT INTO notifications (house_id, date, message, type_notification) VALUES ($1, $2, $3, $4)"
    _, err := pg.conn.ExecutePreparedQuery(query, notification.HouseID, notification.Date, notification.Message, notification.TypeNotification)
    if err != nil {
        log.Printf("Error al insertar notificación: %v", err)
        return err
    }
    return nil
}

func (pg *Postgres) GetAll() ([]entities.Notification, error) {
    query := "SELECT id, house_id, date, message, type_notification FROM notifications"
    rows := pg.conn.FetchRows(query)
    defer rows.Close()

    var notifications []entities.Notification

    for rows.Next() {
        var notification entities.Notification
        if err := rows.Scan(&notification.ID, &notification.HouseID, &notification.Date, &notification.Message, &notification.TypeNotification); err != nil {
            log.Printf("Error al escanear notificación: %v", err)
            return nil, err
        }
        notifications = append(notifications, notification)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return notifications, nil
}

func (pg *Postgres) GetByID(id int) (*entities.Notification, error) {
    query := "SELECT id, house_id, date, message, type_notification FROM notifications WHERE id = $1"
    rows := pg.conn.FetchRows(query, id)
    defer rows.Close()

    var notification entities.Notification
    if rows.Next() {
        if err := rows.Scan(&notification.ID, &notification.HouseID, &notification.Date, &notification.Message, &notification.TypeNotification); err != nil {
            log.Printf("Error al escanear notificación: %v", err)
            return nil, err
        }
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &notification, nil
}