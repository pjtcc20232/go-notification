package class

import (
	"net/http"

	"github.com/notification/back-end/internal/handler"
)

// Success Message Here

var SuccessHttpMsgToDeleteClass handler.HttpMsg = handler.HttpMsg{
	Msg:  "Ok Class Deleted",
	Code: http.StatusOK,
}

// Erros Message Here

var ErroHttpMsgClassIdIsRequired handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro Class ID is required",
	Code: http.StatusBadRequest,
}

var ErroHttpMsgClassNotFound handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro Class Not Found",
	Code: http.StatusNotFound,
}

var ErroHttpMsgToParseRequestClassToJson handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to parse Request Class to JSON",
	Code: http.StatusInternalServerError,
}

var ErroHttpMsgToParseResponseClassToJson handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to parse Response Class to JSON",
	Code: http.StatusInternalServerError,
}

var ErroHttpMsgToConvertingResponseClassListToJson handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to converting Response Class List to JSON",
	Code: http.StatusInternalServerError,
}

var ErroHttpMsgToInsertClass handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to Insert the Class",
	Code: http.StatusInternalServerError,
}

var ErroHttpMsgToUpdateClass handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to Update the Class",
	Code: http.StatusInternalServerError,
}

var ErroHttpMsgToDeleteClass handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to Delete the Class",
	Code: http.StatusInternalServerError,
}
