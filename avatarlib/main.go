package avatarlib

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
)

// AvatarGenerator defines methods for creating avatars, including symmetric variants.
type AvatarGenerator interface {
	GenerateAvatar(width, height int) (*Avatar, error)
	GenerateSymmetric(width, height int) (*Avatar, error)
}

// Avatar represents a binary image by storing its dimensions and pixel data as a bit matrix.
type Avatar struct {
	Width, Height int    // Dimensions of the avatar
	Pixels        []byte // Bit-packed pixel data (1 bit per pixel)
}

// GenerateAvatar returns a new Avatar of given dimensions filled with random pixels.
// Returns an error if dimensions are non-positive or random data cannot be read.
func GenerateAvatar(width, height int) (*Avatar, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("invalid dimensions")
	}

	bits := width * height
	bytesLen := (bits + 7) / 8
	pixels := make([]byte, bytesLen)

	if _, err := rand.Read(pixels); err != nil {
		return nil, fmt.Errorf("random generation failed: %w", err)
	}

	return &Avatar{Width: width, Height: height, Pixels: pixels}, nil
}

// GetPixel returns the boolean value of the pixel at (x, y).
// If the coordinates are out of bounds, it returns false.
func (a *Avatar) GetPixel(x, y int) bool {
	if x < 0 || x >= a.Width || y < 0 || y >= a.Height {
		return false
	}
	idx := y*a.Width + x
	byteIdx := idx / 8
	bitIdx := uint(idx % 8)
	return (a.Pixels[byteIdx] & (1 << bitIdx)) != 0
}

// SetPixel sets or clears the pixel at (x, y) to the given boolean value.
// Out-of-bounds coordinates are ignored.
func (a *Avatar) SetPixel(x, y int, val bool) {
	if x < 0 || x >= a.Width || y < 0 || y >= a.Height {
		return
	}
	idx := y*a.Width + x
	byteIdx := idx / 8
	bitIdx := uint(idx % 8)
	if val {
		a.Pixels[byteIdx] |= 1 << bitIdx
	} else {
		a.Pixels[byteIdx] &^= 1 << bitIdx
	}
}

// GenerateSymmetric returns a new Avatar with pixels randomly set in the left half
// and mirrored across the vertical center for symmetry.
// Returns an error if dimensions are non-positive or random data cannot be read.
func GenerateSymmetric(width, height int) (*Avatar, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("invalid dimensions")
	}

	a := &Avatar{Width: width, Height: height, Pixels: make([]byte, (width*height+7)/8)}
	halfWidth := (width + 1) / 2

	for y := 0; y < height; y++ {
		for x := 0; x < halfWidth; x++ {
			b := make([]byte, 1)
			if _, err := rand.Read(b); err != nil {
				return nil, fmt.Errorf("random read failed: %w", err)
			}
			val := (b[0] & 1) == 1

			a.SetPixel(x, y, val)
			mirrorX := width - 1 - x
			if mirrorX != x {
				a.SetPixel(mirrorX, y, val)
			}
		}
	}

	return a, nil
}

// KeyAvatar associates a unique string key with an Avatar instance.
type KeyAvatar struct {
	Key    string  // Unique identifier for the avatar
	Avatar *Avatar // Underlying avatar data
}

// GenerateKeyAvatar creates a random Avatar and wraps it with the provided key.
// Delegates to GenerateAvatar and returns an error on failure.
func GenerateKeyAvatar(key string, width, height int, method string) (*KeyAvatar, error) {
	var avatar *Avatar
	var err error

	switch method {
	case "symmetric":
		avatar, err = GenerateSymmetric(width, height)
	case "none", "":
		fallthrough
	default:
		avatar, err = GenerateAvatar(width, height)
	}

	if err != nil {
		return nil, err
	}
	return &KeyAvatar{Key: key, Avatar: avatar}, nil
}

// SaveAvatar renders the KeyAvatar as a PNG and writes it to folder/<key>.png.
// Creates the folder if it does not exist.
func SaveAvatar(folder string, ka *KeyAvatar) error {
	pngData, err := RenderPNG(ka.Avatar)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(folder, 0755); err != nil {
		return err
	}

	filename := filepath.Join(folder, ka.Key+".png")
	return os.WriteFile(filename, pngData, 0644)
}

// GetAvatarPath constructs the filesystem path for the avatar PNG by key.
func GetAvatarPath(folder, key string) string {
	return filepath.Join(folder, key+".png")
}

// GetAvatarHTML returns an HTML <img> tag referencing the avatar under baseURL.
func GetAvatarHTML(baseURL, key string, width, height int) string {
	return fmt.Sprintf(`<img src="%s%s.png" width="%d" height="%d" alt="Avatar %s">`,
		baseURL, key, width, height, key)
}

// DeleteAvatar removes the avatar PNG file identified by key from the folder.
func DeleteAvatar(folder, key string) error {
	filename := filepath.Join(folder, key+".png")
	return os.Remove(filename)
}

// RenderPNG encodes the Avatar into a grayscale PNG image.
// Set pixels (true) map to white and unset pixels (false) map to black.
func RenderPNG(a *Avatar) ([]byte, error) {
	img := image.NewGray(image.Rect(0, 0, a.Width, a.Height))

	for y := 0; y < a.Height; y++ {
		for x := 0; x < a.Width; x++ {
			if a.GetPixel(x, y) {
				img.SetGray(x, y, color.Gray{Y: 255})
			} else {
				img.SetGray(x, y, color.Gray{Y: 0})
			}
		}
	}

	buf := &bytes.Buffer{}
	if err := png.Encode(buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
