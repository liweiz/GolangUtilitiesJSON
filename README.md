Still WIP at early stage.

This is a utility package to parse JSON in Golang.

Many of quality APIs that serve JSON come with fixed schemas. However, we can not expect to get a complete schema with a single JSON since some of the data is probably not available for that specific request. What I can think of to get the complete schema is to get some JSONs from the API and find out the schema manually.

Or we might be able to make this progress automatic. Having the JSON schema with type info could help us to convert JSON to what we need faster. A map[string]reflect.Type is what we need for this job. The principal is to only include the part that we can confirm and for the unsure ones, just leave them our and explore it till enough info comes in.

Here is a more detailed process based on what we mentioned above. For the first time a JSON is received, it's available part of schema is kept as a map[string]reflect.Type. If there is any info missing to help make the decision, e.g. [] value and null, they are simply ignored and not merged into the type map. After the initial setup, each time when a JSON is received, 1) convert known values for keys directly since we already know their types from our type map; 2) leave out keys with incomplete info; 3) add unknown keys and types to type map; 4) convert values for unknown keys with known types; 5) set nil as value to known keys in map type without a meaningful value in this JSON.
