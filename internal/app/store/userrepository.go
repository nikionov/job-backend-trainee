package store

import "github.com/nikionov/job-backend-trainee/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error){
	if err := r.store.db.QueryRow("INSERT INTO users (email, encryptedpassword) VALUES ($1, $2) RETURNING id", u.Email, u.EncryptedPassword).Scan(&u.Id); err != nill {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string)(*model.User, error){
	return nil, nil
}
