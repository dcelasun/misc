namespace go bug
typedef string Foo
service FooService{
	void bar (1:set<Foo> foos)
}
