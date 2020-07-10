package models

type FilterOptions struct {
	MinAge   int  `json:"minAge"`
	MaxAge   int  `json:"maxAge"`
	IsOnline bool `json:"isOnline"`
	MyRegion bool `json:"myRegion"`
	MyCity   bool `json:"myCity"`
}
