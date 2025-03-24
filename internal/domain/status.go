package domain

type Status string

const (
	Pending Status = "pending"
	Accepted Status = "accepted"
	Rejected Status = "rejected"
)