package item

type ItemService struct {
	*ItemRepository
}

func NewItemService(itemRepository *ItemRepository) {
	return &ItemService{
		ItemRepository: itemRepository,
	}
}
