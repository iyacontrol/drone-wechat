package main

import (
	"github.com/iyacontrol/drone-wechat/wechat"
)

type (
	// Repo information.
	Repo struct {
		Owner string
		Name  string
	}

	// Commit information.
	Commit struct {
		Sha     string
		Branch  string
		Link    string
		Author  string
		Email   string
		Message string
	}

	// Build information.
	Build struct {
		Tag      string
		Event    string
		Number   int
		Status   string
		Link     string
		Started  float64
		Finished float64
		PR       string
	}

	// Config for the plugin.
	Config struct {
		CorpID          string
		CorpSecret      string
		Token           string
		EncodingAESKey  string
		Text            string
	}

	// Plugin values.
	Plugin struct {
		Repo   Repo
		Commit Commit
		Build  Build
		Config Config
	}
)

// Exec executes the plugin.
func (p Plugin) Exec() error {
	if p.Config.Text == "" {
		return errors.New("text is emplty")
	}
	//配置微信参数
	config := &wechat.Config{
		CorpID:          p.Config.CorpID,
		CorpSecret:      p.Config.CorpSecret,
		Token:           p.Config.Token,
		EncodingAESKey:  p.Config.EncodingAESKey,
	}
	wc := wechat.NewWechat(config)
	wc.Send(p.Config.Text)
	return nil
}