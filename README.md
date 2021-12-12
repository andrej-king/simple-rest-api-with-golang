# Simple rest api for practice with golang

## Folders:
* `internal` specific code for current service.
* `pkg` repeatable independent code.

## External package list
* [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)

## API
* GET `/users` -- user list --`200, 404, 500`
* GET `/users/:id` -- user by id --`200, 404, 500`
* POST `/users` -- create user --`201, 500`
* PUT `/users/:id` -- full update user data --`204/200, 404, 500`
* PATCH `/users/:id` -- partially update user data --`204/200, 404, 500`
* DELETE `/users/:id` -- delete user by id -- `204, 404, 500`
