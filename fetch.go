package ImageFetch

import (
	"encoding/xml"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// Fetch fetches a random art of a character.
func Fetch(character CharacterInfo) (CharacterArt, error) {
	var (
		response *http.Response
		data     []byte
		result   BoardPosts
		post     BoardPost
	)
	response, err = http.Get(SafeTouhouQuery + character.SearchString)
	if err != nil {
		return CharacterArt{}, err
	}
	defer func() { _ = response.Body.Close() }()
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return CharacterArt{}, err
	}
	err = xml.Unmarshal(data, &result)
	if err != nil {
		return CharacterArt{}, err
	}
	rand.Seed(time.Now().UnixNano())
	if result.Count < 1 {
		return CharacterArt{}, ErrNoArtAvailable
	}
	if result.Count < 100 {
		post = result.Posts[rand.Intn(result.Count-1)]
	} else {
		post = result.Posts[rand.Intn(99)]
	}
	return CharacterArt{
		ImageURL:  post.FileURL,
		SourceURL: post.Source,
		Character: character,
	}, nil
}
