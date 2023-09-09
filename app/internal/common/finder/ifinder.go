package finder

type IFinder interface {
	SetUp(selfAddr string, opts any)
	Start() error
	GetAllAddrs() (addrs []string, err error)
}

const (
	IFinderPrefix = "IFinder-"
)

var (
	iFinderInstance IFinder = nil
)

type IFinderUtil struct{}

func (ifu IFinderUtil) Initial() {
	iFinderInstance = &RedisFinder{}
}

func (ifu IFinderUtil) I() IFinder {
	return iFinderInstance
}
