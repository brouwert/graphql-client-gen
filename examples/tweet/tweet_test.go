package tweet

import (
	"testing"
)

func TestTweetQuery(t *testing.T) {
	q := Tweet{
		ID:   Get.ID,
		Body: &Get.String,
		Author: &User{
			Username: &Get.String,
		},
		Stats: &Stat{
			Views: &Get.Int32,
			Likes: &Get.Int32,
		},
		Responders: &TweetRespondersFunction{
			Limit: 10,
			User: User{
				ID: Get.ID,
			},
		},
	}
	t.Log("\n", BuildQuery(q))
}

var Get = struct {
	String string
	ID     interface{}
	Int32  int32
}{
	"",
	0,
	0,
}
