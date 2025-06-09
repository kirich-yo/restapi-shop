package postgresql

import (
	"restapi-sportshop/internal/storage"
	"fmt"
)

func (s *Storage) GetCategories() ([]storage.Category, error) {
        const op = "internal.storage.GetCategories"

        stmt, err := s.DB.Prepare("SELECT * FROM category;")
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        rows, err := stmt.Query()
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }
        defer rows.Close()

        categories := make([]storage.Category, 0)

        for rows.Next() {
                var category storage.Category

                err = rows.Scan(&category.ID, &category.Name)
                if err != nil {
                        return nil, fmt.Errorf("%s: %w", op, err)
                }

                categories = append(categories, category)
        }

        return categories, nil
}

func (s *Storage) GetCategory(id int) (*storage.Category, error) {
        const op = "internal.storage.GetCategory"
        var res storage.Category

        stmt, err := s.DB.Prepare("SELECT * FROM category WHERE id = $1;")
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        err = stmt.QueryRow(id).Scan(&res.ID, &res.Name)
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        return &res, nil
}
