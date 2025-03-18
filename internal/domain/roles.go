package domain

type Role string

const (
	Admin     Role = "admin"
	Moderator Role = "mod"
	Client    Role = "user"
	Guest     Role = "guest"
)
