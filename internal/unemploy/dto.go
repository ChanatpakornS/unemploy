package unemploy

type UnemployedRequestParams struct {
	Start string `query:"start" validate:"required,datetime=2006-01-02"`
}

type WallpaperRequestParams struct {
	Start  string `query:"start" validate:"required,datetime=2006-01-02"`
	Width  int    `query:"width" validate:"required,min=800,max=7680"`
	Height int    `query:"height" validate:"required,min=600,max=4320"`
}
