## Example datastore file as JSON


This is example output as json
```json
{
  "tables": {
    "users": {
      "id": "int;auto-increment",
      "email": "string",
      "password": "[]byte"
    },
    "posts": {
      "id": "int;auto-increment",
      "title": "string",
      "content": "string"
    }
  },
  "data": {
    "users": [
      {
        "id": 0,
        "name": "Test User",
        "email": "test@gmail.com",
        "password": "$kl1kl.34k@#k3r2"
      }
    ],
    "posts": [
      {
        "id": 0,
        "title": "First Post",
        "content": "This is the first post"
      },
      {
        "id": 1,
        "title": "Contentless Post",
        "content": ""
      }
    ]
  }
}
```
