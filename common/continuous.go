package common

type ContinuousResponse struct {
	Value   interface{}
	NextURL string
}

func WrapContinuous(doc interface{}, next string, err error) (*ContinuousResponse, error) {
	if err != nil {
		return nil, err
	} else {
		return &ContinuousResponse{Value: doc, NextURL: next}, nil
	}
}
