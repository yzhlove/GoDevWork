package main

type HelloService struct{}

func (p *HelloService) Hello(request *String, replay *String) error {
	replay.Value = "hello :" + request.GetValue()
	return nil
}

func main() {

}
