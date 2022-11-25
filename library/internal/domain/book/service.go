package book

import "context"

type Service struct {
	storage Storage
}

func (receiver *Service) Create(ctx context.Context, data Book) {
	_, err := receiver.storage.Create(ctx, data)
	if err != nil {
		return
	}
}

func (receiver *Service) GetAll(ctx context.Context) ([]Book, error) {
	books, err := receiver.storage.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func NewService(s Storage) *Service {
	return &Service{storage: s}
}
