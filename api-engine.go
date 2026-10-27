package thorvg

// Initializes the ThorVG engine.
//
// Sets up the internal task scheduler and launches the specified number of worker threads for parallel rendering operations.
//
//   - threads : The number of worker threads to create. A value of 0 disables worker threads.
//
// Return: Tvg_Result indicating success or failure of initialization.
//
// Note: Initialization is reference-counted. Each successful call increments the reference count.       The worker thread count is determined by the first successful initialization and remains       unchanged until the engine is fully terminated.
//
// Warning: Each [EngineInit] call must be paired with a [EngineTerm] call on the same thread.
//
// see [EngineTerm]
func EngineInit(threads int) error {
	if err := initLibThorvg(); err != nil {
		return err
	}

	return tvg_engine_init(threads).error()
}

// Terminates the ThorVG engine.
//
// Releases internal resources and shuts down worker threads initialized by tvg_engine_init().
//
// Returns [RESULT_INSUFFICIENT_CONDITION] : If the engine is not initialized.
//
// Note: Initialization is reference-counted. Each call to [EngineTerm] decrements the reference count,       and the engine is fully terminated only when the count reaches zero.
//
// Warning: Each [EngineInit] call must be paired with a [EngineTerm] call on the same thread.
//
// see [EngineInit]
func EngineTerm() error {
	return tvg_engine_term().error()
}

// Retrieves the version of the TVG engine.
//
//   - major : A major version number.
//   - minor : A minor version number.
//   - micro : A micro version number.
//   - version : The version of the engine in the format major.minor.micro, or a nullptr in case of an internal error.
//
// Returns [RESULT_SUCCESS] :
//
// Since: 0.15
func Version() (int, int, int, string, string, error) {
	var major uint32
	var minor uint32
	var micro uint32
	var version *byte
	result := tvg_engine_version(&major, &minor, &micro, &version)

	return int(major), int(minor), int(micro), goString(version), libthorvgVersion, result.error()
}
