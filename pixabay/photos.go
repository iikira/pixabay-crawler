package pixabay

import (
	"fmt"
	"github.com/json-iterator/go"
	"net/url"
	"path"
	"strconv"
	"strings"
)

// PhotoParameter 搜索文件参数
type PhotoParameter struct {
	Q             string
	Lang          string // 语言
	ImageType     string // Filter results by image type.
	Orientation   string
	Category      string
	MinWidth      int
	MinHeight     int
	Colors        string
	EditorsChoice bool
	Safesearch    bool
	Order         string
	Page          int
	PerPage       int
	Pretty        bool
}

// PhotoInfo 图片信息
type PhotoInfo struct {
	LargeImageURL   string `json:"largeImageURL"`
	WebformatHeight int    `json:"webformatHeight"`
	WebformatWidth  int    `json:"webformatWidth"`
	Likes           int    `json:"likes"`
	ImageWidth      int    `json:"imageWidth"`
	ID              int    `json:"id"`
	UserID          int    `json:"user_id"`
	ImageURL        string `json:"imageURL"`
	Views           int    `json:"views"`
	Comments        int    `json:"comments"`
	PageURL         string `json:"pageURL"`
	ImageHeight     int    `json:"imageHeight"`
	WebformatURL    string `json:"webformatURL"`
	IDHash          string `json:"id_hash"`
	Type            string `json:"type"`
	PreviewHeight   int    `json:"previewHeight"`
	Tags            string `json:"tags"`
	Downloads       int    `json:"downloads"`
	User            string `json:"user"`
	Favorites       int    `json:"favorites"`
	ImageSize       int    `json:"imageSize"`
	PreviewWidth    int    `json:"previewWidth"`
	UserImageURL    string `json:"userImageURL"`
	FullHDURL       string `json:"fullHDURL"`
	PreviewURL      string `json:"previewURL"`
}

// GetPhotos 获取图片的信息
func (p *Pixabay) GetPhotos(param *PhotoParameter) (pis []*PhotoInfo, err error) {
	uri := fmt.Sprintf("%s/?key=%s%s", p.url, p.APIKey, param.URLEncode())

	resp, err := p.Client.Req("GET", uri, nil, nil)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return
	}

	jsonData := &struct {
		Hits []*PhotoInfo `json:"hits"`
	}{}

	d := jsoniter.NewDecoder(resp.Body)
	d.Decode(jsonData)

	return jsonData.Hits, err
}

// URLEncode URL encode
func (param *PhotoParameter) URLEncode() string {
	uv := url.Values{}
	if param.Q != "" {
		uv.Set("Q", param.Q)
	}
	if param.Lang != "" {
		uv.Set("lang", param.Lang)
	}
	if param.ImageType != "" {
		uv.Set("image_type", param.ImageType)
	}
	if param.Orientation != "" {
		uv.Set("orientation", param.Orientation)
	}
	if param.Category != "" {
		uv.Set("category", param.Category)
	}
	if param.MinWidth != 0 {
		uv.Set("min_width", strconv.Itoa(param.MinWidth))
	}
	if param.MinHeight != 0 {
		uv.Set("min_height", strconv.Itoa(param.MinHeight))
	}
	if param.Colors != "" {
		uv.Set("colors", param.Colors)
	}
	if param.EditorsChoice {
		uv.Set("editors_choice", "true")
	}
	if param.Safesearch {
		uv.Set("safesearch", "true")
	}
	if param.Order != "" {
		uv.Set("order", param.Order)
	}
	if param.Page != 0 {
		uv.Set("page", strconv.Itoa(param.Page))
	}
	if param.PerPage != 0 {
		uv.Set("per_page", strconv.Itoa(param.PerPage))
	}
	if param.Pretty {
		uv.Set("pretty", "true")
	}

	return uv.Encode()
}

// Filename 生成文件名
func (pi *PhotoInfo) Filename() string {
	if pi.PreviewURL == "" {
		return path.Base(pi.ImageURL)
	}

	name := path.Base(pi.PreviewURL)

	return strings.Replace(name, "_150", "", 1)
}
