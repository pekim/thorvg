package internal

/*
EngineInit initializes the ThorVG engine.

ThorVG requires an active runtime environment to operate.
Internally, it utilizes a task scheduler to efficiently parallelize rendering operations.
You can specify the number of worker threads using the @p threads parameter.
During initialization, ThorVG will spawn the specified number of threads.

@param[in] threads The number of worker threads to create. A value of zero indicates that only the main thread will be used.

@note The initializer uses internal reference counting to track multiple calls.
The number of threads is fixed on the first call to tvg_engine_init() and cannot be changed in subsequent calls.

@see tvg_engine_term()
*/
func EngineInit(threads int) Result {
	return tvg_engine_init(threads)
}

/*
EngineTerm terminates the ThorVG engine.

Cleans up resources and stops any internal threads initialized by tvg_engine_init().

@retval TVG_RESULT_INSUFFICIENT_CONDITION Returned if there is nothing to terminate (e.g., tvg_engine_init() was not called).

@note The initializer maintains a reference count for safe repeated use. Only the final call to tvg_engine_term() will fully shut down the engine.

@see tvg_engine_init()
*/
func EngineTerm() Result {
	return tvg_engine_term()
}

/**
 * @brief Retrieves the version of the TVG engine.
 *
 * @param[out] major A major version number.
 * @param[out] minor A minor version number.
 * @param[out] micro A micro version number.
 * @param[out] version The version of the engine in the format major.minor.micro, or a @p nullptr in case of an internal error.
 *
 * @retval TVG_RESULT_SUCCESS.
 *
 * @since 0.15
 */
func Version() (Result, int, int, int, string) {
	var major uint32
	var minor uint32
	var micro uint32
	var version *byte
	result := tvg_engine_version(&major, &minor, &micro, &version)

	return result, int(major), int(minor), int(micro), GoString(version)
}
