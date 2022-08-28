package server

import (
	"fmt"
	"net/http"
)

func requestSearch(word string, lang string, langCode string, page string) (resp *http.Response) {

	url := fmt.Sprintf(
		"https://krdict.korean.go.kr/%s/dicSearch/search?nation=%s&nationCode=%s&ParaWordNo=&mainSearchWord=%s&currentPage=%s",
		lang,
		lang,
		langCode,
		word,
		page)

	client, err := createClient()
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp
}

func requestView(id string, lang string, langCode string) (resp *http.Response) {

	url := fmt.Sprintf(
		"https://krdict.korean.go.kr/%s/dicSearch/SearchView?ParaWordNo=%s&nation=%s&nationCode=%s&viewType=A&blockCount=10&viewTypes=on",
		lang,
		id,
		lang,
		langCode,
	)

	client, err := createClient()
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp
}
