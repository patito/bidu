package handler

import (
	"database/sql"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
	"github.com/patito/bidu/entity"
)

// CreateSite creates a new site
func (h *Handler) CreateSite(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var site entity.Site

	if err := decodeBody(r, &site); err != nil {
		respondErr(w, r, http.StatusBadRequest, "Failed to read Site from request: ", err)
		return
	}

	err := h.model.CreateSite(site)
	if err != nil {
		respondErr(w, r, http.StatusInternalServerError, "Failed to create new site: ", err)
		return
	}

	respond(w, r, http.StatusCreated, nil)
}

// GetAllSites returns all sites
func (h *Handler) GetAllSites(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u, _ := url.Parse(r.URL.String())
	queryParams := u.Query()

	sites, err := h.model.GetAllSites(queryParams)
	if err != nil {
		respondErr(w, r, http.StatusInternalServerError, "Failed to get all sites: ", err)
	}

	w.Header().Set("Content-type", "application/json")
	respond(w, r, http.StatusOK, &sites)
}

// GetSite returns a site
func (h *Handler) GetSite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	site, err := h.model.GetSiteByName(ps.ByName("name"))
	if err == sql.ErrNoRows {
		respondErr(w, r, http.StatusNotFound, "Site Not Found: ", err)
		return
	}

	if err != nil {
		respondErr(w, r, http.StatusInternalServerError, "Failed to get Site ", err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	respond(w, r, http.StatusOK, &site)
}

// DeleteSite deletes an specific site
func (h *Handler) DeleteSite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	err := h.model.DeleteSiteByName(ps.ByName("name"))
	if err == sql.ErrNoRows {
		respondErr(w, r, http.StatusNotFound, "Site Not Found: ", err)
		return
	}

	if err != nil {
		respondErr(w, r, http.StatusInternalServerError, "Failed to get Site ", err)
		return
	}

	respond(w, r, http.StatusOK, nil)
}

// DeleteAllSites deletes all sites
func (h *Handler) DeleteAllSites(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	err := h.model.DeleteAllSites()
	if err != nil {
		respondErr(w, r, http.StatusInternalServerError, "Failed to remove all sites ", err)
		return
	}

	respond(w, r, http.StatusOK, nil)
}

// UpdateSite updates an specific site
func (h *Handler) UpdateSite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var site entity.Site

	if err := decodeBody(r, &site); err != nil {
		respondErr(w, r, http.StatusBadRequest, "Failed to read Site from request: ", err)
		return
	}

	err := h.model.UpdateSiteByName(ps.ByName("name"), site)
	if err == sql.ErrNoRows {
		respondErr(w, r, http.StatusNotFound, "Site Not Found: ", err)
	}

	if err != nil {
		respondErr(w, r, http.StatusInternalServerError, "Failed to update Site ", err)
	}

	respond(w, r, http.StatusOK, nil)
}
