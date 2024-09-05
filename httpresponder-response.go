package httpresponder

import (
	"encoding/json"
	"net/http"
)

// ES: Función para enviar una respuesta con el error.
// RU: Функция для отправки ответа с ошибкой
func OtherResponse(w http.ResponseWriter, statusCode int, message string, data interface{}, err error) {
	var errMes string
	if err != nil {
		errMes = err.Error()
	} else {
		errMes = ""
	}
	response := NewJsonResponse(message, data, errMes)
	RespondWithJSON(w, statusCode, response)
}

func ResponseData(w http.ResponseWriter, statusCode int, itemsList interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(itemsList)
}

// ----

// 200 - OK: Standard response for successful HTTP requests / Respuesta estándar para solicitudes HTTP exitosas / Стандартный ответ для успешных HTTP-запросов
func Success_OK(w http.ResponseWriter) { SUCCESS_200(w) }

// 201 - Created: The request has been fulfilled, resulting in the creation of a new resource / La solicitud se ha completado, resultando en la creación de un nuevo recurso / Запрос выполнен, что привело к созданию нового ресурса
func Success_Created(w http.ResponseWriter) { SUCCESS_201(w) }

// 202 - Accepted: The request has been accepted for processing, but the processing has not been completed / La solicitud ha sido aceptada para su procesamiento, pero el procesamiento no se ha completado / Запрос принят для обработки, но обработка еще не завершена
func Success_Accepted(w http.ResponseWriter) { SUCCESS_202(w) }

// 203 - Non-Authoritative Information: The server is returning a modified version of the origin's response / El servidor está devolviendo una versión modificada de la respuesta original / Сервер возвращает модифицированную версию исходного ответа
func Success_NonAuthoritativeInfo(w http.ResponseWriter) { SUCCESS_203(w) }

// 204 - No Content: The server successfully processed the request, and is not returning any content / El servidor procesó correctamente la solicitud y no está devolviendo ningún contenido / Сервер успешно обработал запрос и не возвращает никакого содержимого
func Success_NoContent(w http.ResponseWriter) { SUCCESS_204(w) }

// 205 - Reset Content: The server successfully processed the request and asks that the requester reset its document view / El servidor procesó correctamente la solicitud y solicita al cliente que restablezca la vista del documento / Сервер успешно обработал запрос и просит клиента сбросить отображение документа
func Success_ResetContent(w http.ResponseWriter) { SUCCESS_205(w) }

// 206 - Partial Content: The server is delivering only part of the resource due to a range header sent by the client / El servidor está entregando solo una parte del recurso debido a un encabezado de rango enviado por el cliente / Сервер передает только часть ресурса из-за заголовка диапазона, отправленного клиентом
func Success_PartialContent(w http.ResponseWriter) { SUCCESS_206(w) }

// 207 - Multi-Status (WebDAV): The response contains multiple status codes for multiple resources / La respuesta contiene múltiples códigos de estado para múltiples recursos / Ответ содержит несколько кодов состояния для нескольких ресурсов
func Success_MultiStatus(w http.ResponseWriter) { SUCCESS_207(w) }

// 208 - Already Reported (WebDAV): The members of a DAV binding have already been enumerated and are not included again / Los miembros de una vinculación DAV ya han sido enumerados y no se incluyen de nuevo / Участники связи DAV уже перечислены и повторно не включаются
func Success_AlreadyReported(w http.ResponseWriter) { SUCCESS_208(w) }

// 226 - IM Used: The server has fulfilled a request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance / El servidor ha cumplido con una solicitud para el recurso, y la respuesta es una representación del resultado de una o más manipulaciones de instancia aplicadas a la instancia actual / Сервер выполнил запрос на ресурс, и ответ является представлением результата одной или нескольких манипуляций с экземпляром, примененных к текущему экземпляру
func Success_IMUsed(w http.ResponseWriter) { SUCCESS_226(w) }

// -----------------------------------------------------

// 400 - Bad Request: The server cannot process the request due to client error / El servidor no puede procesar la solicitud debido a un error del cliente / Сервер не может обработать запрос из-за ошибки клиента
func Client_BadRequest(w http.ResponseWriter, message string) { ERROR_400(w, message) }

// 401 - Unauthorized: Authentication is required and has failed or not been provided / Se requiere autenticación y ha fallado o no se ha proporcionado / Требуется аутентификация, которая не удалась или не была предоставлена
func Client_Unauthorized(w http.ResponseWriter, message string) { ERROR_401(w, message) }

// 402 - Payment Required: Reserved for future use, currently used by some services for payment-related issues / Reservado para uso futuro, actualmente utilizado por algunos servicios para problemas relacionados con el pago / Зарезервировано для будущего использования, в настоящее время используется некоторыми сервисами для проблем, связанных с оплатой
func Client_PaymentRequired(w http.ResponseWriter, message string) { ERROR_402(w, message) }

// 403 - Forbidden: The server understands the request but refuses to authorize it / El servidor entiende la solicitud pero se niega a autorizarla / Сервер понимает запрос, но отказывается его авторизовать
func Client_Forbidden(w http.ResponseWriter, message string) { ERROR_403(w, message) }

// 404 - Not Found: The requested resource could not be found on the server / El recurso solicitado no se pudo encontrar en el servidor / Запрашиваемый ресурс не найден на сервере
func Client_NotFound(w http.ResponseWriter, message string) { ERROR_404(w, message) }

// 405 - Method Not Allowed: The request method is not supported for the requested resource / El método de solicitud no es compatible con el recurso solicitado / Метод запроса не поддерживается для запрашиваемого ресурса
func Client_MethodNotAllowed(w http.ResponseWriter, message string) { ERROR_405(w, message) }

// 406 - Not Acceptable: The requested resource is capable of generating only content not acceptable according to the Accept headers / El recurso solicitado solo puede generar contenido no aceptable según los encabezados Accept enviados en la solicitud / Запрашиваемый ресурс может генерировать только неприемлемый контент в соответствии с заголовками Accept, отправленными в запросе
func Client_NotAcceptable(w http.ResponseWriter, message string) { ERROR_406(w, message) }

// 407 - Proxy Authentication Required: The client must authenticate itself with the proxy / El cliente debe autenticarse primero con el proxy / Клиент должен сначала пройти аутентификацию через прокси
func Client_ProxyAuthenticationRequired(w http.ResponseWriter, message string) {
	ERROR_407(w, message)
}

// 408 - Request Timeout: The server timed out waiting for the request / El servidor agotó el tiempo de espera para la solicitud / Сервер ожидал запрос слишком долго
func Client_RequestTimeout(w http.ResponseWriter, message string) { ERROR_408(w, message) }

// 409 - Conflict: The request could not be processed because of conflict in the current state of the resource / La solicitud no se pudo procesar debido a un conflicto en el estado actual del recurso / Запрос не может быть обработан из-за конфликта в текущем состоянии ресурса
func Client_Conflict(w http.ResponseWriter, message string) { ERROR_409(w, message) }

// 410 - Gone: The resource requested is no longer available and will not be available again / El recurso solicitado ya no está disponible y no estará disponible de nuevo / Запрашиваемый ресурс больше недоступен и не будет доступен снова
func Client_Gone(w http.ResponseWriter, message string) { ERROR_410(w, message) }

// 411 - Length Required: The request did not specify the length of its content, which is required by the requested resource / La solicitud no especificó la longitud de su contenido, que es requerida por el recurso solicitado / В запросе не указана длина его содержимого, что требуется для запрашиваемого ресурса
func Client_LengthRequired(w http.ResponseWriter, message string) { ERROR_411(w, message) }

// 412 - Precondition Failed: The server does not meet one of the preconditions that the requester put on the request / El servidor no cumple con una de las condiciones previas que el solicitante puso en los campos del encabezado de la solicitud / Сервер не соответствует одному из предварительных условий, установленных клиентом в заголовках запроса
func Client_PreconditionFailed(w http.ResponseWriter, message string) { ERROR_412(w, message) }

// 413 - Payload Too Large: The request is more extensive than the server is willing or able to process / La solicitud es más extensa de lo que el servidor está dispuesto o puede procesar / Запрос превышает максимальный размер, который сервер готов обработать
func Client_PayloadTooLarge(w http.ResponseWriter, message string) { ERROR_413(w, message) }

// 414 - URI Too Long: The URI provided was too long for the server to process / El URI proporcionado era demasiado largo para que el servidor lo procese / Указанный URI слишком длинный для обработки сервером
func Client_URITooLong(w http.ResponseWriter, message string) { ERROR_414(w, message) }

// 415 - Unsupported Media Type: The request entity has a media type which the server or resource does not support / La entidad de la solicitud tiene un tipo de medio que el servidor o el recurso no admite / Тело запроса имеет тип медиафайла, который не поддерживается сервером или ресурсом
func Client_UnsupportedMediaType(w http.ResponseWriter, message string) { ERROR_415(w, message) }

// 416 - Range Not Satisfiable: The client asked for a portion of the file that the server cannot supply / El cliente pidió una parte del archivo que el servidor no puede suministrar / Клиент запросил часть файла, которую сервер не может предоставить
func Client_RangeNotSatisfiable(w http.ResponseWriter, message string) { ERROR_416(w, message) }

// 417 - Expectation Failed: The server cannot meet the requirements of the Expect request-header field / El servidor no puede cumplir con los requisitos del campo de encabezado Expect de la solicitud / Сервер не может соответствовать требованиям заголовка Expect запроса
func Client_ExpectationFailed(w http.ResponseWriter, message string) { ERROR_417(w, message) }

// 418 - I'm a teapot: An Easter egg status code defined as a joke in RFC 2324 / Un código de estado de broma definido en RFC 2324 / Пасхальное яйцо, определенное в статусе RFC 2324
func Client_Teapot(w http.ResponseWriter, message string) { ERROR_418(w, message) }

// 421 - Misdirected Request: The request was directed at a server that cannot produce a response / La solicitud fue dirigida a un servidor que no puede producir una respuesta / Запрос был направлен на сервер, который не может дать ответ
func Client_MisdirectedRequest(w http.ResponseWriter, message string) { ERROR_421(w, message) }

// 422 - Unprocessable Content: The request was well-formed but could not be processed / La solicitud estaba bien formada pero no pudo ser procesada / Запрос был корректно сформирован, но не может быть обработан
func Client_UnprocessableContent(w http.ResponseWriter, message string) { ERROR_422(w, message) }

// 423 - Locked: The resource being accessed is locked / El recurso al que se intenta acceder está bloqueado / Ресурс, к которому осуществляется доступ, заблокирован
func Client_Locked(w http.ResponseWriter, message string) { ERROR_423(w, message) }

// 424 - Failed Dependency: The request failed because it depended on another request that failed / La solicitud falló porque dependía de otra solicitud que falló / Запрос не выполнен, так как зависел от другого запроса, который не выполнен
func Client_FailedDependency(w http.ResponseWriter, message string) { ERROR_424(w, message) }

// 425 - Too Early: The server is unwilling to process a request that might be replayed / El servidor no está dispuesto a procesar una solicitud que podría repetirse / Сервер не хочет обрабатывать запрос, который может быть повторен
func Client_TooEarly(w http.ResponseWriter, message string) { ERROR_425(w, message) }

// 426 - Upgrade Required: The client should switch to a different protocol / El cliente debe cambiar a un protocolo diferente / Клиент должен переключиться на другой протокол
func Client_UpgradeRequired(w http.ResponseWriter, message string) { ERROR_426(w, message) }

// 428 - Precondition Required: The origin server requires the request to be conditional to prevent conflicts / El servidor de origen requiere que la solicitud sea condicional para prevenir conflictos / Сервер-источник требует, чтобы запрос был условным для предотвращения конфликтов
func Client_PreconditionRequired(w http.ResponseWriter, message string) { ERROR_428(w, message) }

// 429 - Too Many Requests: The user has sent too many requests in a given amount of time / El usuario ha enviado demasiadas solicitudes en un tiempo dado / Пользователь отправил слишком много запросов за определенное время
func Client_TooManyRequests(w http.ResponseWriter, message string) { ERROR_429(w, message) }

// 431 - Request Header Fields Too Large: The server is unwilling to process the request because the header fields are too large / El servidor no está dispuesto a procesar la solicitud porque los campos de encabezado son demasiado grandes / Сервер не желает обрабатывать запрос, поскольку поля заголовков слишком велики
func Client_RequestHeaderFieldsTooLarge(w http.ResponseWriter, message string) {
	ERROR_431(w, message)
}

// 451 - Unavailable For Legal Reasons: A server operator has received a legal demand to deny access to a resource / Un operador de servidor ha recibido una orden legal para denegar el acceso a un recurso / Оператор сервера получил юридическое требование отказать в доступе к ресурсу
func Client_UnavailableForLegalReasons(w http.ResponseWriter, message string) {
	ERROR_451(w, message)
}

// -----------------------------------------------------

// 500 - Internal Server Error: Generic error message for unexpected server conditions / Mensaje de error genérico para condiciones inesperadas del servidor / Общий текст ошибки для неожиданных ситуаций на сервере
func Server_InternalServerError(w http.ResponseWriter, message string) { ERROR_500(w, message) }

// 501 - Not Implemented: The server does not recognize the request method, or it lacks the ability to fulfill the request / El servidor no reconoce el método de solicitud, o carece de la capacidad para cumplir con la solicitud / Сервер не распознает метод запроса или не имеет возможности выполнить запрос
func Server_NotImplemented(w http.ResponseWriter, message string) { ERROR_501(w, message) }

// 502 - Bad Gateway: The server received an invalid response from the upstream server while acting as a gateway or proxy / El servidor recibió una respuesta inválida del servidor ascendente mientras actuaba como una puerta de enlace o proxy / Сервер получил неверный ответ от вышестоящего сервера, действуя как шлюз или прокси
func Server_BadGateway(w http.ResponseWriter, message string) { ERROR_502(w, message) }

// 503 - Service Unavailable: The server is currently unable to handle the request due to being overloaded or down for maintenance / El servidor no puede manejar la solicitud actualmente debido a una sobrecarga o mantenimiento / В настоящее время сервер не может обработать запрос из-за перегрузки или на техническом обслуживании
func Server_ServiceUnavailable(w http.ResponseWriter, message string) { ERROR_503(w, message) }

// 504 - Gateway Timeout: The server did not receive a timely response from the upstream server while acting as a gateway or proxy / El servidor no recibió una respuesta oportuna del servidor ascendente mientras actuaba como una puerta de enlace o proxy / Сервер не получил своевременного ответа от вышестоящего сервера, действуя как шлюз или прокси
func Server_GatewayTimeout(w http.ResponseWriter, message string) { ERROR_504(w, message) }

// 505 - HTTP Version Not Supported: The server does not support the HTTP version used in the request / El servidor no admite la versión HTTP utilizada en la solicitud / Сервер не поддерживает версию HTTP, используемую в запросе
func Server_HTTPVersionNotSupported(w http.ResponseWriter, message string) { ERROR_505(w, message) }

// 506 - Variant Also Negotiates: Transparent content negotiation results in a circular reference / La negociación de contenido transparente da como resultado una referencia circular / Прозрачные переговоры о контенте приводят к циклической ссылке
func Server_VariantAlsoNegotiates(w http.ResponseWriter, message string) { ERROR_506(w, message) }

// 507 - Insufficient Storage: The server is unable to store the representation needed to complete the request (WebDAV) / El servidor no puede almacenar la representación necesaria para completar la solicitud (WebDAV) / Сервер не может сохранить представление, необходимое для выполнения запроса (WebDAV)
func Server_InsufficientStorage(w http.ResponseWriter, message string) { ERROR_507(w, message) }

// 508 - Loop Detected: The server detected an infinite loop while processing the request (WebDAV) / El servidor detectó un bucle infinito al procesar la solicitud (WebDAV) / Сервер обнаружил бесконечный цикл при обработке запроса (WebDAV)
func Server_LoopDetected(w http.ResponseWriter, message string) { ERROR_508(w, message) }

// 510 - Not Extended: Further extensions to the request are required for the server to fulfill it / Se requieren más extensiones para que el servidor cumpla con la solicitud / Для выполнения запроса серверу требуются дополнительные расширения
func Server_NotExtended(w http.ResponseWriter, message string) { ERROR_510(w, message) }

// 511 - Network Authentication Required: The client needs to authenticate to gain network access / El cliente necesita autenticarse para obtener acceso a la red / Клиент должен пройти аутентификацию для доступа к сети
func Server_NetworkAuthenticationRequired(w http.ResponseWriter, message string) {
	ERROR_511(w, message)
}

// ----

// 200 - OK: Standard response for successful HTTP requests / Respuesta estándar para solicitudes HTTP exitosas / Стандартный ответ для успешных HTTP-запросов
func SUCCESS_200(w http.ResponseWriter) {
	message := "OK"
	response := NewJsonResponse(message, nil, "")
	RespondWithJSON(w, http.StatusOK, response)
}

// 201 - Created: The request has been fulfilled, resulting in the creation of a new resource / La solicitud se ha completado, resultando en la creación de un nuevo recurso / Запрос выполнен, что привело к созданию нового ресурса
func SUCCESS_201(w http.ResponseWriter) {
	message := "Created"
	response := NewJsonResponse(message, nil, "")
	RespondWithJSON(w, http.StatusCreated, response)
}

// 202 - Accepted: The request has been accepted for processing, but the processing has not been completed / La solicitud ha sido aceptada para su procesamiento, pero el procesamiento no se ha completado / Запрос принят для обработки, но обработка еще не завершена
func SUCCESS_202(w http.ResponseWriter) {
	message := "Accepted"
	response := NewJsonResponse(message, nil, "")
	RespondWithJSON(w, http.StatusAccepted, response)
}

// 203 - Non-Authoritative Information: The server is returning a modified version of the origin's response / El servidor está devolviendo una versión modificada de la respuesta original / Сервер возвращает модифицированную версию исходного ответа
func SUCCESS_203(w http.ResponseWriter) {
	message := "Non-Authoritative Information"
	response := NewJsonResponse(message, nil, "")
	RespondWithJSON(w, http.StatusNonAuthoritativeInfo, response)
}

// 204 - No Content: The server successfully processed the request, and is not returning any content / El servidor procesó correctamente la solicitud y no está devolviendo ningún contenido / Сервер успешно обработал запрос и не возвращает никакого содержимого
func SUCCESS_204(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// 205 - Reset Content: The server successfully processed the request and asks that the requester reset its document view / El servidor procesó correctamente la solicitud y solicita al cliente que restablezca la vista del documento / Сервер успешно обработал запрос и просит клиента сбросить отображение документа
func SUCCESS_205(w http.ResponseWriter) {
	w.WriteHeader(http.StatusResetContent)
}

// 206 - Partial Content: The server is delivering only part of the resource due to a range header sent by the client / El servidor está entregando solo una parte del recurso debido a un encabezado de rango enviado por el cliente / Сервер передает только часть ресурса из-за заголовка диапазона, отправленного клиентом
func SUCCESS_206(w http.ResponseWriter) {
	message := "Partial Content"
	response := NewJsonResponse(message, nil, "")
	RespondWithJSON(w, http.StatusPartialContent, response)
}

// 207 - Multi-Status (WebDAV): The response contains multiple status codes for multiple resources / La respuesta contiene múltiples códigos de estado para múltiples recursos / Ответ содержит несколько кодов состояния для нескольких ресурсов
func SUCCESS_207(w http.ResponseWriter) {
	message := "Multi-Status"
	response := NewJsonResponse(message, nil, "")
	RespondWithJSON(w, http.StatusMultiStatus, response)
}

// 208 - Already Reported (WebDAV): The members of a DAV binding have already been enumerated and are not included again / Los miembros de una vinculación DAV ya han sido enumerados y no se incluyen de nuevo / Участники связи DAV уже перечислены и повторно не включаются
func SUCCESS_208(w http.ResponseWriter) {
	message := "Already Reported"
	response := NewJsonResponse(message, nil, "")
	RespondWithJSON(w, http.StatusAlreadyReported, response)
}

// 226 - IM Used: The server has fulfilled a request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance / El servidor ha cumplido con una solicitud para el recurso, y la respuesta es una representación del resultado de una o más manipulaciones de instancia aplicadas a la instancia actual / Сервер выполнил запрос на ресурс, и ответ является представлением результата одной или нескольких манипуляций с экземпляром, примененных к текущему экземпляру
func SUCCESS_226(w http.ResponseWriter) {
	message := "IM Used"
	response := NewJsonResponse(message, nil, "")
	RespondWithJSON(w, http.StatusIMUsed, response)
}

// -----------------------------------------------------

// 400 - Bad Request: The server cannot process the request due to client error / El servidor no puede procesar la solicitud debido a un error del cliente / Сервер не может обработать запрос из-за ошибки клиента
func ERROR_400(w http.ResponseWriter, message string) {
	err := "Bad Request"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusBadRequest, response)
}

// 401 - Unauthorized: Authentication is required and has failed or not been provided / Se requiere autenticación y ha fallado o no se ha proporcionado / Требуется аутентификация, которая не удалась или не была предоставлена
func ERROR_401(w http.ResponseWriter, message string) {
	err := "Unauthorized"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusUnauthorized, response)
}

// 402 - Payment Required: Reserved for future use, currently used by some services for payment-related issues / Reservado para uso futuro, actualmente utilizado por algunos servicios para problemas relacionados con el pago / Зарезервировано для будущего использования, в настоящее время используется некоторыми сервисами для проблем, связанных с оплатой
func ERROR_402(w http.ResponseWriter, message string) {
	err := "Payment Required"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusPaymentRequired, response)
}

// 403 - Forbidden: The server understands the request but refuses to authorize it / El servidor entiende la solicitud pero se niega a autorizarla / Сервер понимает запрос, но отказывается его авторизовать
func ERROR_403(w http.ResponseWriter, message string) {
	err := "Forbidden"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusForbidden, response)
}

// 404 - Not Found: The requested resource could not be found on the server / El recurso solicitado no se pudo encontrar en el servidor / Запрашиваемый ресурс не найден на сервере
func ERROR_404(w http.ResponseWriter, message string) {
	err := "Not Found"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusNotFound, response)
}

// 405 - Method Not Allowed: The request method is not supported for the requested resource / El método de solicitud no es compatible con el recurso solicitado / Метод запроса не поддерживается для запрашиваемого ресурса
func ERROR_405(w http.ResponseWriter, message string) {
	err := "Method Not Allowed"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusMethodNotAllowed, response)
}

// 406 - Not Acceptable: The requested resource is capable of generating only content not acceptable according to the Accept headers / El recurso solicitado solo puede generar contenido no aceptable según los encabezados Accept enviados en la solicitud / Запрашиваемый ресурс может генерировать только неприемлемый контент в соответствии с заголовками Accept, отправленными в запросе
func ERROR_406(w http.ResponseWriter, message string) {
	err := "Not Acceptable"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusNotAcceptable, response)
}

// 407 - Proxy Authentication Required: The client must authenticate itself with the proxy / El cliente debe autenticarse primero con el proxy / Клиент должен сначала пройти аутентификацию через прокси
func ERROR_407(w http.ResponseWriter, message string) {
	err := "Proxy Authentication Required"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusProxyAuthRequired, response)
}

// 408 - Request Timeout: The server timed out waiting for the request / El servidor agotó el tiempo de espera para la solicitud / Сервер ожидал запрос слишком долго
func ERROR_408(w http.ResponseWriter, message string) {
	err := "Request Timeout"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusRequestTimeout, response)
}

// 409 - Conflict: The request could not be processed because of conflict in the current state of the resource / La solicitud no se pudo procesar debido a un conflicto en el estado actual del recurso / Запрос не может быть обработан из-за конфликта в текущем состоянии ресурса
func ERROR_409(w http.ResponseWriter, message string) {
	err := "Conflict"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusConflict, response)
}

// 410 - Gone: The resource requested is no longer available and will not be available again / El recurso solicitado ya no está disponible y no estará disponible de nuevo / Запрашиваемый ресурс больше недоступен и не будет доступен снова
func ERROR_410(w http.ResponseWriter, message string) {
	err := "Gone"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusGone, response)
}

// 411 - Length Required: The request did not specify the length of its content, which is required by the requested resource / La solicitud no especificó la longitud de su contenido, que es requerida por el recurso solicitado / В запросе не указана длина его содержимого, что требуется для запрашиваемого ресурса
func ERROR_411(w http.ResponseWriter, message string) {
	err := "Length Required"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusLengthRequired, response)
}

// 412 - Precondition Failed: The server does not meet one of the preconditions that the requester put on the request / El servidor no cumple con una de las condiciones previas que el solicitante puso en los campos del encabezado de la solicitud / Сервер не соответствует одному из предварительных условий, установленных клиентом в заголовках запроса
func ERROR_412(w http.ResponseWriter, message string) {
	err := "Precondition Failed"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusPreconditionFailed, response)
}

// 413 - Payload Too Large: The request is more extensive than the server is willing or able to process / La solicitud es más extensa de lo que el servidor está dispuesto o puede procesar / Запрос превышает максимальный размер, который сервер готов обработать
func ERROR_413(w http.ResponseWriter, message string) {
	err := "Payload Too Large"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusRequestEntityTooLarge, response)
}

// 414 - URI Too Long: The URI provided was too long for the server to process / El URI proporcionado era demasiado largo para que el servidor lo procese / Указанный URI слишком длинный для обработки сервером
func ERROR_414(w http.ResponseWriter, message string) {
	err := "URI Too Long"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusRequestURITooLong, response)
}

// 415 - Unsupported Media Type: The request entity has a media type which the server or resource does not support / La entidad de la solicitud tiene un tipo de medio que el servidor o el recurso no admite / Тело запроса имеет тип медиафайла, который не поддерживается сервером или ресурсом
func ERROR_415(w http.ResponseWriter, message string) {
	err := "Unsupported Media Type"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusUnsupportedMediaType, response)
}

// 416 - Range Not Satisfiable: The client asked for a portion of the file that the server cannot supply / El cliente pidió una parte del archivo que el servidor no puede suministrar / Клиент запросил часть файла, которую сервер не может предоставить
func ERROR_416(w http.ResponseWriter, message string) {
	err := "Range Not Satisfiable"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusRequestedRangeNotSatisfiable, response)
}

// 417 - Expectation Failed: The server cannot meet the requirements of the Expect request-header field / El servidor no puede cumplir con los requisitos del campo de encabezado Expect de la solicitud / Сервер не может соответствовать требованиям заголовка Expect запроса
func ERROR_417(w http.ResponseWriter, message string) {
	err := "Expectation Failed"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusExpectationFailed, response)
}

// 418 - I'm a teapot: An Easter egg status code defined as a joke in RFC 2324 / Un código de estado de broma definido en RFC 2324 / Пасхальное яйцо, определенное в статусе RFC 2324
func ERROR_418(w http.ResponseWriter, message string) {
	err := "I'm a teapot"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusTeapot, response)
}

// 421 - Misdirected Request: The request was directed at a server that cannot produce a response / La solicitud fue dirigida a un servidor que no puede producir una respuesta / Запрос был направлен на сервер, который не может дать ответ
func ERROR_421(w http.ResponseWriter, message string) {
	err := "Misdirected Request"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusMisdirectedRequest, response)
}

// 422 - Unprocessable Content: The request was well-formed but could not be processed / La solicitud estaba bien formada pero no pudo ser procesada / Запрос был корректно сформирован, но не может быть обработан
func ERROR_422(w http.ResponseWriter, message string) {
	err := "Unprocessable Content"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusUnprocessableEntity, response)
}

// 423 - Locked: The resource being accessed is locked / El recurso al que se intenta acceder está bloqueado / Ресурс, к которому осуществляется доступ, заблокирован
func ERROR_423(w http.ResponseWriter, message string) {
	err := "Locked"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusLocked, response)
}

// 424 - Failed Dependency: The request failed because it depended on another request that failed / La solicitud falló porque dependía de otra solicitud que falló / Запрос не выполнен, так как зависел от другого запроса, который не выполнен
func ERROR_424(w http.ResponseWriter, message string) {
	err := "Failed Dependency"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusFailedDependency, response)
}

// 425 - Too Early: The server is unwilling to process a request that might be replayed / El servidor no está dispuesto a procesar una solicitud que podría repetirse / Сервер не хочет обрабатывать запрос, который может быть повторен
func ERROR_425(w http.ResponseWriter, message string) {
	err := "Too Early"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusTooEarly, response)
}

// 426 - Upgrade Required: The client should switch to a different protocol / El cliente debe cambiar a un protocolo diferente / Клиент должен переключиться на другой протокол
func ERROR_426(w http.ResponseWriter, message string) {
	err := "Upgrade Required"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusUpgradeRequired, response)
}

// 428 - Precondition Required: The origin server requires the request to be conditional to prevent conflicts / El servidor de origen requiere que la solicitud sea condicional para prevenir conflictos / Сервер-источник требует, чтобы запрос был условным для предотвращения конфликтов
func ERROR_428(w http.ResponseWriter, message string) {
	err := "Precondition Required"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusPreconditionRequired, response)
}

// 429 - Too Many Requests: The user has sent too many requests in a given amount of time / El usuario ha enviado demasiadas solicitudes en un tiempo dado / Пользователь отправил слишком много запросов за определенное время
func ERROR_429(w http.ResponseWriter, message string) {
	err := "Too Many Requests"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusTooManyRequests, response)
}

// 431 - Request Header Fields Too Large: The server is unwilling to process the request because the header fields are too large / El servidor no está dispuesto a procesar la solicitud porque los campos de encabezado son demasiado grandes / Сервер не желает обрабатывать запрос, поскольку поля заголовков слишком велики
func ERROR_431(w http.ResponseWriter, message string) {
	err := "Request Header Fields Too Large"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusRequestHeaderFieldsTooLarge, response)
}

// 451 - Unavailable For Legal Reasons: A server operator has received a legal demand to deny access to a resource / Un operador de servidor ha recibido una orden legal para denegar el acceso a un recurso / Оператор сервера получил юридическое требование отказать в доступе к ресурсу
func ERROR_451(w http.ResponseWriter, message string) {
	err := "Unavailable For Legal Reasons"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusUnavailableForLegalReasons, response)
}

// 5xx ERRORS -----------------------------------------------------

// 500 - Internal Server Error: Generic error message for unexpected server conditions / Mensaje de error genérico para condiciones inesperadas del servidor / Общий текст ошибки для неожиданных ситуаций на сервере
func ERROR_500(w http.ResponseWriter, message string) {
	err := "Internal Server Error"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusInternalServerError, response)
}

// 501 - Not Implemented: The server does not recognize the request method, or it lacks the ability to fulfill the request / El servidor no reconoce el método de solicitud, o carece de la capacidad para cumplir con la solicitud / Сервер не распознает метод запроса или не имеет возможности выполнить запрос
func ERROR_501(w http.ResponseWriter, message string) {
	err := "Not Implemented"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusNotImplemented, response)
}

// 502 - Bad Gateway: The server received an invalid response from the upstream server while acting as a gateway or proxy / El servidor recibió una respuesta inválida del servidor ascendente mientras actuaba como una puerta de enlace o proxy / Сервер получил неверный ответ от вышестоящего сервера, действуя как шлюз или прокси
func ERROR_502(w http.ResponseWriter, message string) {
	err := "Bad Gateway"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusBadGateway, response)
}

// 503 - Service Unavailable: The server is currently unable to handle the request due to being overloaded or down for maintenance / El servidor no puede manejar la solicitud actualmente debido a una sobrecarga o mantenimiento / В настоящее время сервер не может обработать запрос из-за перегрузки или на техническом обслуживании
func ERROR_503(w http.ResponseWriter, message string) {
	err := "Service Unavailable"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusServiceUnavailable, response)
}

// 504 - Gateway Timeout: The server did not receive a timely response from the upstream server while acting as a gateway or proxy / El servidor no recibió una respuesta oportuna del servidor ascendente mientras actuaba como una puerta de enlace o proxy / Сервер не получил своевременного ответа от вышестоящего сервера, действуя как шлюз или прокси
func ERROR_504(w http.ResponseWriter, message string) {
	err := "Gateway Timeout"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusGatewayTimeout, response)
}

// 505 - HTTP Version Not Supported: The server does not support the HTTP version used in the request / El servidor no admite la versión HTTP utilizada en la solicitud / Сервер не поддерживает версию HTTP, используемую в запросе
func ERROR_505(w http.ResponseWriter, message string) {
	err := "HTTP Version Not Supported"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusHTTPVersionNotSupported, response)
}

// 506 - Variant Also Negotiates: Transparent content negotiation results in a circular reference / La negociación de contenido transparente da como resultado una referencia circular / Прозрачные переговоры о контенте приводят к циклической ссылке
func ERROR_506(w http.ResponseWriter, message string) {
	err := "Variant Also Negotiates"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusVariantAlsoNegotiates, response)
}

// 507 - Insufficient Storage: The server is unable to store the representation needed to complete the request (WebDAV) / El servidor no puede almacenar la representación necesaria para completar la solicitud (WebDAV) / Сервер не может сохранить представление, необходимое для выполнения запроса (WebDAV)
func ERROR_507(w http.ResponseWriter, message string) {
	err := "Insufficient Storage"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusInsufficientStorage, response)
}

// 508 - Loop Detected: The server detected an infinite loop while processing the request (WebDAV) / El servidor detectó un bucle infinito al procesar la solicitud (WebDAV) / Сервер обнаружил бесконечный цикл при обработке запроса (WebDAV)
func ERROR_508(w http.ResponseWriter, message string) {
	err := "Loop Detected"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusLoopDetected, response)
}

// 510 - Not Extended: Further extensions to the request are required for the server to fulfill it / Se requieren más extensiones para que el servidor cumpla con la solicitud / Для выполнения запроса серверу требуются дополнительные расширения
func ERROR_510(w http.ResponseWriter, message string) {
	err := "Not Extended"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusNotExtended, response)
}

// 511 - Network Authentication Required: The client needs to authenticate to gain network access / El cliente necesita autenticarse para obtener acceso a la red / Клиент должен пройти аутентификацию для доступа к сети
func ERROR_511(w http.ResponseWriter, message string) {
	err := "Network Authentication Required"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusNetworkAuthenticationRequired, response)
}

// Internet Infromation Services

// 440 Login Time-out : The client's session has expired and must log in again. / La sesión del cliente ha expirado y debe iniciar sesión nuevamente. / Сеанс клиента истек, и ему необходимо войти в систему снова.
func ERROR_440(w http.ResponseWriter, message string) {
	err := "Login Time-out"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusUpgradeRequired, response)
}

// 440 Login Time-out : The client's session has expired and must log in again. / La sesión del cliente ha expirado y debe iniciar sesión nuevamente. / Сеанс клиента истек, и ему необходимо войти в систему снова.
func Client_LoginTimeOut(w http.ResponseWriter, message string) { ERROR_440(w, message) }

// 449 Retry With : The server cannot honour the request because the user has not provided the required information. / El servidor no puede cumplir con la solicitud porque el usuario no ha proporcionado la información requerida. / Сервер не может выполнить запрос, так как пользователь не предоставил необходимую информацию.
func ERROR_449(w http.ResponseWriter, message string) {
	err := "Retry With"
	response := NewJsonResponse(message, nil, err)
	RespondWithJSON(w, http.StatusUpgradeRequired, response)
}

// 449 Retry With : The server cannot honour the request because the user has not provided the required information. / El servidor no puede cumplir con la solicitud porque el usuario no ha proporcionado la información requerida. / Сервер не может выполнить запрос, так как пользователь не предоставил необходимую информацию.
func Client_RetryWith(w http.ResponseWriter, message string) { ERROR_449(w, message) }
