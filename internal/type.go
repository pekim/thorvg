package internal

type Tvg_Canvas uintptr

type Tvg_Result int

const (
	TVG_RESULT_SUCCESS                Tvg_Result = 0    ///< The value returned in case of a correct request execution.
	TVG_RESULT_INVALID_ARGUMENT       Tvg_Result = iota ///< The value returned in the event of a problem with the arguments given to the API - e.g. empty paths or null pointers.
	TVG_RESULT_INSUFFICIENT_CONDITION                   ///< The value returned in case the request cannot be processed - e.g. asking for properties of an object, which does not exist.
	TVG_RESULT_FAILED_ALLOCATION                        ///< The value returned in case of unsuccessful memory allocation.
	TVG_RESULT_MEMORY_CORRUPTION                        ///< The value returned in the event of bad memory handling - e.g. failing in pointer releasing or casting
	TVG_RESULT_NOT_SUPPORTED                            ///< The value returned in case of choosing unsupported engine features(options).
	TVG_RESULT_UNKNOWN                Tvg_Result = 255  ///< The value returned in all other cases.
)

type Tvg_Engine_Option int

const (
	TVG_ENGINE_OPTION_NONE         Tvg_Engine_Option = 0      /**< No engine options are enabled. This may be used to explicitly disable all optional behaviors. */
	TVG_ENGINE_OPTION_DEFAULT      Tvg_Engine_Option = 1 << 0 /**< Uses the default rendering mode. */
	TVG_ENGINE_OPTION_SMART_RENDER Tvg_Engine_Option = 1 << 1 /**< Enables automatic partial (smart) rendering optimizations. */
)
