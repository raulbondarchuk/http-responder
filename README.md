# httpresponder

`httpresponder` es una biblioteca en Go diseñada para simplificar la gestión de respuestas HTTP y errores comunes en aplicaciones web. La biblioteca proporciona una estructura uniforme para manejar respuestas en formato JSON y facilita la validación de datos en las solicitudes HTTP.

## Características

- **Respuestas JSON estándar**: Estructura consistente para las respuestas HTTP en formato JSON.
- **Manejo de errores HTTP**: Funciones específicas para responder con códigos de estado HTTP comunes, tanto de éxito como de error.
- **Validación de campos**: Validación sencilla de campos de entrada para evitar datos vacíos o no válidos.
- **Gestión de solicitudes HTTP**: Funciones para verificar y decodificar el cuerpo de las solicitudes HTTP.

## Instalación

Para instalar la biblioteca, usa el siguiente comando:

```bash
go get github.com/tu-usuario/http-responder
```
# Uso

## Ejemplo Básico

```go
package main

import (
	"net/http"
	resp "github.com/tu-usuario/http-responder"
	"github.com/gin-gonic/gin"
)

func main() {
	// Gin
	r := gin.Default()

	// Ruta principal
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	// Ruta que usa la biblioteca httpresponder
	r.GET("/test", func(c *gin.Context) {
		resp.Success_OK(c.Writer)
	})

	// Ruta que verifica el encabezado Auth
	r.GET("/test3", func(c *gin.Context) {
		authHeader := c.GetHeader("Auth")
		if authHeader == "" {
			resp.Client_Forbidden(c.Writer, "TOKEN DONDE ESTÁ?¿?")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Authorized",
		})
	})

	// Iniciar servidor en el puerto 3000
	r.Run(":3000")
}
```

## Validación de Campos

La función `ValidateFields` permite verificar si los campos requeridos en tu aplicación no están vacíos ni contienen espacios:

```go
err := resp.ValidateFields("nombre", 123, "apellido")
if err != nil {
	fmt.Println("Error:", err)
}
```

## Gestión de Solicitudes HTTP

Para verificar y decodificar el cuerpo de una solicitud HTTP en JSON:

```go
var datos MiEstructura
err := resp.CheckAndRespondJSON(w, r, &datos)
if err != nil {
	resp.Client_BadRequest(w, "Datos no válidos")
	return
}
```

## Documentación de las Funciones

### Respuestas de Éxito

- **`Success_OK`**/**`SUCCESS_200`**: Responde con un código HTTP 200 (OK).
- **`Success_Created`**/**`SUCCESS_201`**: Responde con un código HTTP 201 (Created).
- **`Success_Accepted`**/**`SUCCESS_202`**: Responde con un código HTTP 202 (Accepted).
- **`Success_NoContent`**/**`SUCCESS_204`**: Responde con un código HTTP 204 (No Content).
- **`Success_PartialContent`**/**`SUCCESS_206`**: Responde con un código HTTP 206 (Partial Content).
- **`Success_MultiStatus`**/**`SUCCESS_207`**: Responde con un código HTTP 207 (Multi-Status).

### Respuestas de Error del Cliente

- **`Client_BadRequest`**/**`ERROR_400`**: Responde con un código HTTP 400 (Bad Request).
- **`Client_Unauthorized`**/**`ERROR_401`**: Responde con un código HTTP 401 (Unauthorized).
- **`Client_Forbidden`**/**`ERROR_403`**: Responde con un código HTTP 403 (Forbidden).
- **`Client_NotFound`**/**`ERROR_404`**: Responde con un código HTTP 404 (Not Found).
- **`Client_MethodNotAllowed`**/**`ERROR_405`**: Responde con un código HTTP 405 (Method Not Allowed).
- **`Client_RequestTimeout`**/**`ERROR_408`**: Responde con un código HTTP 408 (Request Timeout).
- **`Client_Conflict`**/**`ERROR_409`**: Responde con un código HTTP 409 (Conflict).
- **`Client_Gone`**/**`ERROR_410`**: Responde con un código HTTP 410 (Gone).
- **`Client_LengthRequired`**/**`ERROR_411`**: Responde con un código HTTP 411 (Length Required).
- **`Client_PreconditionFailed`**/**`ERROR_412`**: Responde con un código HTTP 412 (Precondition Failed).
- **`Client_PayloadTooLarge`**/**`ERROR_413`**: Responde con un código HTTP 413 (Payload Too Large).
- **`Client_URITooLong`**/**`ERROR_414`**: Responde con un código HTTP 414 (URI Too Long).
- **`Client_UnsupportedMediaType`**/**`ERROR_415`**: Responde con un código HTTP 415 (Unsupported Media Type).
- **`Client_ExpectationFailed`**/**`ERROR_417`**: Responde con un código HTTP 417 (Expectation Failed).
- **`Client_Teapot`**/**`ERROR_418`**: Responde con un código HTTP 418 (I'm a teapot).
- **`Client_MisdirectedRequest`**/**`ERROR_421`**: Responde con un código HTTP 421 (Misdirected Request).
- **`Client_UnprocessableContent`**/**`ERROR_422`**: Responde con un código HTTP 422 (Unprocessable Content).
- **`Client_Locked`**/**`ERROR_423`**: Responde con un código HTTP 423 (Locked).
- **`Client_FailedDependency`**/**`ERROR_424`**: Responde con un código HTTP 424 (Failed Dependency).
- **`Client_TooEarly`**/**`ERROR_424`**: Responde con un código HTTP 425 (Too Early).
- **`Client_UpgradeRequired`**/**`ERROR_426`**: Responde con un código HTTP 426 (Upgrade Required).
- **`Client_PreconditionRequired`**/**`ERROR_428`**: Responde con un código HTTP 428 (Precondition Required).
- **`Client_TooManyRequests`**/**`ERROR_429`**: Responde con un código HTTP 429 (Too Many Requests).
- **`Client_RequestHeaderFieldsTooLarge`**/**`ERROR_431`**: Responde con un código HTTP 431 (Request Header Fields Too Large).
- **`Client_UnavailableForLegalReasons`**/**`ERROR_451`**: Responde con un código HTTP 451 (Unavailable For Legal Reasons).

### Respuestas de Error del Servidor

- **`Server_InternalServerError`**/**`ERROR_500`**: Responde con un código HTTP 500 (Internal Server Error).
- **`Server_NotImplemented`**/**`ERROR_501`**: Responde con un código HTTP 501 (Not Implemented).
- **`Server_BadGateway`**/**`ERROR_502`**: Responde con un código HTTP 502 (Bad Gateway).
- **`Server_ServiceUnavailable`**/**`ERROR_503`**: Responde con un código HTTP 503 (Service Unavailable).
- **`Server_GatewayTimeout`**/**`ERROR_504`**: Responde con un código HTTP 504 (Gateway Timeout).
- **`Server_HTTPVersionNotSupported`**/**`ERROR_505`**: Responde con un código HTTP 505 (HTTP Version Not Supported).
- **`Server_VariantAlsoNegotiates`**/**`ERROR_506`**: Responde con un código HTTP 506 (Variant Also Negotiates).
- **`Server_InsufficientStorage`**/**`ERROR_507`**: Responde con un código HTTP 507 (Insufficient Storage).
- **`Server_LoopDetected`**/**`ERROR_508`**: Responde con un código HTTP 508 (Loop Detected).
- **`Server_NotExtended`**/**`ERROR_510`**: Responde con un código HTTP 510 (Not Extended).
- **`Server_NetworkAuthenticationRequired`**/**`ERROR_511`**: Responde con un código HTTP 511 (Network Authentication Required).

### Internet Infromation Services

- **`Client_LoginTimeOut`**/**`ERROR_440`**: Responde con un código HTTP 440 (Login Time-out).
- **`Client_RetryWith`**/**`ERROR_449`**: Responde con un código HTTP 449 (Retry With).

## [Referencia: Lista de códigos de estado HTTP](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes)

