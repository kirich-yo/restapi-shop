package postgresql

import (
	"restapi-sportshop/internal/storage"
	"fmt"
)

func (s *Storage) GetItems() ([]storage.Item, error) {
        const op = "internal.storage.GetItems"

        stmt, err := s.DB.Prepare("SELECT * FROM item;")
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        rows, err := stmt.Query()
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

func (s *Storage) AddItem(item *storage.Item) error {
        const op = "internal.storage.AddItem"

        stmt, err := s.DB.Prepare("INSERT INTO item (name, price, sale_price, photo_url) VALUES ($1, $2, $3, $4);")
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        _, err = stmt.Exec(item.Name, item.Price, item.SalePrice, item.PhotoURL)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        return nil
}

func (s *Storage) GetItem(id int) (*storage.Item, error) {
        const op = "internal.storage.GetItem"
        var res storage.Item

        stmt, err := s.DB.Prepare("SELECT * FROM item WHERE id = $1;")
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        err = stmt.QueryRow(id).Scan(&res.ID, &res.Name, &res.Price, &res.SalePrice, &res.PhotoURL)
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        return &res, nil
}

func (s *Storage) UpdateItem(id int, item *storage.Item) error {
        const op = "internal.storage.UpdateItem"

        stmt, err := s.DB.Prepare(`UPDATE item
                SET name = $1,
                    price = $2,
                    sale_price = $3,
                    photo_url = $4
                WHERE id = $5;`)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        _, err = stmt.Exec(
                item.Name,
                item.Price,
                item.SalePrice,
                item.PhotoURL,
                id)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        return nil
}

func (s *Storage) DeleteItem(id int) error {
        const op = "internal.storage.DeleteItem"

        stmt, err := s.DB.Prepare("DELETE FROM item WHERE id = $1;")
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        _, err = stmt.Exec(id)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        return nil
}

func (s *Storage) GetItemReviews(id int) ([]storage.Review, error) {
        const op = "internal.storage.GetItemReviews"

        stmt, err := s.DB.Prepare("SELECT * FROM review WHERE item_id = $1;")
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }

        rows, err := stmt.Query(id)
        if err != nil {
                return nil, fmt.Errorf("%s: %w", op, err)
        }
        defer rows.Close()

        reviews := make([]storage.Review, 0)

        for rows.Next() {
                var review storage.Review

                err = rows.Scan(
                        &review.ID,
                        &review.Rating,
                        &review.Advantages,
                        &review.Disadvantages,
                        &review.Description,
                        &review.UserID,
                        &review.ItemID)
                if err != nil {
                        return nil, fmt.Errorf("%s: %w", op, err)
                }

                reviews = append(reviews, review)
        }

        return reviews, nil
}

func (s *Storage) AddItemReview(review *storage.Review) error {
        const op = "internal.storage.AddItemReview"

        stmt, err := s.DB.Prepare(`INSERT INTO review (
                rating,
                advantages,
                disadvantages,
                description,
                user_id,
                item_id
        ) VALUES ($1, $2, $3, $4, $5, $6);`)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        _, err = stmt.Exec(
                review.Rating,
                review.Advantages,
                review.Disadvantages,
                review.Description,
                review.UserID,
                review.ItemID)
        if err != nil {
                return fmt.Errorf("%s: %w", op, err)
        }

        return nil
}
