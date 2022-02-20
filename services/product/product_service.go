package product

type ProductService interface {
	Find() (Product, error)
}

type productService struct {
	repo ProductRepo
}

func NewProductService(repo ProductRepo) ProductService {
	return &productService{repo: repo}
}

func (p *productService) Find() (Product, error) {
	return p.repo.Find()
}

//asdasd
