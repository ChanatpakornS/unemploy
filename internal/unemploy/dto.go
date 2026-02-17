package unemploy

type UnemployedRequestParams struct {
	Start string `query:"start" validate:"required,datetime=2006-01-02"`
}
