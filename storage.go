package xstorage

import (
	"github.com/google/uuid"
)

type (
	Item struct {
		// ID the item id
		ID uuid.UUID
		// AccountID the account owning the item
		AccountID uuid.UUID
		// ObjectKey the key is the file name
		ObjectKey string
		// ContentType image etc.
		ContentType string
		// Size the size of the file
		Size int64
		// CreatedAt the timestamp the item was created
		CreatedAt time.Time
	}
)
