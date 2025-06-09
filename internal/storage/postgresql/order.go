package postgresql

import (
	"restapi-sportshop/internal/storage"
	"fmt"
)

func (s *Storage) AddOrder(order *storage.Order) error {
        const op = "internal.storage.AddOrder"

        stmt, err := s.DB.Prepare(`INSERT INTO order (user_id,
		item_id,
		quantity,
		ship_date,
		status) VALUES ($1, $2, $3, $4, $5);`)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        _, err = stmt.Exec(order.UserID, order.ItemID, order.Quantity, order.ShipDate, order.Status)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        return nil
}

func (s *Storage) GetOrder(id int) (*storage.Order, error) {
        const op = "internal.storage.GetOrder"
        var res storage.Order

        stmt, err := s.DB.Prepare("SELECT * FROM order WHERE id = $1;")
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        err = stmt.QueryRow(id).Scan(&res.ID, &res.UserID, &res.ItemID, &res.Quantity, &res.ShipDate, &res.Status)
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        return &res, nil
}

func (s *Storage) UpdateOrder(id int, order *storage.Order) error {
        const op = "internal.storage.UpdateOrder"

        stmt, err := s.DB.Prepare(`UPDATE order
                SET user_id = $1,
                    item_id = $2,
                    quantity = $3,
                    ship_date = $4,
		    status = $5
                WHERE id = $6;`)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        _, err = stmt.Exec(
                order.UserID,
                order.ItemID,
                order.Quantity,
                order.ShipDate,
		order.Status,
                id)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        return nil
}

func (s *Storage) DeleteOrder(id int) error {
        const op = "internal.storage.DeleteOrder"

        stmt, err := s.DB.Prepare("DELETE FROM order WHERE id = $1;")
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        _, err = stmt.Exec(id)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        return nil
}
