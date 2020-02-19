//Please include GOPATH for execution of this package

package usfield

type Employer struct {
	employee //Struct embedding is done here
	Dismiss bool
}

//Here, the struct name is started with small letter which means that this struct is unexported and invisible to other pacakges
//But, the fields inside it start with capital letter i.e. the fields are exported and visible and usable by other packages
type employee struct {
	Name string
	Number int
}

