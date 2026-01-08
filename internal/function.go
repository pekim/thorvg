package internal

func EngineInit(threads int) Result {
	return tvg_engine_init(threads)
}

func SwCanvasCreate(option EngineOption) Canvas {
	return tvg_swcanvas_create(option)
}
