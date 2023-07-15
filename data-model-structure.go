package awesomerust

type Filename = string
type Sha1Blob = string
type Sha1Tree = string
type Sha1Commit = string
type Sha1BlobOrTree = string
type Sha1Any = string

// START OMIT

type Blob struct {
	Content []byte
}

type Tree struct {
	Content map[Filename]Sha1BlobOrTree
}

type Commit struct {
	Root      Sha1Tree
	Parents   []Sha1Commit
	Author    string
	Committer string
	Message   string
}

// END OMIT
