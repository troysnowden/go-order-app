package main

type jsonSuccessResponse struct {
	RequiredPacks []requiredPackInfo `json:"requiredPacks"`
	Response      string             `json:"response"`
}

type jsonErrorResppnse struct {
	ErrorMessage string `json:"error"`
	Response     string `json:"response"`
}
