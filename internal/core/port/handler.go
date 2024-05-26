package port

type PriceEstimationHandler interface {
	Create(ctx interface{})
	Update(ctx interface{})
	Delete(ctx interface{})
	GetAll(ctx interface{})
	List(ctx interface{})
}

type LocationHandler interface {
	Create(ctx interface{})
	Get(ctx interface{})
	GetAll(ctx interface{})
	Delete(ctx interface{})
}

type UserHandler interface {
	Create(ctx interface{})
	Get(ctx interface{})
	GetAll(ctx interface{})
	Delete(ctx interface{})
}
