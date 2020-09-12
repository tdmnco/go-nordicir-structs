package structs

import "time"

type announcements struct {
	Announcements []announcement `json:"announcements,omitempty"`
}

type announcement struct {
	Company         company   `json:"company,omitempty"`
	Content         content   `json:"content,omitempty"`
	Country         string    `json:"country,omitempty"`
	ID              string    `json:"id,omitempty"`
	LogoURL         string    `json:"logoUrl,omitempty"`
	Paging          paging    `json:"paging,omitempty"`
	PublicationDate time.Time `json:"publicationDate,omitempty"`
}

type attachment struct {
	MIMEType    string `json:"mimetype,omitempty"`
	Title       string `json:"title,omitempty"`
	URL         string `json:"url,omitempty"`
	SizeInBytes uint64 `json:"sizeInBytes,omitempty"`
}

type category struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type company struct {
	Name         string `json:"name,omitempty"`
	PressRoomURL string `json:"pressroomUrl,omitempty"`
}

type content struct {
	Danish  localizedContent `json:"da,omitempty"`
	English localizedContent `json:"en,omitempty"`
}

type correctionOf struct {
	ID         string `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	WebsiteURL string `json:"websiteUrl,omitempty"`
}

type exchange struct {
	Category category `json:"category,omitempty"`
	ISIN     string   `json:"isin,omitempty"`
	Name     string   `json:"name,omitempty"`
}

type id struct {
	Description string `json:"description,omitempty"`
	Example     string `json:"example,omitempty"`
	Type        string `json:"type,omitempty"`
}

type links struct {
	First    string `json:"first,omitempty"`
	Last     string `json:"last,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"prev,omitempty"`
}

type localizedContent struct {
	Attachments  []attachment `json:"attachments,omitempty"`
	CorrectionOf correctionOf `json:"correctionOf,omitempty"`
	Exchanges    []exchange   `json:"exchanges,omitempty"`
	HTML         string       `json:"html,omitempty"`
	PDFURL       string       `json:"pdfUrl,omitempty"`
	Title        string       `json:"title,omitempty"`
	WebsiteURL   string       `json:"websiteUrl,omitempty"`
}

type paging struct {
	ItemsPerPage int   `json:"itemsPerPage,omitempty"`
	Links        links `json:"links,omitempty"`
	Page         int   `json:"page,omitempty"`
	Total        int   `json:"total,omitempty"`
}

type properties struct {
	ID         id         `json:"id,omitempty"`
	Title      title      `json:"title,omitempty"`
	WebsiteURL websiteURL `json:"websiteUrl,omitempty"`
}

type title struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"title,omitempty"`
}

type websiteURL struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
}
