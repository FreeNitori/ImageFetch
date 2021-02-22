package imagefetch

import (
	"encoding/xml"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// FetchPosts fetches a posts of a character.
func FetchPosts(query string, character CharacterInfo) (BoardPosts, error) {
	var (
		response *http.Response
		data     []byte
		result   BoardPosts
	)
	response, err = http.Get(query + character.SearchString)
	if err != nil {
		return BoardPosts{}, err
	}
	defer func() { _ = response.Body.Close() }()
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return BoardPosts{}, err
	}
	err = xml.Unmarshal(data, &result)
	if err != nil {
		return BoardPosts{}, err
	}
	if result.Count < 1 {
		return BoardPosts{}, ErrNoArtAvailable
	}
	return result, nil
}

// FetchRandomPost fetches a random post.
func FetchRandomPost(posts BoardPosts) (post BoardPost) {
	rand.Seed(time.Now().UnixNano())
	if posts.Count < 100 {
		if posts.Count == 1 {
			post = posts.Posts[0]
		} else {
			post = posts.Posts[rand.Intn(posts.Count-1)]
		}
	} else {
		post = posts.Posts[rand.Intn(99)]
	}
	return
}

// Fetch fetches a random art of a character.
func Fetch(query string, character CharacterInfo) (CharacterArt, error) {
	var (
		posts BoardPosts
		post  BoardPost
	)
	posts, err = FetchPosts(query, character)
	if err != nil {
		return CharacterArt{}, err
	}
	post = FetchRandomPost(posts)
	return post.Art(), nil
}

// FetchRandom fetches art of a random character in a collection.
func FetchRandom(query string, collection CharacterCollection) (CharacterInfo, CharacterArt, error) {
	var (
		character CharacterInfo
		posts     BoardPosts
		ok        bool
	)
	character, ok = collection.CharacterRandom()
	if !ok {
		return CharacterInfo{}, CharacterArt{}, ErrNoArtAvailable
	}
	posts, err = FetchPosts(query, character)
	if err != nil {
		return CharacterInfo{}, CharacterArt{}, err
	}
	return character, FetchRandomPost(posts).Art(), nil
}
