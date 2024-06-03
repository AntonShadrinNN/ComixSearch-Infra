package xkcd

type ComixApiResponse struct {
	Comices map[string]string `json:"comices"`
	Error   string            `json:"error"`
}
