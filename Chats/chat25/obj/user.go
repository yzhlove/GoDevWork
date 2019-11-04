package obj

import "fmt"

//go:generate msgp -io=false -tests=false
type UserInterface interface {
	Show()
	MarshalMsg(bts []byte) (o []byte, err error)
	UnmarshalMsg(bts []byte) (o []byte, err error)
	ToString() string
}

type Student struct {
	Name string
	Age  int
}

type Teacher struct {
	Id   int
	Name string
	Age  int
}

//msgp:ignore Controller
type Controller struct {
	List []UserInterface
}

//真正序列化和反序列化的结构
type ParseController struct {
	UserInterMap map[string][][]byte
}

func NewParse() *ParseController {
	return &ParseController{UserInterMap: make(map[string][][]byte)}
}

func (s Student) Show() {
	fmt.Println("student:", s)
}

func (s Student) ToString() string {
	return "Student"
}

func (t Teacher) Show() {
	fmt.Println("teacher:", t)
}

func (t Teacher) ToString() string {
	return "Teacher"
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) EncodeMarshalMsg() ([]byte, error) {
	parse := NewParse()
	for _, it := range c.List {
		if data, err := it.MarshalMsg(nil); err != nil {
			panic(err)
		} else {
			parse.UserInterMap[it.ToString()] =
				append(parse.UserInterMap[it.ToString()], data)
		}
	}
	return parse.MarshalMsg(nil)
}

func (c *Controller) DecodeUnmarshalMsg(data []byte) {
	parse := NewParse()
	if _, err := parse.UnmarshalMsg(data); err != nil {
		panic(err)
	}
	for t, datas := range parse.UserInterMap {
		var decoed UserInterface
		switch t {
		case "Student":
			decoed = new(Student)
		case "Teacher":
			decoed = new(Teacher)
		default:
			panic("type err")
		}
		for _, data := range datas {
			if _, err := decoed.UnmarshalMsg(data); err != nil {
				continue
			}
			c.List = append(c.List, decoed)
		}
	}

	for _, utc := range c.List {
		fmt.Printf("======")
		utc.Show()
	}

}
