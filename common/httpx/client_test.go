package httpx

import (
	"context"
	"github.com/cdumange/gocess"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestClient_Do(t *testing.T) {
	ctx := context.Background()

	t.Run("no presteps", func(t *testing.T) {
		mockHttp := new(MockDoer)
		mockHttp.On("Do", mock.Anything, mock.Anything).Return(&http.Response{}, nil).Once()
		req, err := http.NewRequest(http.MethodGet, "", nil)
		require.NoError(t, err)
		c := Client{
			http:     mockHttp,
			presteps: gocess.NewGoProcess[*http.Request](),
		}
		_, err = c.Do(ctx, req)
		assert.NoError(t, err)

		mockHttp.AssertExpectations(t)
	})

	t.Run("no presteps", func(t *testing.T) {
		URL, err := url.Parse("https://test")
		URL.Path = "/test"
		req, err := http.NewRequest(http.MethodGet, URL.String(), nil)
		require.NoError(t, err)

		mockStep := new(gocess.MockStep[*http.Request])
		mockStep.On("Execute", mock.Anything, mock.MatchedBy(func(i *http.Request) bool {
			return i.URL.Path == "/test"
		})).Return(func(ctx context.Context, i *http.Request) *http.Request {
			i.URL.Path = "/retest"
			return i
		}, nil).
			Once()

		mockHttp := new(MockDoer)
		mockHttp.On("Do", mock.Anything, mock.MatchedBy(func(r *http.Request) bool {
			return strings.Contains(r.URL.Path, "/retest")
		})).Return(&http.Response{
			StatusCode: http.StatusOK,
		}, nil).Once()

		c := Client{
			http:     mockHttp,
			presteps: gocess.NewGoProcess[*http.Request](mockStep),
		}

		_, err = c.Do(ctx, req)
		assert.NoError(t, err)

		mockStep.AssertExpectations(t)
		mockHttp.AssertExpectations(t)
	})
}
