package thorvg

/*
EngineInit initializes the ThorVG engine.

Sets up the internal task scheduler and launches the specified
number of worker threads for parallel rendering operations.

	@param[in] threads The number of worker threads to create. A value of 0 disables worker threads.

	@return Tvg_Result indicating success or failure of initialization.

	@note Initialization is reference-counted. Each successful call increments the reference count.
				The worker thread count is determined by the first successful initialization and remains
				unchanged until the engine is fully terminated.

	@warning Each @ref tvg_engine_init() call must be paired with a @ref tvg_engine_term() call on the same thread.

	@see tvg_engine_term()
*/
func EngineInit(threads int) error {
	if err := initLibThorvg(); err != nil {
		return err
	}

	return tvg_engine_init(threads).error()
}

/*
EngineTerm terminates the ThorVG engine.

Releases internal resources and shuts down worker threads initialized by tvg_engine_init().

	@retval TVG_RESULT_INSUFFICIENT_CONDITION If the engine is not initialized.

	@note Initialization is reference-counted. Each call to @ref tvg_engine_term() decrements the reference count,
				and the engine is fully terminated only when the count reaches zero.

	@warning Each @ref tvg_engine_init() call must be paired with a @ref tvg_engine_term() call on the same thread.

	@see tvg_engine_init()
*/
func EngineTerm() error {
	return tvg_engine_term().error()
}

/*
Version retrieves the version of the TVG engine.

	@param[out] major A major version number.
	@param[out] minor A minor version number.
	@param[out] micro A micro version number.
	@param[out] version The version of the engine in the format major.minor.micro, or a @p nullptr in case of an internal error.

	@retval TVG_RESULT_SUCCESS.

	@since 0.15
*/
func Version() (int, int, int, string, string, error) {
	var major uint32
	var minor uint32
	var micro uint32
	var version *byte
	result := tvg_engine_version(&major, &minor, &micro, &version)

	return int(major), int(minor), int(micro), goString(version), libthorvgVersion, result.error()
}
