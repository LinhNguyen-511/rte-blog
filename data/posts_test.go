package data

import (
	"rte-blog/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutTitle(t *testing.T) {
	t.Run("update the title in the database with the new title", func(t *testing.T) {
		postModel := PostModel{db}
		post := types.Post{Id: 1, Title: "sample"}

		got, _ := postModel.PutTitle(post)
		assert.Equal(t, post, got)
	})
}
