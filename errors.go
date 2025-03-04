package slack

import "github.com/okovelkov/slack/internal/errorsx"

// Errors returned by various methods.
const (
	ErrAlreadyDisconnected  = errorsx.String("Invalid call to Disconnect - Slack API is already disconnected")
	ErrRTMDisconnected      = errorsx.String("disconnect received while trying to connect")
	ErrParametersMissing    = errorsx.String("received empty parameters")
	ErrInvalidConfiguration = errorsx.String("invalid configuration")
	ErrMissingHeaders       = errorsx.String("missing headers")
	ErrExpiredTimestamp     = errorsx.String("timestamp is too old")
)

// internal errors
const (
	errPaginationComplete = errorsx.String("pagination complete")
)
