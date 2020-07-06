package models

type Likes struct {
	YesLikes    []string `json:"yesLikes"`
	NoLikes     []string `json:"noLikes"`
	SmbLikes    []Like   `json:"smbLikes"`
	SympatSet   []string `json:"sympatSet"`
	SympatFetch string   `json:"sympatFetch"`
}
