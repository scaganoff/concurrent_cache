Concurrent Cache
================

A simple exercise in building a concurrent cache using go and channels.

The cache has in input channel for requests and an output channel for responses
All activity is serialised through the input channel effectively acting as a mutex.
