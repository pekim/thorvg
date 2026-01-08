package internal

type enum int

/*
Result is an enumeration specifying the result from the APIs.

All ThorVG APIs could potentially return one of the values in the list.
Please note that some APIs may additionally specify the reasons that trigger their return values.
*/
type Result enum

const (
	RESULT_SUCCESS                Result = 0    ///< The value returned in case of a correct request execution.
	RESULT_INVALID_ARGUMENT       Result = iota ///< The value returned in the event of a problem with the arguments given to the API - e.g. empty paths or null pointers.
	RESULT_INSUFFICIENT_CONDITION               ///< The value returned in case the request cannot be processed - e.g. asking for properties of an object, which does not exist.
	RESULT_FAILED_ALLOCATION                    ///< The value returned in case of unsuccessful memory allocation.
	RESULT_MEMORY_CORRUPTION                    ///< The value returned in the event of bad memory handling - e.g. failing in pointer releasing or casting
	RESULT_NOT_SUPPORTED                        ///< The value returned in case of choosing unsupported engine features(options).
	RESULT_UNKNOWN                Result = 255  ///< The value returned in all other cases.
)

/*
EngineOption is an enumeration to specify rendering engine behavior.

The availability or behavior of @c ENGINE_OPTION_SMART_RENDER may vary depending on platform or backend support.
It attempts to optimize rendering performance by updating only the regions  of the canvas that have
changed between frames (partial redraw). This can be highly effective in scenarios  where most of the
canvas remains static and only small portions are updated—such as simple animations or GUI interactions.
However, in complex scenes where a large portion of the canvas changes frequently (e.g., full-screen animations
or heavy object movements), the overhead of tracking changes and managing update regions may outweigh the benefits,
resulting in decreased performance compared to the default rendering mode. Thus, it is recommended to benchmark
both modes in your specific use case to determine the optimal setting.

@since 1.0
*/
type EngineOption enum

const (
	ENGINE_OPTION_NONE         EngineOption = 0      /**< No engine options are enabled. This may be used to explicitly disable all optional behaviors. */
	ENGINE_OPTION_DEFAULT      EngineOption = 1 << 0 /**< Uses the default rendering mode. */
	ENGINE_OPTION_SMART_RENDER EngineOption = 1 << 1 /**< Enables automatic partial (smart) rendering optimizations. */
)
