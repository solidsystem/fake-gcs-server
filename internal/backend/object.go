package backend

// Object represents the object that is stored within the fake server.
type Object struct {
	BucketName string `json:"-"`
	Name       string `json:"-"`
	Content    []byte
	Crc32c     string
	// Generation of this object. Used to check preconditions.
	Generation int64 `json:"generation,omitempty,string"`
}

// ID is useful for comparing objects
func (o *Object) ID() string {
	return o.BucketName + "/" + o.Name
}
