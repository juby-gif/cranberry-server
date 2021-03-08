# cranberry-server
## Usage
 1. Add a new number
```bash
curl -X POST -d '{"number":51.34}' "http://127.0.0.1:3000/api/v1/add"
```
2. Get the sum, average and count
```bash
curl -X GET  "http://127.0.0.1:3000/api/v1/calc"
```