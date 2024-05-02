package usecase

import (
	"github.com/mrjonze/goexpert-clean-architecture/internal/entity"
	"github.com/mrjonze/goexpert-clean-architecture/pkg/events"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrdersListed    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface,
	OrdersListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrdersListed:    OrdersListed,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return make([]OrderOutputDTO, 0), err
	}

	var dtos []OrderOutputDTO
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		dtos = append(dtos, dto)
	}

	c.OrdersListed.SetPayload(dtos)

	c.EventDispatcher.Dispatch(c.OrdersListed)

	return dtos, nil
}
