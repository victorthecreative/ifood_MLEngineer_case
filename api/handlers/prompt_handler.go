package handlers

import (
	"encoding/json"
	"ifood_case/api/utils"
	"net/http"
)

type PromptRequest struct {
	Prompt string `json:"prompt"`
	MaxLen int    `json:"max_len"`
}

type PromptResponse struct {
	Response string `json:"response"`
}

// @Summary Process a prompt
// @Description Process a given prompt using a Python script
// @Accept  json
// @Produce  json
// @Param   prompt body PromptRequest true "Prompt"
// @Success 200 {object} PromptResponse
// @Failure 400 {object} map[string]string
// @Router /api/v1/prompt [post]
func PromptHandler(w http.ResponseWriter, r *http.Request) {
	var req PromptRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := utils.CallPythonScript(req.Prompt, req.MaxLen)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := PromptResponse{Response: response}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
