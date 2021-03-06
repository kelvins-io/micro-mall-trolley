package vars

type EmailConfigSettingS struct {
	Enable   bool   `json:"enable"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type MongoDBSettingS struct {
	Uri         string `json:"uri"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Database    string `json:"database"`
	AuthSource  string `json:"auth_source"`
	MaxPoolSize int    `json:"max_pool_size"`
	MinPoolSize int    `json:"min_pool_size"`
}
