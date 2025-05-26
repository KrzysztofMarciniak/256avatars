# 256avatars

Used in:
* [minimal forum](https://github.com/KrzysztofMarciniak/minimal-forum)

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

	"github.com/krzysztofmarciniak/256avatars/avatarlib"
)

func main() {
	folder := "./avatars"
	key := "user123"
	width, height := 256, 256

	ka, err := avatarlib.GenerateKeyAvatar(key, width, height)
	if err != nil {
		log.Fatal(err)
	}

	err = avatarlib.SaveAvatar(folder, ka)
	if err != nil {
		log.Fatal(err)
	}

	path := avatarlib.GetAvatarPath(folder, key)
	fmt.Println("Avatar saved to:", path)

	imgTag := avatarlib.GetAvatarHTML("/avatars/", key, width, height)
	fmt.Println("HTML tag:", imgTag)

    // Uncomment the following lines to delete the avatar

	// err = avatarlib.DeleteAvatar(folder, key)
	// if err != nil {
	// 	log.Fatal(err)
	// }
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

### `GenerateKeySymmetricAvatar(key string, width, height int) (*KeyAvatar, error)`

Creates an symmetric avatar wrapped in a `KeyAvatar` with the given key.

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
