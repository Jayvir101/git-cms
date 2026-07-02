package content

import "context"

// ContentStore defines the storage operations required for CMS posts.
type ContentStore interface {
	Create(ctx context.Context, post *Post) error
	GetBySlug(ctx context.Context, slug string) (*Post, error)
	List(ctx context.Context) ([]Post, error)
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, slug string) error
}
