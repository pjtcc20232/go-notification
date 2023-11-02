package course

import (
	"net/http"

	"github.com/notification/back-end/internal/handler"
)

// Success Message Here

var SuccessHttpMsgToDeleteCourse handler.HttpMsg = handler.HttpMsg{
	Msg:  "Ok Course Deleted",
	Code: http.StatusOK,
}

// Erros Message Here

var ErroHttpMsgCourseIdIsRequired handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro Course ID is required",
	Code: http.StatusBadRequest,
}

var ErroHttpMsgCourseNotFound handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro Course Not Found",
	Code: http.StatusNotFound,
}

var ErroHttpMsgToParseRequestCourseToJson handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to parse Request Course to JSON",
	Code: http.StatusInternalServerError,
}

var ErroHttpMsgToParseResponseCourseToJson handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to parse Response Course to JSON",
	Code: http.StatusInternalServerError,
}

var ErroHttpMsgToConvertingResponseCourseListToJson handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to converting Response Course List to JSON",
	Code: http.StatusInternalServerError,
}

var ErroHttpMsgToInsertCourse handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to Insert the Course",
	Code: http.StatusInternalServerError,
}

var ErroHttpMsgToUpdateCourse handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to Update the Course",
	Code: http.StatusInternalServerError,
}

var ErroHttpMsgToDeleteCourse handler.HttpMsg = handler.HttpMsg{
	Msg:  "Erro to Delete the Course",
	Code: http.StatusInternalServerError,
}
