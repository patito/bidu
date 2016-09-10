package model

import (
	"fmt"
	"strings"

	"github.com/patito/bidu/entity"
)

// CreateSite creates a new record of Site
func (m *Model) CreateSite(site entity.Site) error {

	query := `INSERT INTO
	          site(name, slug, physical_address, comments)
	          VALUES($1, $2, $3, $4)`

	stmt, err := m.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(site.Name,
		site.Slug,
		site.PhysicalAddress,
		site.Comments)
	if err != nil {
		return err
	}

	return nil
}

// GetSiteByName retuns the Site based on Name
func (m *Model) GetSiteByName(name string) (site entity.Site, err error) {

	query := `SELECT
	            site_id,
	            name,
	            slug,
	            comments,
	            physical_address
	          FROM site
	          WHERE name = $1`

	site = entity.Site{}
	err = m.db.QueryRow(
		query,
		name,
	).Scan(
		&site.ID,
		&site.Name,
		&site.Slug,
		&site.Comments,
		&site.PhysicalAddress)

	return
}

// GetAllSites retuns the Site based on ID
func (m *Model) GetAllSites(qstrings map[string][]string) ([]entity.Site, error) {

	filters := []string{}
	if len(qstrings) > 0 {
		if val, ok := qstrings["name"]; ok {
			filters = append(filters, "AND name = '"+val[0]+"'")
		}

		if val, ok := qstrings["slug"]; ok {
			filters = append(filters, "AND slug = '"+val[0]+"'")
		}

		if val, ok := qstrings["physical_address"]; ok {
			filters = append(filters, "AND physical_address = '"+val[0]+"'")
		}
	}

	query := fmt.Sprintf(`SELECT
	            site_id,
	            name,
	            slug,
	            comments,
	            physical_address
	          FROM site
	          WHERE site_id IS NOT NULL
	          %s`,
		strings.Join(filters, " "),
	)

	rows, err := m.db.Query(query)
	if err != nil {
		return []entity.Site{}, err
	}
	defer rows.Close()

	var sites = []entity.Site{}
	for rows.Next() {
		site := entity.Site{}
		err = rows.Scan(&site.ID,
			&site.Name,
			&site.Slug,
			&site.Comments,
			&site.PhysicalAddress)
		if err != nil {
			return []entity.Site{}, err
		}
		sites = append(sites, site)
	}

	return sites, nil
}

// DeleteSiteByName remove an specific site
func (m *Model) DeleteSiteByName(name string) error {

	query := "DELETE FROM site WHERE name = $1 returning site_id"

	var siteID int
	err := m.db.QueryRow(
		query,
		name,
	).Scan(&siteID)

	return err
}

// DeleteAllSites remove all sites
func (m *Model) DeleteAllSites() error {

	query := "DELETE FROM site"

	_, err := m.db.Exec(query)

	return err
}

// UpdateSiteByName remove an specific site
func (m *Model) UpdateSiteByName(name string, site entity.Site) error {
	query := `UPDATE site 
	          SET name = $1, slug = $2, physical_address = $3
	          WHERE name = $4 returning site_id`

	var siteID int
	err := m.db.QueryRow(
		query,
		site.Name,
		site.Slug,
		site.PhysicalAddress,
		name,
	).Scan(&siteID)

	return err
}
