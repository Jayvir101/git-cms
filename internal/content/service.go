package content

import "context"

// ContentService contains application-level content rules and coordinates storage operations.
type ContentService struct {
	store ContentStore
}

// NewContentService creates a ContentService using the provided content store.
func NewContentService(store ContentStore) *ContentService {
	return &ContentService{
		store: store,
	}
}

// List returns all posts from the configured content store.
func (s *ContentService) List(ctx context.Context) ([]Post, error) {
	return s.store.List(ctx)
}
