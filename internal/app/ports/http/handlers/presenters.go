package handlers

type userCreateRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

func UserSuccessResponse(status string) map[string]any {
	resp := make(map[string]any)

	resp["status"] = status
	resp["error"] = nil

	return resp
}

func UserErrorResponse(err error) map[string]any {
	resp := make(map[string]any)

	resp["data"] = nil
	resp["error"] = err

	return resp
}
