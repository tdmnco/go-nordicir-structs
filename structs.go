package structs

import "time"

// CompanyAnnouncements contains an array of company announcements
type CompanyAnnouncements struct {
	CompanyAnnouncements []CompanyAnnouncement `json:"announcements,omitempty"`
}

// CompanyAnnouncement represents a single company announcement from the Nordic IR service
type CompanyAnnouncement struct {
	Company         Company   `json:"company,omitempty"`
	Content         Content   `json:"content,omitempty"`
	Country         string    `json:"country,omitempty"`
	ID              string    `json:"id,omitempty"`
	Identifier      string    `json:"identifier,omitempty"`
	LogoURL         string    `json:"logoUrl,omitempty"`
	Paging          Paging    `json:"paging,omitempty"`
	PublicationDate time.Time `json:"publicationDate,omitempty"`
}

// Attachment contains any attachment related to a given company message (PDF files etc.)
type Attachment struct {
	MIMEType    string `json:"mimetype,omitempty"`
	Title       string `json:"title,omitempty"`
	URL         string `json:"url,omitempty"`
	SizeInBytes uint64 `json:"sizeInBytes,omitempty"`
}

// Category contains the name and ID of a company announcement category
type Category struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Company is a representation of the given company that has published an announcement
type Company struct {
	Name         string `json:"name,omitempty"`
	PressRoomURL string `json:"pressroomUrl,omitempty"`
}

// Content holds any amount of localized content of a company announcement
type Content struct {
	Danish  LocalizedContent `json:"da,omitempty"`
	English LocalizedContent `json:"en,omitempty"`
}

// CorrectionOf is used to correct any given company announcement
type CorrectionOf struct {
	ID         string `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	WebsiteURL string `json:"websiteUrl,omitempty"`
}

// Exchange information including ISIN number, category and name
type Exchange struct {
	Category Category `json:"category,omitempty"`
	ISIN     string   `json:"isin,omitempty"`
	Name     string   `json:"name,omitempty"`
}

// ID metadata of company announcement properties
type ID struct {
	Description string `json:"description,omitempty"`
	Example     string `json:"example,omitempty"`
	Type        string `json:"type,omitempty"`
}

// Links used in relation with paging of company announcements
type Links struct {
	First    string `json:"first,omitempty"`
	Last     string `json:"last,omitempty"`
	Next     string `json:"next,omitempty"`
	Previous string `json:"prev,omitempty"`
}

// LocalizedContent contains localized content of a company announcement in a given language
type LocalizedContent struct {
	Attachments  []Attachment `json:"attachments,omitempty"`
	CorrectionOf CorrectionOf `json:"correctionOf,omitempty"`
	Exchanges    []Exchange   `json:"exchanges,omitempty"`
	HTML         string       `json:"html,omitempty"`
	PDFURL       string       `json:"pdfUrl,omitempty"`
	Title        string       `json:"title,omitempty"`
	WebsiteURL   string       `json:"websiteUrl,omitempty"`
}

// Paging holds metadata about the paging of company announcements coming from the Nordic IR endpoints
type Paging struct {
	ItemsPerPage int   `json:"itemsPerPage,omitempty"`
	Links        Links `json:"links,omitempty"`
	Page         int   `json:"page,omitempty"`
	Total        int   `json:"total,omitempty"`
}

// Properties of a company announcement
type Properties struct {
	ID         ID         `json:"id,omitempty"`
	Title      Title      `json:"title,omitempty"`
	WebsiteURL WebsiteURL `json:"websiteUrl,omitempty"`
}

// Title of a company announcement
type Title struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"title,omitempty"`
}

// WebsiteURL contained in the properties of a company announcement
type WebsiteURL struct {
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
}
