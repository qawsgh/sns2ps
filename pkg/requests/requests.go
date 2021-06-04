// Use local files to create json instead of making web requests
package requests

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type HTTPError struct {
	Err        error
	StatusCode int    // http status code
	URL        string // URL
}

func (f *HTTPError) Error() string {
	return fmt.Sprintf("HTTP StatusCode: %d url: %v err: %v", f.StatusCode, f.URL, f.Err)
}

// fileRequest opens a file and reads the content, returning a byteValue.
// It is intended to be called by any function that needs to get json content for
// testing instead of using the web.
func FileRequest(filename string) []byte {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

// webRequest will request content from the specified Shoot 'n Score It API endpoint.
// Endpoints are provided as the URL string and should be the complete URL including
// the match ID, and type of request (competitor, match, squad).
// username and password values are used to authenticate to Shoot 'n Score It.
func WebRequest(url string, username string, password string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.SetBasicAuth(username, password)
	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		fmt.Printf("Could not connect to %v - check your Internet connection and make sure you can access the site in a browser\n", url)
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			URL:        url,
			Err:        errors.New("could not connect - check your Internet connection and make sure you can access the site in a browser"),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not connect to %v - check your Internet connection and make sure you can access the site in a browser\n", url)
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			URL:        url,
			Err:        errors.New("could not connect - check your Internet connection and make sure you can access the site in a browser"),
		}
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == 401 {
			fmt.Println("Unable to get details for this match - please check your username and password and try again.")
			return nil, &HTTPError{
				StatusCode: resp.StatusCode,
				URL:        url,
				Err:        errors.New("unable to get details for this match - please check your username and password and try again"),
			}
		} else if resp.StatusCode == 404 {
			fmt.Println("Unable to get details for this match - please check your matchID and try again.")
			return nil, &HTTPError{
				StatusCode: resp.StatusCode,
				URL:        url,
				Err:        errors.New("unable to get details for this match - please check your matchID and try again"),
			}
		} else {
			fmt.Println("Unknown error: ", resp.StatusCode)
			return nil, &HTTPError{
				StatusCode: resp.StatusCode,
				URL:        url,
				Err:        errors.New("unknown error - check StatusCode"),
			}
		}
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}
	return body, nil
}
