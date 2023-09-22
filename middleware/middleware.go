package middleware

import "net/http"

/* Em palavras bem simples, usar um middleware é muito útil para conseguirmos
evitar duplicidade de código, para não precisarmos ficar copiando e colando a
mesma linha para todas as funções que queremos.
É uma camada intermediária que pode ser usada para processar, modificar ou
inspecionar dados antes que eles alcancem o manipulador final da solicitação.
*/

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}
