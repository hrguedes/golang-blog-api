
# Go Lang - Blog Api

### In progress

Project created for learning with the official Google Go Lang language documentation [Go Lang](https://go.dev/doc/)


The idea is to create a simple api where you can register posts and authors.

First I created everything in the file `main.go`, then I improved the api using the principles of SOLID.
## Roadmap

- Organize for using interfaces

- Dependency injection

- Mongo DB

- Create a rich models


## API Reference

#### Get all Authors

```http
  GET /api/author/all
```

#### Add new  Author

```http
  GET /api/author/add
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `body`      | `json` | **Required**. |

```json 
{
	"Name": "Lara Croft da Silva"
}
```
Return all of authors


#### Get all Posts

```http
  GET /api/post/all
```

#### Add new  Author

```http
  GET /api/post/add
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `body`      | `json` | **Required**. |

```json 
{
	"title": "Blue Train",
	"author": {
		"id": "bb8333a7-c4ee-4bc6-9ec5-1b4a275eae9c"
	},
	"content": "Content is here my friend"
}
```

Return all of posts

## Authors

- [@hrguedes](https://github.com/hrguedes)


## References

 - [Go Lang Documentation](https://go.dev/doc/)
 - [Gorila mux Package](https://github.com/gorilla/mux)
 - [Awesome README](https://readme.so/editor)

