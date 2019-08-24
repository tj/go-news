package news_test

import (
	"testing"

	"github.com/tj/assert"

	"github.com/tj/go-news"
)

// newStorage helper.
func newStorage(t testing.TB) *news.Store {
	return news.New("news_test")
}

// Test adding subscribers.
func TestStore_AddSubscriber(t *testing.T) {
	db := newStorage(t)

	err := db.AddSubscriber("general", "manny@apex.sh")
	assert.NoError(t, err, "adding subscriber")

	err = db.AddSubscriber("general", "tobi@apex.sh")
	assert.NoError(t, err, "adding subscriber")

	err = db.AddSubscriber("general", "loki@apex.sh")
	assert.NoError(t, err, "adding subscriber")

	err = db.AddSubscriber("general", "jane@apex.sh")
	assert.NoError(t, err, "adding subscriber")

	err = db.AddSubscriber("blog", "tj@apex.sh")
	assert.NoError(t, err, "adding subscriber")

	err = db.AddSubscriber("blog", "jane@apex.sh")
	assert.NoError(t, err, "adding subscriber")
}

// Test removing subscribers.
func TestStore_RemoveSubscriber(t *testing.T) {
	db := newStorage(t)

	err := db.AddSubscriber("up", "luna@apex.sh")
	assert.NoError(t, err, "adding subscriber")

	emails, err := db.GetSubscribers("up")
	assert.NoError(t, err, "listing subscribers")
	assert.Len(t, emails, 1, "emails")

	err = db.RemoveSubscriber("up", "luna@apex.sh")
	assert.NoError(t, err, "removing subscriber")

	emails, err = db.GetSubscribers("up")
	assert.NoError(t, err, "listing subscribers")
	assert.Len(t, emails, 0, "emails")

	err = db.RemoveSubscriber("up", "luna@apex.sh")
	assert.NoError(t, err, "removing subscriber")
}

// Test listing subscribers.
func TestStore_GetSubscribers(t *testing.T) {
	db := newStorage(t)

	emails, err := db.GetSubscribers("general")
	assert.NoError(t, err, "listing subscribers")
	assert.Len(t, emails, 4, "emails")

	emails, err = db.GetSubscribers("blog")
	assert.NoError(t, err, "listing subscribers")
	assert.Len(t, emails, 2, "emails")
}
