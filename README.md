# MoviesAPI ![Code Quality Score](https://www.code-inspector.com/project/29263/status/svg)

## What is MoviesAPI ?

MoviesAPI is a REST API written in Golang, made for cinephiles it has the following features:

## Features

- [X] `GET` A random movie
- [X] `GET` A movie by its name
- [X] `GET` All the movies

## Documentation

### `GET` /api/movie

#### Query parameters

| Name                  | Type          | Description
|-----------------------|---------------|-------------------------|
| movie_name (Optional) | String        | Get a movie by its name |

If no parameters are specified a random movie will be returned.

#### Response fields

| Name                  | Type          | Description
|-----------------------|---------------|----------------------------------|
| movie_name            | String        | The name of the movie            |
| movie_desc            | String        | A brief description of the movie |

#### Example

soon

---

### `GET` /api/movies

#### Query parameters

| Name                  | Type          | Description
|-----------------------|---------------|-----------------------------------------------|
| max_movies (Optional) | Integer       | Specify the maximum number of movies returned |

#### Response fields

| Name                  | Type          | Description
|-----------------------|---------------|-------------------------------------------------------------------------|
| movie_name            | String        | The name of the movie                                                   |
| movie_desc            | String        | A brief description of the movie                                        |

#### Example

soon

## Credits
Thanks to [this](github.com/BonBot-Team/BonAPI) repository for its inspiration.
