# 256avatars

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL_v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0.html)

A Go library for generating pixel-art style avatars. Provides functionality to create random and symmetric avatars, save and delete PNG files, and generate HTML tags for easy integration in web applications.

## Features

* **Random Avatars**: Generate avatars with random pixel patterns at specified dimensions.
* **Symmetric Avatars**: Create horizontally mirrored avatars for visually appealing designs.
* **Keyed Avatars**: Associate avatars with string keys for easy lookup and management.
* **PNG Rendering**: Render avatars to grayscale PNG images (white-on-black).
* **File Operations**: Save, retrieve path, and delete avatar files on disk.
* **HTML Integration**: Generate `<img>` tags to embed avatars in web pages.

## Installation

```bash
go get github.com/krzysztofmarciniak/256avatars
```

## Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/krzysztofmarciniak/256avatars"
)

func main() {
    // Generate a 64x64 random avatar
    avatar, err := avatarlib.GenerateAvatar(64, 64)
    if err != nil {
        log.Fatalf("failed to generate avatar: %v", err)
    }

    // Render as PNG
    pngData, err := avatarlib.RenderPNG(avatar)
    if err != nil {
        log.Fatalf("failed to render PNG: %v", err)
    }

    // Save to disk
    keyAvatar, _ := avatarlib.GenerateKeyAvatar("user123", 64, 64)
    if err := avatarlib.SaveAvatar("./avatars", keyAvatar); err != nil {
        log.Fatalf("failed to save avatar: %v", err)
    }

    // Output HTML tag
    html := avatarlib.GetAvatarHTML("/avatars/", "user123", 64, 64)
    fmt.Println(html)
}
```

## API Reference

### `GenerateAvatar(width, height int) (*Avatar, error)`

Generates a random avatar of the specified size. Returns an error if dimensions are invalid.

### `GenerateSymmetric(width, height int) (*Avatar, error)`

Generates a symmetric avatar (horizontal mirror) of the specified size.

### `Avatar` Struct

* `Width int`
* `Height int`
* `Pixels []byte`

Methods:

* `GetPixel(x, y int) bool` — Get pixel value (white = true, black = false).
* `SetPixel(x, y int, val bool)` — Set pixel value.

### `GenerateKeyAvatar(key string, width, height int) (*KeyAvatar, error)`

Creates an avatar wrapped in a `KeyAvatar` with the given key.

### `SaveAvatar(folder string, ka *KeyAvatar) error`

Saves the avatar PNG to `folder/<key>.png`, creating the folder if necessary.

### `GetAvatarPath(folder, key string) string`

Returns the file path for the avatar PNG.

### `GetAvatarHTML(baseURL, key string, width, height int) string`

Returns an HTML `<img>` tag pointing to the avatar.

### `DeleteAvatar(folder, key string) error`

Deletes the avatar file.

### `RenderPNG(a *Avatar) ([]byte, error)`

Encodes the avatar into a PNG byte slice.

## License

This project is licensed under the **AGPL v3** License. See the [LICENSE](LICENSE) file for details.
