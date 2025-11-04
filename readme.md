# Go MockAPI

A simple yet flexible Go module that provides a ready-to-use mock API for books data. Perfect for testing, prototyping, or learning purposes.

## About This Project

This is my first time publishing a Go module, so bear with me on this journey. The module has some implementation examples in subfolders that you can use as references or starting points.

Through building this project, I've learned quite a bit about:
- **Module versioning**: How to properly version and publish Go modules
- **Golang embed file**: Using `embed.FS` to bundle static assets directly into the binary
- **Modular code**: Designing clean interfaces that allow flexible implementations

## Features

- Pre-populated dataset of 50 programming books
- Paginated book listing with search functionality
- Get book by ID endpoint
- Bundled static image files for book covers
- Interface-based design for easy customization

## Installation

```bash
go get github.com/anggaaryas/go-mockapi
```

## Usage

### Option 1: Bring Your Own Implementation

You can implement the data source and routing by yourself. Just implement these interfaces:

**DataSource Interface:**
```go
type DataSource interface {
    PopulateData() error
    GetBookByID(id string) (Book, error)
    GetBooks(page int, pageSize int, search string) ([]Book, error)
    GetBooksCount(search string) (int64, error)
}
```

**Router Interface:**
```go
type Router interface {
    SetupMockApiRoute(service Service) error
}
```

Then use it like this:

```go
import "github.com/anggaaryas/go-mockapi"

func main() {
    dataSource := NewYourCustomDataSource()
    router := NewYourCustomRouter()
    
    mockapi.Use(dataSource, router)
}
```

### Option 2: Use the Provided Examples

If you want to get up and running quickly, you can use the GORM and Gin implementations that come with the module.

**Install the sub-modules:**

```bash
go get github.com/anggaaryas/go-mockapi/datasource/gorm
go get github.com/anggaaryas/go-mockapi/router/ginrouter
```

**Example implementation:**

```go
package main

import (
    "github.com/anggaaryas/go-mockapi"
    gormsql "github.com/anggaaryas/go-mockapi/datasource/gorm"
    "github.com/anggaaryas/go-mockapi/router/ginrouter"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    // Setup database
    db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Setup Gin router
    r := gin.Default()

    // Create datasource and router
    dataSource := gormsql.Create(db)
    router := ginrouter.Create(r)

    // Initialize mock API
    mockapi.Use(dataSource, router)

    // Run server
    r.Run(":8080")
}
```

## API Endpoints

Once running, you'll have access to these endpoints:

- `GET /api/books` - Get paginated list of books
  - Query params: `page` (default: 1), `page_size` (default: 10), `search` (optional)
- `GET /api/books/:id` - Get a specific book by ID
- `GET /mockapi/static/image/:filename` - Access book cover images

**Example requests:**

```bash
# Get first page of books
curl http://localhost:8080/api/books?page=1&page_size=10

# Search for books
curl http://localhost:8080/api/books?search=Go

# Get specific book
curl http://localhost:8080/api/books/1
```

## Environment Variables

- `BASE_URL` - Base URL for generating cover image URLs (default: `http://localhost:8080`)

## Project Structure

```
mockapi/
├── datasource/
│   └── gorm/          # GORM implementation example
├── router/
│   └── ginrouter/     # Gin router implementation example
└── static/
    └── image/         # Embedded book cover images
```

## Contributing

This is a learning project, so feel free to open issues or submit PRs if you find any bugs or have suggestions for improvements.

## License
BSD 3-Clause License

Copyright (c) 2025 Angga Arya Saputra. All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.