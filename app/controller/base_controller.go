package controller

import (
  "strconv"
  "net/http"
  "encoding/json"
  "log"
  "errors"
  "fmt"
  "io"
  "strings"

  "github.com/fikrirnurhidayat/codeotd/app/entity"
)

type okResponse struct {
  Status string       `json:"status"`
  Meta   interface{}  `json:"meta,omitempty"`
  Data   interface{}  `json:"data"`
}

type failResponse struct {
  Status  string  `json:"status"`
  Message string  `json:"message"`
}

type errorResponse struct {
  Status  string  `json:"status"`
  Message string  `json:"message"`
}

func ok(data interface{}, meta interface{}) *okResponse {
  return &okResponse{
    Status: "OK",
    Meta: meta,
    Data: data,
  }
}

func fail(err error) *failResponse {
  return &failResponse{
    Status: "FAIL",
    Message: err.Error(),
  }
}

func fatal(err error) *errorResponse {
  return &errorResponse{
    Status: "ERROR",
    Message: err.Error(),
  }
}

type malformedRequest struct {
    status int
    msg    string
}

func (mr *malformedRequest) Error() string {
    return mr.msg
}

// NOTE: Stolen from this https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
  if contentType := r.Header.Get("Content-Type"); contentType != "" {
    if contentType != "application/json" {
      msg := "content-type header is not application/json"
      return &malformedRequest{status: http.StatusUnsupportedMediaType, msg: msg}
    }
  }

  r.Body = http.MaxBytesReader(w, r.Body, 1048576)

  dec := json.NewDecoder(r.Body)
  dec.DisallowUnknownFields()

  err := dec.Decode(&dst)
  if err != nil {
    var syntaxError *json.SyntaxError
    var unmarshalTypeError *json.UnmarshalTypeError

    switch {
    case errors.As(err, &syntaxError):
      msg := fmt.Sprintf("request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
      return &malformedRequest{status: http.StatusBadRequest, msg: msg}

    case errors.Is(err, io.ErrUnexpectedEOF):
      msg := "request body contains badly-formed JSON"
      return &malformedRequest{status: http.StatusBadRequest, msg: msg}

    case errors.As(err, &unmarshalTypeError):
      msg := fmt.Sprintf("request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
      return &malformedRequest{status: http.StatusBadRequest, msg: msg}

    case strings.HasPrefix(err.Error(), "json: unknown field "):
      fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
      msg := fmt.Sprintf("request body contains unknown field %s", fieldName)
      return &malformedRequest{status: http.StatusBadRequest, msg: msg}

    case errors.Is(err, io.EOF):
      msg := "request body must not be empty"
      return &malformedRequest{status: http.StatusBadRequest, msg: msg}

    case err.Error() == "http: request body too large":
      msg := "request body must not be larger than 1MB"
      return &malformedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}

    default:
      return err
    }
  }

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
    msg := "request body must only contain a single JSON object"
    return &malformedRequest{status: http.StatusBadRequest, msg: msg}
  }

  return nil
}

func encodeJSONBody(w http.ResponseWriter, dst interface{}) {
  encoder := json.NewEncoder(w)
  encoder.Encode(dst)
}

func getPagination(r *http.Request) (entity.Pagination, error) {
  var pagination entity.Pagination

  pageQuery, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32) 

  if err != nil {
    if strings.EqualFold(err.Error(), "strconv.ParseInt: parsing \"\": invalid syntax") {
      pageQuery = 0
    } else {
      return pagination, errors.New("page must be a valid number")
    }
  }

  pageSizeQuery, err := strconv.ParseInt(r.URL.Query().Get("page_size"), 10, 32)

  if err != nil {
    if strings.EqualFold(err.Error(), "strconv.ParseInt: parsing \"\": invalid syntax") {
      pageSizeQuery = 0
    } else {
      return pagination, errors.New("page_size must be a valid number")
    }
  }

  pagination = entity.Pagination{
    Page: int32(pageQuery),
    PageSize: int32(pageSizeQuery),
  }

  return pagination, nil
}

func LogRequest(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		log.Printf("[INFO] %s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	})
}
