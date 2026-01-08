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
