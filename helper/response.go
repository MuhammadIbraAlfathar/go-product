package helper

import "github.com/MuhammadIbraAlfathar/go-product/dto"

func Response(params dto.ResponseParams) any {
	var response any
	var status string

	if params.StatusCode >= 200 && params.StatusCode <= 299 {
		status = "Success"
	} else {
		status = "Failed"
	}

	if params.Data != nil {
		response = &dto.ResponseWithData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
			Data:    params.Data,
		}
	} else {
		response = &dto.ResponseWithoutData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
		}
	}

	return response
}
