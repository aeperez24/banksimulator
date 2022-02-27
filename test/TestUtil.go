package test

import "testing"

type AssertHelper struct {
	T *testing.T
}

func (helper AssertHelper) Assert(expected interface{}, actual interface{}) {
	if expected != actual {
		helper.T.Errorf("expected %v and actual %v", expected, actual)
	}

}
