package obj

//go:generate msgp -io=false -tests=false
type UserInfo struct {
	Name     string
	Age      int
	Birthday string
}
