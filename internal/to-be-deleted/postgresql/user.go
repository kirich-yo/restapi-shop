package postgresql

import (
	"restapi-sportshop/internal/storage"
	"fmt"
)

func (s *Storage) AddUser(user *storage.User) error {
        const op = "internal.storage.AddUser"
        stmt, err := s.DB.Prepare(`INSERT INTO public.user (username,
                first_name,
                last_name,
                date_of_birth,
                photo_url,
		role_id
                password) VALUES ($1, $2, $3, $4, $5, $6, $7)`)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        _, err = stmt.Exec(user.Username,
                user.FirstName,
                user.LastName,
                user.DateOfBirth,
                user.PhotoURL,
                user.Password)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        return nil
}

func (s *Storage) GetUser(id int) (*storage.User, error) {
        const op = "internal.storage.GetUser"
        var res storage.User

        stmt, err := s.DB.Prepare("SELECT * FROM public.user WHERE id = $1;")
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        err = stmt.QueryRow(id).Scan(&res.ID,
		&res.Username,
		&res.FirstName,
		&res.LastName,
		&res.DateOfBirth,
		&res.PhotoURL,
		&res.RoleID,
		&res.Password)
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        return &res, nil
}

func (s *Storage) UpdateUser(id int, user *storage.User) error {
        const op = "internal.storage.UpdateUser"

        stmt, err := s.DB.Prepare(`UPDATE public.user
                SET username = $1,
                    first_name = $2,
                    last_name = $3,
                    date_of_birth = $4,
                    photo_url = $5,
		    role_id = $6,
		    password = $7
                WHERE id = $8;`)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        _, err = stmt.Exec(
                user.Username,
                user.FirstName,
                user.LastName,
                user.DateOfBirth,
                user.PhotoURL,
		user.RoleID,
		user.Password,
                id)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        return nil
}

func (s *Storage) DeleteUser(id int) error {
        const op = "internal.storage.DeleteUser"

        stmt, err := s.DB.Prepare("DELETE FROM public.user WHERE id = $1;")
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        _, err = stmt.Exec(id)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        return nil
}
