package postgresql

import (
	"restapi-sportshop/internal/storage"
	"fmt"
)

func (s *Storage) GetRoles() ([]storage.Role, error) {
        const op = "internal.storage.GetRoles"

        stmt, err := s.DB.Prepare("SELECT * FROM role;")
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        rows, err := stmt.Query()
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }
        defer rows.Close()

        roles := make([]storage.Role, 0)

        for rows.Next() {
                var role storage.Role

                err = rows.Scan(&role.ID, &role.Name)
                if err != nil {
                        return nil, fmt.Errorf("%s: %w", op, err)
                }

                roles = append(roles, role)
        }

        return roles, nil
}

func (s *Storage) GetRole(id int) (*storage.Role, error) {
        const op = "internal.storage.GetRole"
        var res storage.Role

        stmt, err := s.DB.Prepare("SELECT * FROM role WHERE id = $1;")
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        err = stmt.QueryRow(id).Scan(&res.ID, &res.Name)
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        return &res, nil
}
