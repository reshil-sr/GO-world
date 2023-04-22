package mock

import (
	"database/sql"
	"fmt"
	/* "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface" */)

//approach 1 is look up the function signature itself then create a higer order function out of OpenDB function
// and pass in as an argument, and use that instead of sql.open
type sqlOpener func(string, string) (*sql.DB, error)

//approach 2 instead of passing it in to the function
//we make a package level variable and point that to real function
var SQLOpen = sql.Open

//and in OpenDB function will call that variable instead of passing it as an arg

//approach 1
func OpenDB(user, password, addr, db string, opener sqlOpener) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@%s/%s", user, password, addr, db)
	// return sql.Open("mysql", conn) //here the sql.Open is a truble maker, we don't want to test that so we mock it
	//apprach 1
	return opener("mysql", conn)
	//approch 2
	//return SQLOpen("mysql", conn)
}

//====  I have a large inteface, I need to mock a small set of its methods ==============
// let say I want GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
// we make a mock struct and defines the fields for what we want it to return
// and embed that big huge interface, and what that does is, it say okay I'am going to satisfy this interface implicitly

/* type mockDynamoDBClient struct {
	dynamodbiface.DynamoDBAPI
	output *dynamodb.GetItemOutput
	err    error
}

// and here we define GetItem to do what we want
func (m *mockDynamoDBClient) GetItem(_ *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return m.output, m.err
}

//============= "I need to mock a method on a type" ===============
//here os.file is a concrete type and we are make use of 2 methods of that type f.Clsoe() and f.read()
func ReadFileContents(f *os.File, numBytes int) ([]byte, error) {
	defer f.Close()
	data := make([]byte, numBytes)
	_, err := f.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil

} */

//Another exm (https://stackoverflow.com/questions/19167970/mock-functions-in-go)
//Method 1: Pass get_page() as a parameter of downloader()
/* type PageGetter func(url string) string

func downloader(pageGetterFunc PageGetter) {
	// ...
	content := pageGetterFunc(BASE_URL)
	// ...
} */

//Main:

//func get_page(url string) string { /* ... */ }

/* func main() {
	downloader(get_page)
}

//Test:

func mock_get_page(url string) string {
	// mock your 'get_page()' function here
}

func TestDownloader(t *testing.T) {
	downloader(mock_get_page)
} */
//=======================================================================
