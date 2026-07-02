package content

// Post represents a CMS post independent of how it is stored or delivered.
type Post struct {
	Title string
	Slug  string
	Body  string
}
