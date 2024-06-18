package server

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"rte-blog/types"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type StubPostModel struct {
	Store *sql.DB
}

func (model *StubPostModel) Create(title string) (int, error) {
	return 1, nil
}

func (model *StubPostModel) GetById(id int) (post *types.Post, err error) {
	return nil, nil
}

func (model *StubPostModel) PutTitle(post types.Post) (types.Post, error) {
	return post, nil
}

func (model *StubPostModel) CreatePostContent(id int) (*types.Content, error) {
	return &types.Content{Id: 1}, nil
}

func TestHandleGetPost(t *testing.T) {
	t.Run("returns a post with title, meta-data and content", func(t *testing.T) {
		server := createServer(t)

		context, response := makeRequest(t, http.MethodGet, "/posts/1", nil)
		context.SetParamNames("id")
		context.SetParamValues("1")

		if assert.NoError(t, server.handleGetPost(context)) {
			assert.Equal(t, http.StatusOK, response.Code)

			doc, err := goquery.NewDocumentFromReader(strings.NewReader(response.Body.String()))
			if err != nil {
				log.Fatal(err)
			}

			title := doc.Find("h1").First().Text()

			assert.Equal(t, "Sample post", title)
		}
	})
}

func TestHandleContentCreate(t *testing.T) {
	t.Run("returns new main element when creating a new content block", func(t *testing.T) {
		server := createServer(t)

		context, response := makeRequest(t, http.MethodGet, "/posts/1/paragraphs", nil)
		context.SetParamNames("id")
		context.SetParamValues("1")

		if assert.NoError(t, server.handleGetPost(context)) {
			assert.Equal(t, http.StatusOK, response.Code)

			responseBody := response.Body.String()
			fmt.Print(responseBody)
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(responseBody))
			if err != nil {
				log.Fatal(err)
			}

			doc.Find("article").Each(func(i int, s *goquery.Selection) {
				hasNewParagraph := s.Is("[data-testid='paragraph-1']")

				assert.Equal(t, true, hasNewParagraph)
			})
		}
	})
}

func createServer(t *testing.T) *server {
	t.Helper()

	postModel := &StubPostModel{Store: &sql.DB{}}

	server := &server{
		config:    &http.Server{},
		postModel: postModel,
	}

	return server
}

func makeRequest(t *testing.T, method string, target string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	t.Helper()

	e := echo.New()
	request := httptest.NewRequest(method, target, body)
	response := httptest.NewRecorder()
	context := e.NewContext(request, response)

	return context, response
}
