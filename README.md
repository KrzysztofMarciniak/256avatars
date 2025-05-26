# 256avatars

Used in:  
* [minimal forum](https://github.com/KrzysztofMarciniak/minimal-forum)

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL_v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0.html)

A Go library for generating pixel-art style avatars. Provides functionality to create random and symmetric avatars, save and delete PNG files, and generate HTML tags for easy integration in web applications.

## Features

* **Random Avatars**: Generate avatars with random pixel patterns at specified dimensions.
* **Symmetric Avatars**: Create vertically mirrored avatars for visually appealing designs.
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
	keysym := "user1234"
	width, height := 256, 256
	scale := 4

	ka, err := avatarlib.GenerateKeyAvatar(key, width, height)
	if err != nil {
		log.Fatal(err)
	}
	ksa, err := avatarlib.GenerateKeySymmetricAvatar(keysym, width, height)
	if err != nil {
		log.Fatal(err)
	}

	err = avatarlib.SaveAvatar(folder, ksa, scale)
	if err != nil {
		log.Fatal(err)
	}
	err = avatarlib.SaveAvatar(folder, ka, scale)
	if err != nil {
		log.Fatal(err)
	}

	path := avatarlib.GetAvatarPath(folder, key)
	fmt.Println("Avatar saved to:", path)
	sympath := avatarlib.GetAvatarPath(folder, keysym)
	fmt.Println("Symmetric avatar saved to:", sympath)

	imgTag := avatarlib.GetAvatarHTML("/avatars/", key, width*scale, height*scale)
	fmt.Println("HTML tag:", imgTag)
	imgTagSym := avatarlib.GetAvatarHTML("/avatars/", keysym, width*scale, height*scale)
	fmt.Println("Symmetric HTML tag:", imgTagSym)

	// Delete avatar files
	// err = avatarlib.DeleteAvatar(folder, keysym)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = avatarlib.DeleteAvatar(folder, key)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
```
## API Reference

### Core Functions

`GenerateAvatar(width, height int) (*Avatar, error)`  
Generates a random avatar of the specified size. Returns an error if dimensions are invalid.

`GenerateSymmetric(width, height int) (*Avatar, error)`  
Generates a symmetric avatar (vertical mirror) of the specified size.

### Avatar Struct

```go
type Avatar struct {
    Width  int
    Height int
    Pixels []byte
}
```

**Methods:**

- `GetPixel(x, y int) bool` — Get pixel value (white = true, black = false)
- `SetPixel(x, y int, val bool)` — Set pixel value

### Key Avatar Functions

`GenerateKeyAvatar(key string, width, height int) (*KeyAvatar, error)`  
Creates an avatar wrapped in a KeyAvatar with the given key.

`GenerateKeySymmetricAvatar(key string, width, height int) (*KeyAvatar, error)`  
Creates a symmetric avatar wrapped in a KeyAvatar with the given key.

### File Operations

`SaveAvatar(folder string, ka *KeyAvatar, scale int) error`  
Saves the avatar PNG to folder/<key>.png at the specified scale, creating the folder if necessary.

`GetAvatarPath(folder, key string) string`  
Returns the file path for the avatar PNG.

`GetAvatarHTML(baseURL, key string, width, height int) string`  
Returns an HTML <img> tag pointing to the avatar.

`DeleteAvatar(folder, key string) error`  
Deletes the avatar file.

### Rendering

`RenderPNG(a *Avatar, scale int) ([]byte, error)`  
Encodes the avatar into a scaled PNG byte slice.