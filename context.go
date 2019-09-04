package gocleanarch

/**
Context is an application context to ensure that certain objects
are static during the lifetime of the application
*/

// CodecastRepo is a expected to be a concrete implementation of some form of persistence
var CodecastRepo CodecastGateway

var UserRepo UserGateway

var LicenseRepo LicenseGateway
