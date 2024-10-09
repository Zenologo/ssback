package user

import (
	"database/sql"
	"fmt"
	logger "ssproxy/back/internal/pkg"
	"ssproxy/back/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			fmt.Println("error: ", err)
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

/**
* if we can insert this user, return
 */
func (s *Store) IsNewUser(email string) (bool, error) {
	rows, err := s.db.Query("SELECT COUNT(*) FROM users WHERE email = ?", email)
	if err != nil {
		return false, err
	}

	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			logger.Error.Println(err)
			return false, err
		}

	}

	fmt.Println(count)

	if count == 0 {
		return true, nil
	}

	return false, fmt.Errorf("this user is already in system")
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, err
	}

	return u, nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedDate,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
