package utils

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/arlert/malcolm-ui/dist"
// 	assetfs "github.com/elazarl/go-bindata-assetfs"
// )

// type binaryFileSystem struct {
// 	fs http.FileSystem
// }

// func (b *binaryFileSystem) Open(name string) (http.File, error) {
// 	return b.fs.Open(name)
// }

// func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {

// 	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
// 		if _, err := b.fs.Open(p); err != nil {
// 			return false
// 		}
// 		return true
// 	}
// 	return false
// }

// func Frontend(root string) *binaryFileSystem {
// 	fs := &assetfs.AssetFS{
// 		Asset:     dist.Asset,
// 		AssetDir:  dist.AssetDir,
// 		AssetInfo: dist.AssetInfo,
// 		Prefix:    root,
// 	}

// 	return &binaryFileSystem{
// 		fs,
// 	}
// }
