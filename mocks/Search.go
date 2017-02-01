package mocks

import "github.com/stretchr/testify/mock"

import "github.com/google/go-github/github"

type Search struct {
	mock.Mock
}

func (_m *Search) Issues(query string, opt *github.SearchOptions) (*github.IssuesSearchResult, *github.Response, error) {
	ret := _m.Called(query, opt)

	var r0 *github.IssuesSearchResult
	if rf, ok := ret.Get(0).(func(string, *github.SearchOptions) *github.IssuesSearchResult); ok {
		r0 = rf(query, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*github.IssuesSearchResult)
		}
	}

	var r1 *github.Response
	if rf, ok := ret.Get(1).(func(string, *github.SearchOptions) *github.Response); ok {
		r1 = rf(query, opt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, *github.SearchOptions) error); ok {
		r2 = rf(query, opt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
