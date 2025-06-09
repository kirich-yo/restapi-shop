package postgresql

import (
	"restapi-sportshop/internal/storage"
	"fmt"
)

func (s *Storage) GetWarehouses() ([]storage.Warehouse, error) {
        const op = "internal.storage.GetWarehouses"

        stmt, err := s.DB.Prepare("SELECT * FROM warehouse;")
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        rows, err := stmt.Query()
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }
        defer rows.Close()

        warehouses := make([]storage.Warehouse, 0)

        for rows.Next() {
                var warehouse storage.Warehouse

                err = rows.Scan(&warehouse.ID, &warehouse.Address)
                if err != nil {
                        return nil, fmt.Errorf("%s: %w", op, err)
                }

                warehouses = append(warehouses, warehouse)
        }

        return warehouses, nil
}

func (s *Storage) GetWarehouse(id int) (*storage.Warehouse, error) {
        const op = "internal.storage.GetWarehouse"
        var res storage.Warehouse

        stmt, err := s.DB.Prepare("SELECT * FROM warehouse WHERE id = $1;")
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        err = stmt.QueryRow(id).Scan(&res.ID, &res.Address)
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        return &res, nil
}

func (s *Storage) GetWarehouseItems(id int) ([]storage.Item, error) {
        const op = "internal.storage.GetWarehouseItems"

        stmt, err := s.DB.Prepare(`SELECT item.id, name, price, sale_price, photo_url
                FROM warehouse_item JOIN item
                ON warehouse_item.item_id = item.id
                WHERE warehouse_id = $1
                ORDER BY item.id;`)
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        rows, err := stmt.Query(id)
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }
        defer rows.Close()

        items := make([]storage.Item, 0)

        for rows.Next() {
                var item storage.Item

                err = rows.Scan(&item.ID, &item.Name, &item.Price, &item.SalePrice, &item.PhotoURL)
                if err != nil {
                        return nil, fmt.Errorf("%s: %w", op, err)
                }

                items = append(items, item)
        }

        return items, nil
}
