# examples

## error handling

To keep the examples relatively simple and uncluttered,
error handling is kept very simple.
Any error returned from a thorvg function results in a panic.

This is achieved by using the `SetErrorHandler` function,
providing an error handling function that simply panics.
This avoids having to check every error value returned from
functions.
