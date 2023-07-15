package awesomerust

// START OMIT

type Object interface {
	Blob | Tree | Commit
}

type ObjectStore[O Object] struct {
	objects map[Sha1Any]O
}

// END OMIT

func NewObjectStore[O Object]() *ObjectStore[O] {
	return &ObjectStore[O]{
		objects: make(map[Sha1Any]O),
	}
}

// START2 OMIT

func (store *ObjectStore[O]) RetrieveObject(sha1 Sha1Any) (O, bool) {
	object, ok := store.objects[sha1]
	return object, ok
}

func (store *ObjectStore[O]) StoreObject(object O) {
	objectId := CalculateSha1(object)
	store.objects[objectId] = object
}

func CalculateSha1[O Object](obj O) Sha1Any {
	return "IMAGINE CALCULATING SHA-1"
}

// END2 OMIT
