package handlers

import
(
	"encoding/json"
	"kt_project/internal/models"
	"kt_project/internal/repository"
	"net/http"
)

type Handler struct
{
	Repo *repository.Storage
}

func (h *Handler) DumpMetric(w http.REsponseWriter, r *http.Request)
{
	var m models.Stat

	// parse json
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil
	{
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// save to db
	h.Repo.Save(m)

	// response to agent
	w.WriteHeader(http.StatusCreated)
}
