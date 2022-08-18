package main

import "database/sql"

var _ UserCRUD = (*UserStore)(nil)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) Create(u *User) (int, error) {
	stmt := "INSERT INTO users VALUES(NULL, ?, ?, ?, ?, ?)"
	res, err := us.db.Exec(stmt, u.Name, u.Email, u.Password, u.Salt, u.Admin)
	if err != nil {
		return 0, nil
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return 0, err
	}

	return int(id), nil
}

func (us *UserStore) Get(id int) (*User, error) {
	getQuery := "SELECT * FROM users WHERE id=?"

	row := us.db.QueryRow(getQuery, id)

	u := &User{}
	var err error

	if err = row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Salt, &u.Admin); err == sql.ErrNoRows {
		return u, sql.ErrNoRows
	}

	return u, err
}

func (us *UserStore) List(offset int) ([]*User, error) {
	return nil, nil
}

func (us *UserStore) Delete(id string) error {
	return nil
}

func (us *UserStore) Update(u *User) (*User, error) {
	return nil, nil
}
