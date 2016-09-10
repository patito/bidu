package entity

// Site stores all information about site.
type Site struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Comments        string `json:"comments"`
	PhysicalAddress string `json:"physical_address"`
}
