package dataFetcher

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetPosts(t *testing.T) {
	mockResponse := "[{\"userId\":1,\"id\":1,\"title\":\"sunt aut facere repellat provident occaecati excepturi optio reprehenderit\",\"body\":\"quia et suscipit\\nsuscipit recusandae consequuntur expedita et cum\\nreprehenderit molestiae ut ut quas totam\\nnostrum rerum est autem sunt rem eveniet architecto\"},{\"userId\":1,\"id\":2,\"title\":\"qui est esse\",\"body\":\"est rerum tempore vitae\\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\\nqui aperiam non debitis possimus qui neque nisi nulla\"}]"
	fakeServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		_, err := writer.Write([]byte(mockResponse))
		if err != nil {
			t.Errorf("Error writing mock response: %v", err)
		}
	}))
	defer fakeServer.Close()

	res, err := convertJSONResToStruct[Post](fakeServer.URL)
	if err != nil {
		t.Errorf("Error getting posts: %v", err)
	}

	shouldReturn := []Post{
		{UserID: 1, ID: 1, Title: "sunt aut facere repellat provident occaecati excepturi optio reprehenderit", Body: "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"},
		{UserID: 1, ID: 2, Title: "qui est esse", Body: "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla"},
	}

	if !(reflect.DeepEqual(res, shouldReturn)) {
		t.Errorf("GetPosts() = \n%v\nwant\n%v", res, shouldReturn)
	}
}

func TestGetComments(t *testing.T) {
	mockResponse := "[{\"postId\":1,\"id\":1,\"name\":\"id labore ex et quam laborum\",\"email\":\"Eliseo@gardner.biz\",\"body\":\"laudantium enim quasi est quidem magnam voluptate ipsam eos\\ntempora quo necessitatibus\\ndolor quam autem quasi\\nreiciendis et nam sapiente accusantium\"},{\"postId\":1,\"id\":2,\"name\":\"quo vero reiciendis velit similique earum\",\"email\":\"Jayne_Kuhic@sydney.com\",\"body\":\"est natus enim nihil est dolore omnis voluptatem numquam\\net omnis occaecati quod ullam at\\nvoluptatem error expedita pariatur\\nnihil sint nostrum voluptatem reiciendis et\"}]"
	fakeServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		_, err := writer.Write([]byte(mockResponse))
		if err != nil {
			t.Errorf("Error writing mock response: %v", err)
		}
	}))
	defer fakeServer.Close()

	res, err := convertJSONResToStruct[Comment](fakeServer.URL)
	if err != nil {
		t.Errorf("Error getting posts: %v", err)
	}

	shouldReturn := []Comment{
		{PostID: 1, ID: 1, Name: "id labore ex et quam laborum", Email: "Eliseo@gardner.biz", Body: "laudantium enim quasi est quidem magnam voluptate ipsam eos\ntempora quo necessitatibus\ndolor quam autem quasi\nreiciendis et nam sapiente accusantium"},
		{PostID: 1, ID: 2, Name: "quo vero reiciendis velit similique earum", Email: "Jayne_Kuhic@sydney.com", Body: "est natus enim nihil est dolore omnis voluptatem numquam\net omnis occaecati quod ullam at\nvoluptatem error expedita pariatur\nnihil sint nostrum voluptatem reiciendis et"},
	}

	if !(reflect.DeepEqual(res, shouldReturn)) {
		t.Errorf("GetPosts() = \n%v\nwant\n%v", res, shouldReturn)
	}
}

func TestReturnsCorrectQueryParams(t *testing.T) {
	mockPosts := []Post{
		{UserID: 1, ID: 1, Title: "sunt aut facere repellat provident occaecati excepturi optio reprehenderit", Body: "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"},
		{UserID: 1, ID: 2, Title: "qui est esse", Body: "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla"},
	}

	shouldReturn := "fakeURL/comments?&postId=1&postId=2"

	res := getEndpointWithQueryParameters(mockPosts, "fakeURL")

	if res != shouldReturn {
		t.Errorf("GetPosts() = \n%v\nwant\n%v", res, shouldReturn)
	}

}

func TestGetPostsWComments(t *testing.T) {
	mockPosts := "[{\"userId\":1,\"id\":1,\"title\":\"sunt aut facere repellat provident occaecati excepturi optio reprehenderit\",\"body\":\"quia et suscipit\\nsuscipit recusandae consequuntur expedita et cum\\nreprehenderit molestiae ut ut quas totam\\nnostrum rerum est autem sunt rem eveniet architecto\"},{\"userId\":1,\"id\":2,\"title\":\"qui est esse\",\"body\":\"est rerum tempore vitae\\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\\nqui aperiam non debitis possimus qui neque nisi nulla\"}]"
	mockComments := "[{\"postId\":1,\"id\":1,\"name\":\"id labore ex et quam laborum\",\"email\":\"Eliseo@gardner.biz\",\"body\":\"laudantium enim quasi est quidem magnam voluptate ipsam eos\\ntempora quo necessitatibus\\ndolor quam autem quasi\\nreiciendis et nam sapiente accusantium\"},{\"postId\":1,\"id\":2,\"name\":\"quo vero reiciendis velit similique earum\",\"email\":\"Jayne_Kuhic@sydney.com\",\"body\":\"est natus enim nihil est dolore omnis voluptatem numquam\\net omnis occaecati quod ullam at\\nvoluptatem error expedita pariatur\\nnihil sint nostrum voluptatem reiciendis et\"}]"
	var requestCounter int

	fakeServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		requestCounter++
		writer.Header().Set("Content-Type", "application/json")

		var currentResponse []byte
		switch requestCounter {
		case 1:
			currentResponse = []byte(mockPosts)
		case 2:
			currentResponse = []byte(mockComments)
		}

		_, err := writer.Write(currentResponse)
		if err != nil {
			t.Errorf("Error writing mock response: %v", err)
		}
	}))
	defer fakeServer.Close()

	res, err := GetPostsWComments(1, fakeServer.URL)
	if err != nil {
		t.Errorf("Error getting posts: %v", err)
	}

	shouldReturn := []Post{
		{
			UserID: 1,
			ID:     1,
			Title:  "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
			Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
			Comments: []Comment{
				{PostID: 1, ID: 1, Name: "id labore ex et quam laborum", Email: "Eliseo@gardner.biz", Body: "laudantium enim quasi est quidem magnam voluptate ipsam eos\ntempora quo necessitatibus\ndolor quam autem quasi\nreiciendis et nam sapiente accusantium"},
				{PostID: 1, ID: 2, Name: "quo vero reiciendis velit similique earum", Email: "Jayne_Kuhic@sydney.com", Body: "est natus enim nihil est dolore omnis voluptatem numquam\net omnis occaecati quod ullam at\nvoluptatem error expedita pariatur\nnihil sint nostrum voluptatem reiciendis et"},
			},
		},
		{
			UserID: 1,
			ID:     2,
			Title:  "qui est esse",
			Body:   "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla",
		},
	}

	if !(reflect.DeepEqual(res, shouldReturn)) {
		t.Errorf("got\n%v\nwant\n%v", res, shouldReturn)
	}
}
