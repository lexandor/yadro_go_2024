package xkcd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"xkcd/pkg/utils"
)

type XkcdClient struct {
	client *http.Client
}

func NewXkcdClient(timeout time.Duration) *XkcdClient {
	return &XkcdClient{
		client: &http.Client{Timeout: timeout},
	}
}

// Get отправляет GET запрос на указанный URL
func (xkcd *XkcdClient) GetComics(url string, id int) (*Comics, error) {

	req, err := http.NewRequest("GET", url+"/"+strconv.Itoa(id)+"/info.0.json", nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return nil, err
	}
	req.Header.Add("Accept", `application/json`)
	resp, err := xkcd.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("[%d] Comics with id=%d not found", resp.StatusCode, id)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	comics, err := utils.ParseJSON[Comics](string(body))
	if err != nil {
		return nil, err
	}

	return &comics, nil
}
