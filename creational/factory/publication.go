package factory


type iPublication interface {
	setName(name string)
	setPages(pages int)
	setPublished(publisher string)
	getName() string
	getPages() int
	getPublisher() string
}

type publication struct {
	name string
	pages int
	publisher string
}

func (p *publication) setName(name string) {
	p.name = name
}

func (p *publication) setPages(pages int) {
	p.pages = pages
}

func (p * publication) setPublished(published string) {
	p.publisher = published
}

func (p *publication) getName() string {
	return p.name
}

func (p *publication) getPages() int {
	return p.pages
}

func (p *publication) getPublisher() string {
	return p.publisher
}
