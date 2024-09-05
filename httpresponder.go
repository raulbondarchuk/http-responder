package httpresponder

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func Prueba() {
	fmt.Println("HELLO WORLD")
}

// JsonResponse es la estructura de la respuesta en formato JSON
type JsonResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Constructor para la respuesta JsonResponse
func NewJsonResponse(message string, data interface{}, err string) JsonResponse {
	return JsonResponse{
		Message: message,
		Data:    data,
		Error:   err,
	}
}

// Responder con el formato JSON
func RespondWithJSON(w http.ResponseWriter, statusCode int, response JsonResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// ValidateFields comprueba que todos los campos pasados ​​no estén vacíos ni contengan espacios. (string, int)
func ValidateFields(fields ...interface{}) error {
	for _, field := range fields {
		value := reflect.ValueOf(field)
		switch value.Kind() {
		case reflect.String:
			str := value.String()
			if strings.TrimSpace(str) == "" || value.IsZero() {
				return fmt.Errorf("fields cannot be empty or contain spaces")
			}
		case reflect.Int:
			if value.Int() == 0 || value.IsZero() {
				return fmt.Errorf("integer fields cannot be zero")
			}
		default:
			return fmt.Errorf("unsupported field type: %s", value.Kind())
		}
	}
	return nil
}

// Verificar y responder con JSON correcto
func CheckAndRespondJSON(w http.ResponseWriter, r *http.Request, object interface{}) error {
	if r.Body == nil {
		err := errors.New("request body is empty")
		return err
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Evita la decodificación si JSON contiene campos que no están en la estructura
	err := decoder.Decode(object)
	if err != nil {
		return err
	}
	return nil
}
