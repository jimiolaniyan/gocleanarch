package gocleanarch

import "github.com/jimiolaniyan/gocleanarch/gateways"

/**
Context is an application context to ensure that certain objects
are static during the lifetime of the application
*/

// CodecastRepo is a expected to be a concrete implementation of some form of persistence
var CodecastRepo gateways.CodecastGateway

var UserRepo gateways.UserGateway

var LicenseRepo gateways.LicenseGateway

var SessionKeeper *GateKeeper
