package books

import (
	"github.com/upper/db/v4"
)

func (b *BookModule) table(name string) db.Collection {

	return b.sess.Collection(name)
}
