package maps

/*
TODO: add ways to make maps persisting
TODO: remove expiring map, replacing it with a predicate based option; expiring just wraps sync, adds a time.Time field to the KV pairs, and then runs a condition to see if a KV pair is valid. This can, and should be generalized.
*/
