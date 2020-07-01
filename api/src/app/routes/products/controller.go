package products

// Controller struct
type Controller struct {
	Repository Repository
}

// NewController func
func NewController(r Repository) *Controller {
	return &Controller{r}
}
