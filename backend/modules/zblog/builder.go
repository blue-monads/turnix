package zblog

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type BuildOptions struct {
	Site       *SiteModal
	Posts      []PostModal
	BaseFolder *os.Root

	ThemeLoaderFunc func() error
}

// FIXME => support more options https://gohugo.io/getting-started/configuration/#all-configuration-settings

type HugoConfig struct {
	BaseURL      string         `toml:"baseURL" json:"baseURL"`
	LanguageCode string         `toml:"languageCode" json:"languageCode"`
	Title        string         `toml:"title" json:"title"`
	Build        map[string]any `toml:"build" json:"build"`
	Theme        string         `toml:"theme" json:"theme"`
}

func SiteBuild(opts BuildOptions) error {

	opts.BaseFolder.Mkdir("site_working_dir", 0755)

	opts.BaseFolder.Mkdir("site_working_dir/theme", 0755)

	opts.BaseFolder.Mkdir("site_working_dir/content", 0755)

	opts.BaseFolder.Mkdir("site_working_dir/content/posts", 0755)

	err := opts.ThemeLoaderFunc()
	if err != nil {
		return err
	}

	config := HugoConfig{}

	if opts.Site.HugoConfig != "" {
		err = json.Unmarshal([]byte(opts.Site.HugoConfig), &config)
		if err != nil {
			return err
		}
	}

	if config.BaseURL == "" && opts.Site.Domain != "" {
		config.BaseURL = fmt.Sprintf("http://%s", opts.Site.Domain)
	}

	if config.LanguageCode == "" {
		config.LanguageCode = "en-us"
	}

	if config.Title == "" {
		config.Title = opts.Site.Name
	}

	if config.Title == "" {
		config.Title = "Simple site"
	}

	if config.Theme == "" {
		config.Theme = "default"
	}

	f, err := opts.BaseFolder.Create("site_working_dir/hugo.json")
	if err != nil {
		return err
	}

	enc := json.NewEncoder(f)
	err = enc.Encode(config)
	if err != nil {
		f.Close()
		return err
	}

	f.Close()

	for _, post := range opts.Posts {
		// if post.IsPublished {
		// 	continue
		// }

		slug := post.Slug

		if slug == "" {
			slug = slugify(post.Title)
		}

		filename := fmt.Sprintf("site_working_dir/content/posts/%s-%d.md", slug, post.ID)
		f, err := opts.BaseFolder.Create(filename)
		if err != nil {
			return err
		}

		var buf strings.Builder

		buf.WriteString("---\n")
		buf.WriteString(fmt.Sprintf("title: %s\n", post.Title))
		buf.WriteString(fmt.Sprintf("date: %s\n", post.CreatedAt.Format("2006-01-02")))
		buf.WriteString("---\n\n")
		buf.WriteString(post.Content)

		_, err = f.WriteString(buf.String())
		if err != nil {
			f.Close()
			return err
		}

		f.Close()
	}

	return nil

}
