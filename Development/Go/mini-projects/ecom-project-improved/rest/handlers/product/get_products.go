package product

import (
	"ecom-project-improved/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page, limit := addQueryParamsForPagination(w, r)
	data, err := h.svc.List(page, limit)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if data == nil {
		utils.SendData(w, "No products found", 404)
		return
	}
	utils.SendData(w, data, 200)
}

func addQueryParamsForPagination(w http.ResponseWriter, r *http.Request) (int, int) {
	// pagination steps
	// get query parameters
	// get page no
	// get limit no from query parameters 
	reqQuery := r.URL.Query()
	pageStr := reqQuery.Get("page")
	limitStr := reqQuery.Get("limit")
	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		utils.SendData(w, http.StatusInternalServerError, http.StatusInternalServerError)
		return 0, 0
	}
	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		utils.SendData(w, http.StatusBadRequest, http.StatusBadRequest)
		return 0, 0
	}
	if page == 0 {
		page = 1 // initial request
	}
	if limit == 0 {
		limit = 10 // first 10 will be load 
	}

	return int(page), int(limit)
}
