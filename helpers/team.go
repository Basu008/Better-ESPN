package helpers

type Team struct {
	Name    string   `json:"name,omitempty"`
	Players []string `json:"players,omitempty"`
}
