package filesrepo

import "github.com/manicar2093/filestores"

type FileStore interface {
	Get(filepath string) (filestores.ObjectInfo, error)
}
