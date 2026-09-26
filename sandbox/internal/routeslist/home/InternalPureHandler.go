package home

import (
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/api"
	"github.com/MateusMoutinhoOrg/AgnosTeset/sandbox/deps/serverdeps"
)

// InternalPureHandler answers GET /{*Rest}. Every value the route
// declares is already on entries, read off the request by the generic
// RequestHandler, and the response already carries the route's response-type.
//
// Answering — setting a status, or writing a byte, which sends a 200 — is what
// ends the chain. A handler that does neither has declined, and the next route
// matching this request runs — which is how a route becomes a middleware.
// What a middleware in front stored is on route.Locals, read with
// routeio.GetLocal. Return a failure you did not answer yourself with
// routeio.Fail; nil means "done" or "not mine".
func InternalPureHandler(sandbox *api.Sandbox, route *api.Route, entries *Entries, response *serverdeps.Response) error {

	dest := entries.Rest
	if dest == "" {
		dest = "index.html"
	}
	if sandbox.Deps.Stringsdeps.HasPrefix(dest, "/") {
		dest = sandbox.Deps.Stringsdeps.TrimPrefix(dest, "/")
	}

	fullPath := "frontend/" + dest
	content, err := sandbox.Deps.Embeddeps.ReadFile(fullPath)

	if err != nil && !sandbox.Deps.Stringsdeps.HasSuffix(dest, ".html") {
		// Test if dest + "index.html" exists as requested
		altDest := dest
		if altDest != "" && !sandbox.Deps.Stringsdeps.HasSuffix(altDest, "/") {
			altDest += "/"
		}
		altDest += "index.html"
		altPath := "frontend/" + altDest

		if content2, err2 := sandbox.Deps.Embeddeps.ReadFile(altPath); err2 == nil {
			content = content2
			err = nil
			fullPath = altPath
		}
	}

	if err != nil {
		response.SetStatus(api.StatusNotFound)
		response.Write([]byte("Not Found"))
		return nil
	}

	ext := ""
	lastDot := sandbox.Deps.Stringsdeps.LastIndex(fullPath, ".")
	lastSlash := sandbox.Deps.Stringsdeps.LastIndex(fullPath, "/")
	if lastDot != -1 && lastDot > lastSlash {
		ext = fullPath[lastDot:]
	}

	contentType := "text/plain"
	switch ext {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	case ".png":
		contentType = "image/png"
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".svg":
		contentType = "image/svg+xml"
	case ".json":
		contentType = "application/json"
	}

	response.SetHeader("Content-Type", contentType)
	response.SetStatus(api.StatusOk)
	response.Write(content)

	return nil
}
