package dto

type CreateGroupRequest struct{
	Name string `json:"name"`
	MemberIDs []string `json:"memberIds"`
}