/*******************************************************************************
* Package and license  ...
*******************************************************************************/

package users

import "context"
import "sync"

/*******************************************************************************
* Imports ...
*******************************************************************************/

/*******************************************************************************
* Type definitions ...
*******************************************************************************/

// UserAPIMock ...
type UserAPIMock struct {
	sync.Mutex
	users []User
}

/*******************************************************************************
* Init methods ...
*******************************************************************************/

/*******************************************************************************
* Private methods ...
*******************************************************************************/

/*******************************************************************************
* Public methods ...
*******************************************************************************/

// NewUserAPIMock ...
func NewUserAPIMock() *UserAPIMock {
	return &UserAPIMock{}
}

// FindAll ...
func (u *UserAPIMock) FindAll(ctx context.Context) ([]User, error) {
	u.Lock()
	defer u.Unlock()

	// return mock data
	return u.users, nil
}
