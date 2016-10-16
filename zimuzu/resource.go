package zimuzu

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	resInfoURL     = "http://" + domain + "/gresource/%s"
	resFileListURL = "http://" + domain + "/gresource/list/%s"
)

type File struct {
	Name    string
	Format  string
	Season  int
	Episode int
	Links   []FileLink
}

type FileLink struct {
	Type string
	URL  string
}

func (c *Ctx) GetResource(id string) ([]*File, error) {
	if b, err := c.testLogin(); !b || err != nil {
		if err := c.Login(); err != nil {
			return nil, err
		}
	}

	resp, err := c.client.Get(fmt.Sprintf(resFileListURL, id))
	if err != nil {
		if strings.Contains(err.Error(), "2012: getsockopt: connection refused") {
			return nil, fmt.Errorf("ListLinks: resource not found")
		} else {
			return nil, err
		}
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("ListLinks: http %s", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, err
	}

	var files []*File

	doc.Find(".media-list li").Each(func(i int, sel *goquery.Selection) {
		f := &File{}
		f.Name = sel.Find(".fl a.lk").Text()
		f.Format = sel.AttrOr("format", "")
		f.Season, _ = strconv.Atoi(sel.AttrOr("season", "0"))
		f.Episode, _ = strconv.Atoi(sel.AttrOr("episode", "0"))

		sel.Find(".fr a").Each(func(_ int, link *goquery.Selection) {
			var url string

			if _, exists := link.Attr("thundertype"); exists {
				url = link.AttrOr("lhkakwea", "")
			} else if rel, exists := link.Attr("rel"); exists && rel == "xiaomi" {
				url = link.AttrOr("xmhref", "")
			} else {
				url = link.AttrOr("href", "")
			}

			if url != "" {
				f.Links = append(f.Links, FileLink{
					Type: link.Text(),
					URL:  url,
				})
			}
		})

		files = append(files, f)
	})

	return files, nil
}
