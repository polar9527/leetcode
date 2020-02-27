package offer_2

type eagerSingleton struct{}

var insEager *eagerSingleton = &eagerSingleton{}

func GetInsEager() *eagerSingleton {
	return insEager
}
