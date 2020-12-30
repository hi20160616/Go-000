package biz

type Greeter struct {
	ID   int32
	Name string
	Msg  string
}

type GreeterCase struct {
	repo GreeterRepo
}

type GreeterRepo interface {
	SayHi(*Greeter) int32
}

func (g *GreeterCase) SetID(gg *Greeter) {
	// GreeterCase.GreeterRepo.SayHi(*Greeter)
	// So, there, gg.ID's value decided by *Greeter in GreeterRepo
	// In other words, GreeterRepo.SayHi(*Greeter) decide the value
	// So, implement the interface can set the value, Ioc done.
	gg.ID = g.repo.SayHi(gg)
}

func NewGreeterCase(repo GreeterRepo) *GreeterCase {
	return &GreeterCase{repo: repo}
}
