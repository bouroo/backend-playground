package repositories

import (
	"database/sql"
	"go-std/internal/models"

	"github.com/google/uuid"
)

type cartRepositoryImpl struct {
	db *sql.DB
}

type CartRepository interface {
	CreateCart(cart *models.Cart) error
	GetCart(id uuid.UUID) (*models.Cart, error)
	UpdateCart(id uuid.UUID, cart *models.Cart) error
	RemoveCart(id uuid.UUID) error
}

func NewCartRepository(db *sql.DB) CartRepository {
	return &cartRepositoryImpl{db: db}
}

func (r *cartRepositoryImpl) CreateCart(cart *models.Cart) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if cart.ID == uuid.Nil {
		cart.ID, err = uuid.NewV7()
		if err != nil {
			return err
		}
	}

	_, err = tx.Exec("INSERT INTO carts (id, user_id, items) VALUES (?, ?, ?)", cart.ID, cart.UserID, cart.Items)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *cartRepositoryImpl) GetCart(id uuid.UUID) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.QueryRow("SELECT id, user_id, items FROM carts WHERE id = ?", id).Scan(&cart.ID, &cart.UserID, &cart.Items)
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *cartRepositoryImpl) UpdateCart(id uuid.UUID, cart *models.Cart) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE carts SET user_id = ?, items = ? WHERE id = ?", cart.UserID, cart.Items, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *cartRepositoryImpl) RemoveCart(id uuid.UUID) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM carts WHERE id = ?", id)
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM cart_items WHERE cart_id = ?", id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
