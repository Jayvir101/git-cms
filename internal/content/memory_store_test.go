package content

import (
	"context"
	"fmt"
)

type memoryStore struct {
	posts map[string]Post
}

func newMemoryStore() *memoryStore {
	return &memoryStore{
		posts: make(map[string]Post),
	}
}

func (m *memoryStore) Create(ctx context.Context, post *Post) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	m.posts[post.Slug] = *post
	return nil
}

func (m *memoryStore) GetBySlug(ctx context.Context, slug string) (*Post, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	post, ok := m.posts[slug]
	if !ok {
		return nil, fmt.Errorf("post %q not found", slug)
	}

	return &post, nil
}

func (m *memoryStore) List(ctx context.Context) ([]Post, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	posts := make([]Post, 0, len(m.posts))
	for _, post := range m.posts {
		posts = append(posts, post)
	}

	return posts, nil
}

func (m *memoryStore) Update(ctx context.Context, post *Post) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	m.posts[post.Slug] = *post
	return nil
}

func (m *memoryStore) Delete(ctx context.Context, slug string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	delete(m.posts, slug)
	return nil
}

var _ ContentStore = (*memoryStore)(nil)
