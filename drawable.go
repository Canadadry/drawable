package drawable

type Drawable interface {
	Prepare() error
	Draw()
}
