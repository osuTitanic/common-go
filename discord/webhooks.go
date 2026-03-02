package discord

import "time"

type Footer struct {
	Text         string  `json:"text"`
	IconURL      *string `json:"icon_url,omitempty"`
	ProxyIconURL *string `json:"proxy_icon_url,omitempty"`
}

type Image struct {
	URL      string  `json:"url"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

type Thumbnail struct {
	URL      string  `json:"url"`
	ProxyURL *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

type Video struct {
	URL    string `json:"url"`
	Height *int   `json:"height,omitempty"`
	Width  *int   `json:"width,omitempty"`
}

type Provider struct {
	URL  *string `json:"url,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Author struct {
	Name         *string `json:"name,omitempty"`
	URL          *string `json:"url,omitempty"`
	IconURL      *string `json:"icon_url,omitempty"`
	ProxyIconURL *string `json:"proxy_icon_url,omitempty"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type Embed struct {
	Title       *string    `json:"title,omitempty"`
	Type        *string    `json:"type,omitempty"`
	Description *string    `json:"description,omitempty"`
	URL         *string    `json:"url,omitempty"`
	Timestamp   *time.Time `json:"timestamp,omitempty"`
	Color       *int       `json:"color,omitempty"`
	Footer      *Footer    `json:"footer,omitempty"`
	Image       *Image     `json:"image,omitempty"`
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`
	Video       *Video     `json:"video,omitempty"`
	Provider    *Provider  `json:"provider,omitempty"`
	Author      *Author    `json:"author,omitempty"`
	Fields      []Field    `json:"fields,omitempty"`
}

func (e *Embed) AddField(name, value string, inline bool) {
	e.Fields = append(e.Fields, Field{
		Name:   name,
		Value:  value,
		Inline: inline,
	})
}

type File struct {
	Name string
	Data []byte
}

type Webhook struct {
	URL       string
	Content   *string
	Username  *string
	AvatarURL *string
	TTS       *bool
	File      *File
	Embeds    []Embed
}

func (w *Webhook) AddEmbed(embed Embed) {
	w.Embeds = append(w.Embeds, embed)
}

func (w *Webhook) SetFile(name string, data []byte) {
	w.File = &File{Name: name, Data: data}
}

// TODO: Post logic
