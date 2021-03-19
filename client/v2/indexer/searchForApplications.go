package indexer

import (
	"context"

	"github.com/algorand/go-algorand-sdk/client/v2/common"
	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
)

type SearchForApplicationsParams struct {

	// ApplicationId application ID
	ApplicationId uint64 `url:"application-id,omitempty"`

	// IncludeAll include all items including closed accounts, deleted applications,
	// destroyed assets, opted-out asset holdings, and closed-out application
	// localstates.
	IncludeAll bool `url:"include-all,omitempty"`

	// Limit maximum number of results to return.
	Limit uint64 `url:"limit,omitempty"`

	// Next the next page of results. Use the next token provided by the previous
	// results.
	Next string `url:"next,omitempty"`
}

type SearchForApplications struct {
	c *Client

	p SearchForApplicationsParams
}

// ApplicationId application ID
func (s *SearchForApplications) ApplicationId(ApplicationId uint64) *SearchForApplications {
	s.p.ApplicationId = ApplicationId
	return s
}

// IncludeAll include all items including closed accounts, deleted applications,
// destroyed assets, opted-out asset holdings, and closed-out application
// localstates.
func (s *SearchForApplications) IncludeAll(IncludeAll bool) *SearchForApplications {
	s.p.IncludeAll = IncludeAll
	return s
}

// Limit maximum number of results to return.
func (s *SearchForApplications) Limit(Limit uint64) *SearchForApplications {
	s.p.Limit = Limit
	return s
}

// Next the next page of results. Use the next token provided by the previous
// results.
func (s *SearchForApplications) Next(Next string) *SearchForApplications {
	s.p.Next = Next
	return s
}

func (s *SearchForApplications) Do(ctx context.Context, headers ...*common.Header) (response models.ApplicationsResponse, err error) {
	err = s.c.get(ctx, &response, "/v2/applications", s.p, headers)
	return
}