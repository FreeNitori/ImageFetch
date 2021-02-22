package imagefetch

import "encoding/xml"

// BoardPosts represents an array of posts returned by the image board.
type BoardPosts struct {
	XMLName xml.Name    `xml:"posts"`
	Posts   []BoardPost `xml:"post"`
	Count   int         `xml:"count,attr"`
	Offset  int         `xml:"offset,attr"`
}

// BoardPost represents a post on the image board.
type BoardPost struct {
	Height        int    `xml:"height,attr"`
	Title         string `xml:"title,attr"`
	Score         int    `xml:"score,attr"`
	FileURL       string `xml:"file_url,attr"`
	ParentID      int    `xml:"parent_id,attr"`
	SampleURL     string `xml:"sample_url,attr"`
	SampleWidth   int    `xml:"sample_width,attr"`
	SampleHeight  int    `xml:"sample_height,attr"`
	PreviewURL    string `xml:"preview_url,attr"`
	Rating        string `xml:"rating,attr"`
	Tags          string `xml:"tags,attr"`
	ID            int    `xml:"id,attr"`
	Width         int    `xml:"width,attr"`
	Change        int    `xml:"change,attr"`
	MD5           string `xml:"md5,attr"`
	CreatorID     int    `xml:"creator_id,attr"`
	HasChildren   bool   `xml:"has_children,attr"`
	CreatedAt     string `xml:"created_at,attr"`
	Status        string `xml:"status,attr"`
	Source        string `xml:"source,attr"`
	HasNotes      bool   `xml:"has_notes,attr"`
	HasComments   bool   `xml:"has_comments,attr"`
	PreviewWidth  int    `xml:"preview_width,attr"`
	PreviewHeight int    `xml:"preview_height,attr"`
}

func (post BoardPost) Art() CharacterArt {
	return CharacterArt{
		ImageURL:  post.FileURL,
		SourceURL: post.Source,
	}
}
