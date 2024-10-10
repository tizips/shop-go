package common

type ToUserByPaginate struct {
	ID        string                    `json:"id"`
	Nickname  string                    `json:"nickname"`
	Username  string                    `json:"username,omitempty"`
	Mobile    string                    `json:"mobile,omitempty"`
	Email     string                    `json:"email,omitempty"`
	Roles     []ToUserByPaginateOfRoles `json:"roles"`
	IsEnable  uint8                     `json:"is_enable"`
	CreatedAt string                    `json:"created_at"`
}

type ToUserByPaginateOfRoles struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
