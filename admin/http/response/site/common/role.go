package common

type ToRoleByPaginate struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Summary   string `json:"summary"`
	CreatedAt string `json:"created_at"`
}

type ToRoleByInformation struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	Summary     string   `json:"summary"`
	CreatedAt   string   `json:"created_at"`
}
