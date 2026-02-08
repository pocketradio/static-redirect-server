# Redirect Server

Go HTTP server with layered redirects.

```
YAML → map → fallback
```

At startup, YAML is parsed and converted into a redirect map.
The YAMLHandler builds a handler that wraps a MapHandler and returns it ( all this before requests hit the endpoint )
On each request the returned handlers ( essentially handlers created by MapHandler ) run ; no YAML parsing happens again.

Each handler checks for a match and either redirects or passes the request down.
If nothing matches, the fallback mux responds with 'Hello world'

Run:

```
go mod tidy
go run .
```

Visit `localhost:8080`.

---
