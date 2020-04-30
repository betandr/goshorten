# Go URL Shortener

Write code to:
1. Accept HTTP connections and return 200/OK [ `go run ./goal1` | `go test ./goal1` ]
2. Accept JSON containing a URL and parse. [ `go run ./goal2` | `go test ./goal2` ]
```
{ "URL": "https://www.bbc.co.uk/iplayer" }
```
3. Return a shortened URL (using a very simple scheme that doesn't need to be cryptographically secure) in a JSON response [ `go run ./goal3` | `go test ./goal3` ]
```
{ "ShortURL": "http://localhost:8080/1" }
```
4. Avoid generating short URLs for duplicate URLs [ `go run ./goal4` | `go test ./goal4` ]
5. Accept a shortened URL and redirect to correct URL

Stretch goals:
1. Store key/values in a local JSON database
2. Count times the URL has been decoded
3. Delete URL from local database
4. Make the encoded URL "cryptographically secure"