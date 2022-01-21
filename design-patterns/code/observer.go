package code

//
// 主体更新，所有对象收到通知
type SubjectI interface {
	Register(observer Observer)
	Remove(observer Observer)
	Notify()
}

type Subject struct {
	o []Observer
}

func (s Subject) Register(observer Observer) {
	s.o = append(s.o, observer)
}

func (s Subject) Remove(observer Observer) {
	for k, o := range s.o {
		if o == observer {
			s.o = append(s.o[0:k], s.o[k+1:]...)
		}
	}
}

func (s Subject) Notify() {
	for _, observer := range s.o {
		observer.Handler()
	}
}

type Observer interface {
	Handler()
}

type OA struct {
}

func (o OA) Handler() {

}

//
