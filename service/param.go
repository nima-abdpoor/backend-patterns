package service

type SendSMSRequest struct {
	To      string
	Message string
}

type SendSMSResponse struct {
	Success bool
}
