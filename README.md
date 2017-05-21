# Midpoint service

Calculates geographic middle for an array of locations.

Example request for `Den Bosch` and `Amsterdam` :

```
curl -X POST \
  http://localhost:8080/ \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '[{
	"lat": 51.6978,
	"long": 5.3037
},
{
	"lat":52.3702,
	"long": 4.8952
}]'
```