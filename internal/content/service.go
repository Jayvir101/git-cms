package content

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
