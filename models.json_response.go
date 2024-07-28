package main

type jsonSuccessGetResponse struct {
	RequiredPacks []requiredPackInfo `json:"requiredPacks"`
	Response      string             `json:"response"`
}

type jsonSuccessPutResponse struct {
	Response string `json:"response"`
}

type jsonErrorResponse struct {
	ErrorMessage string `json:"error"`
	Response     string `json:"response"`
}
