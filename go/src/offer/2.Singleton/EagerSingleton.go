package __Singleton

type eagerSingleton struct{}

var insEager *eagerSingleton = &eagerSingleton{}

func GetInsEager() *eagerSingleton {
	return insEager
}
