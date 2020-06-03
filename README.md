## RUN:
* Set Environment Variables
```
export PGSQL_BLOG_USER="postgres"
export PGSQL_BLOG_PASSWORD="postgres"
export PGSQL_BLOG_HOST="localhost"
export PGSQL_BLOG_PORT="5432"

export PGSQL_USER_USER="postgres"
export PGSQL_USER_PASSWORD="postgres"
export PGSQL_USER_HOST="localhost"
export PGSQL_USER_PORT="5432"
```
* Create DB using psql command:
```
CREATE DATABASE "go-echo-db-blog";
CREATE DATABASE "go-echo-db-user";
```
* Create tables:
```
cd go-echo-blog/src/db/blog/internal/createtable
chmod +x run.sh
./run.sh
```
```
cd go-echo-blog/src/db/user/internal/createtable
chmod +x run.sh
./run.sh
```
* Start server:
```bash
git clone https://github.com/NavenduDuari/go-echo-blog
cd go-echo-blog
go run main.go
```

## User SignUp:

* Request:
```bash
curl -H "Content-Type: application/json" -X POST http://localhost:8000/public/signup -d '{"phone":"9876543210", "email":"navendu@mail.com", "password":"mypass"}'
```
* Response:
```bash
"SignUp successful"
```
## User Login:

* Request:
```bash
curl -H "Content-Type: application/json" -X GET http://localhost:8000/public/login -d '{"userid":"9876543210", "password":"mypass"}'

```
* Response:
```json
{
  "message": "Login successful!",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImZjOTBjZWM4LWE1MDgtMTFlYS05NDk0LTAyNDJhYzExMDAwMiIsInBob25lIjoiOTg3NjU0MzIxMCIsImVtYWlsIjoibmF2ZW5kdUBtYWlsLmNvbSIsImV4cCI6MTU5MTIxMzY1N30.tlw71Alzi9aga6RBFTNIS2ZMEp2AZev1SMCEfdXY5HM"
}
```
## POST Blog:

* Request:
```bash
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImZjOTBjZWM4LWE1MDgtMTFlYS05NDk0LTAyNDJhYzExMDAwMiIsInBob25lIjoiOTg3NjU0MzIxMCIsImVtYWlsIjoibmF2ZW5kdUBtYWlsLmNvbSIsImV4cCI6MTU5MTIxMzY1N30.tlw71Alzi9aga6RBFTNIS2ZMEp2AZev1SMCEfdXY5HM" -H "Content-Type: application/json" -X POST http://localhost:8000/protected/blogs -d '{"title":"This is post title", "body":"this is post body", "categories":["cat1", "cat2"], "tags":["tag1", "tag2"]}'
```
* Response:
```bash
"Blog POST successful"
```
* Request:
```bash
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImZjOTBjZWM4LWE1MDgtMTFlYS05NDk0LTAyNDJhYzExMDAwMiIsInBob25lIjoiOTg3NjU0MzIxMCIsImVtYWlsIjoibmF2ZW5kdUBtYWlsLmNvbSIsImV4cCI6MTU5MTIxMzY1N30.tlw71Alzi9aga6RBFTNIS2ZMEp2AZev1SMCEfdXY5HM" -H "Content-Type: application/json" -X POST http://localhost:8000/protected/blogs -d '{"title":"This is another post title", "body":"this is another post body", "categories":["cat1", "cat2"], "tags":["tag1", "tag2"]}'
```
* Response:
```bash
"Blog POST successful"
```
## GET Blogs:

* Request:
```bash
curl -H "Content-Type: application/json" -X GET http://localhost:8000/public/blogs
```
* Response:
```json
{
  "blogs": [
    {
      "id": "0f51d9c8-a50c-11ea-bb4d-0242ac110002",
      "title": "This is post title",
      "body": "this is post body",
      "author": "",
      "timestamp": "2020-06-02T20:03:00.340968Z",
      "categories": [
        "cat1",
        "cat2"
      ],
      "tags": [
        "tag1",
        "tag2"
      ]
    },
    {
      "id": "ac7ba170-a50c-11ea-bb4d-0242ac110002",
      "title": "This is another post title",
      "body": "this is another post body",
      "author": "",
      "timestamp": "2020-06-02T20:07:24.018669Z",
      "categories": [
        "cat1",
        "cat2"
      ],
      "tags": [
        "tag1",
        "tag2"
      ]
    }
  ],
  "message": "Blogs FETCH successful!"
}
```
## UPDATE Blogs:

* Request:
```bash
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImZjOTBjZWM4LWE1MDgtMTFlYS05NDk0LTAyNDJhYzExMDAwMiIsInBob25lIjoiOTg3NjU0MzIxMCIsImVtYWlsIjoibmF2ZW5kdUBtYWlsLmNvbSIsImV4cCI6MTU5MTIxMzY1N30.tlw71Alzi9aga6RBFTNIS2ZMEp2AZev1SMCEfdXY5HM" -H "Content-Type: application/json" -X PUT http://localhost:8000/protected/blogs -d '{"id":"0f51d9c8-a50c-11ea-bb4d-0242ac110002", "blog":{"title":"This is updated post title", "body":"this is updated post body", "categories":["cat1", "cat2"], "tags":["tag1", "tag2"]}}
```
* Response:
```bash
"Blog UPDATE successful"
```
* Request:
```bash
curl -H "Content-Type: application/json" -X GET http://localhost:8000/public/blogs
```
* Response:
```json
{
  "blogs": [
    {
      "id": "ac7ba170-a50c-11ea-bb4d-0242ac110002",
      "title": "This is another post title",
      "body": "this is another post body",
      "author": "",
      "timestamp": "2020-06-02T20:07:24.018669Z",
      "categories": [
        "cat1",
        "cat2"
      ],
      "tags": [
        "tag1",
        "tag2"
      ]
    },
    {
      "id": "0f51d9c8-a50c-11ea-bb4d-0242ac110002",
      "title": "This is updated post title",
      "body": "this is updated post body",
      "author": "",
      "timestamp": "2020-06-02T20:15:31.937391Z",
      "categories": [
        "cat1",
        "cat2"
      ],
      "tags": [
        "tag1",
        "tag2"
      ]
    }
  ],
  "message": "Blogs FETCH successful!"
}
```
## DELETE Blogs:

* Request:
```bash
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImZjOTBjZWM4LWE1MDgtMTFlYS05NDk0LTAyNDJhYzExMDAwMiIsInBob25lIjoiOTg3NjU0MzIxMCIsImVtYWlsIjoibmF2ZW5kdUBtYWlsLmNvbSIsImV4cCI6MTU5MTIxMzY1N30.tlw71Alzi9aga6RBFTNIS2ZMEp2AZev1SMCEfdXY5HM" -H "Content-Type: application/json" -X DELETE http://localhost:8000/protected/blogs/ac7ba170-a50c-11ea-bb4d-0242ac110002
```
* Response:
```bash
"Blog DELETE successful"
```
* Request:
```bash
curl -H "Content-Type: application/json" -X GET http://localhost:8000/public/blogs
```
* Response:
```json
{
  "blogs": [
    {
      "id": "0f51d9c8-a50c-11ea-bb4d-0242ac110002",
      "title": "This is updated post title",
      "body": "this is updated post body",
      "author": "",
      "timestamp": "2020-06-02T20:15:31.937391Z",
      "categories": [
        "cat1",
        "cat2"
      ],
      "tags": [
        "tag1",
        "tag2"
      ]
    }
  ],
  "message": "Blogs FETCH successful!"
}
```



