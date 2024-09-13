package storage

var MapURLs map[string]string

func init() {
	MapURLs = make(map[string]string, 100)
}
