package service

import (
	"github.com/Ivlay/upstore"
	htmlParser "github.com/Ivlay/upstore/pkg/parser"
	"github.com/Ivlay/upstore/pkg/repository"
	"go.uber.org/zap"
)

type ProductService struct {
	repo   repository.Product
	parser *htmlParser.HtmlParser
	logger *zap.Logger
}

func NewProductService(repo repository.Product, parser *htmlParser.HtmlParser, logger *zap.Logger) *ProductService {
	return &ProductService{
		repo:   repo,
		parser: parser,
		logger: logger,
	}
}

func (p *ProductService) Prepare() {
	count, err := p.repo.Count()
	if err != nil {
		p.logger.Error("Error while count products", zap.Error(err))
	}

	if count == 0 {
		product := p.parser.PrepareProduct()
		p.logger.Info("Prepare product", zap.Any("product", product))
		err := p.repo.Prepare(product)

		if err != nil {
			p.logger.Error("Error while prepare product", zap.Error(err))
		}
	}
}

func (p *ProductService) Update() (int, error) {
	product := p.parser.PrepareProduct()

	newProduct := upstore.Product{
		Title:   product.Title,
		Price:   product.Price,
		PriceId: product.PriceId,
	}

	return p.repo.Update(newProduct)
}

func (p *ProductService) Get() (upstore.Product, error) {
	return p.repo.Get()
}
