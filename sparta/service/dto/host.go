package dto

type ShutdownHostDTO struct {
	HostIP string `json:"hostIP" binding:"required" message:"hostIP is required"`
}
