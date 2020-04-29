# Go URL Shortener

Write code to:
1. Accept HTTP connections and return 200/OK [ [goal1](https://github.com/betandr/goshorten/tree/goal1) ]
2. Accept JSON containing a URL and parse: [ [goal2](https://github.com/betandr/goshorten/tree/goal2) ]
```
{ "URL": "https://www.bbc.co.uk/iplayer" }
```
3. Return a shortened URL (using a very simple scheme that doesn't need to be cryptographically secure) in a JSON response:
```
{ "ShortURL": "http://localhost:8080/abc123" }
```
4. Avoid generating short URLs for duplicate URLs
4. Accept a shortened URL and redirect to correct URL
5. Store key/values in a local JSON database

Stretch goals:
1. Count times the URL has been decoded
2. Delete URL from local database
3. Make the encoded URL "cryptographically secure"