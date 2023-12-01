/*
Package txscript implements the pyrin transaction script language.

This package provides data structures and functions to parse and execute
pyrin transaction scripts.

# Script Overview

Pyrin transaction scripts are written in a stack-base, FORTH-like language.

The pyrin script language consists of a number of opcodes which fall into
several categories such pushing and popping data to and from the stack,
performing basic and bitwise arithmetic, conditional branching, comparing
hashes, and checking cryptographic signatures. Scripts are processed from left
to right and intentionally do not provide loops.

Typical pyrin scripts at the time of this writing are of several standard
forms which consist of a spender providing a public key and a signature
which proves the spender owns the associated private key. This information
is used to prove the the spender is authorized to perform the transaction.

One benefit of using a scripting language is added flexibility in specifying
what conditions must be met in order to spend pyrin.

# Errors

Errors returned by this package are of type txscript.Error. This allows the
caller to programmatically determine the specific error by examining the
ErrorCode field of the type asserted txscript.Error while still providing rich
error messages with contextual information. A convenience function named
IsErrorCode is also provided to allow callers to easily check for a specific
error code. See ErrorCode in the package documentation for a full list.
*/
package txscript
