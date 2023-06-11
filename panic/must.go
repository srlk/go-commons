package panic

import "go.uber.org/zap"

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustWithLogger(sl *zap.SugaredLogger, err error) {
	if err != nil {
		sl.Panic(err)
	}
}
