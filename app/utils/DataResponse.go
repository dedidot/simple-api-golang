package utils

type ResponseData struct {
  Status int
  Meta interface{}
	Data interface{}
}

func ResponseMessage(statuscode int) (message interface{}) {
	elements := map[int]map[string]string{
    200: map[string]string{
      "message":"The request has succeeded. An entity corresponding to the requested resource is sent in the response.",
      "severity":"OK",
    },
    201: map[string]string{
      "message":"The request has been fulfilled and resulted in a new resource being created.",
      "saverity":"CREATED",
    },
    204: map[string]string{
      "message":"The server successfully processed the request, but is not returning any content.",
      "saverity":"NO CONTENT",
    },
    400: map[string]string{
      "message":"The request could not be understood by the server due to malformed syntax or request.",
      "saverity":"INVALID REQUEST",
    },
    404:  map[string]string{
      "message":"The server has not found anything matching the request.Boron",
      "saverity":"NOT FOUND",
    },
    500:  map[string]string{
      "message":"The server encountered an unexpected condition which prevented it from fulfilling the request.",
      "saverity":"INTERNAL SERVER ERROR",
    },
  }
  var codes interface{}
  codes = elements[statuscode]  
  return codes
}