/*******************************************************************************
* Package and license  ...
*******************************************************************************/

package books

import "context"
import "sync"

/*******************************************************************************
* Imports ...
*******************************************************************************/

/*******************************************************************************
* Type definitions ...
*******************************************************************************/

// BookAPIMock ...
type BookAPIMock struct {
	sync.Mutex
	books []Book
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

// NewBookAPIMock ...
func NewBookAPIMock() *BookAPIMock {
	return &BookAPIMock{}
}

// FindAll ...
func (u *BookAPIMock) FindAll(ctx context.Context) ([]Book, error) {
	u.Lock()
	defer u.Unlock()

	// return mock data
	return u.books, nil
}