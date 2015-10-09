package utilities

// THe scan process:
// 1. Start from root Map.
// 2. One level in one branch scanner. Check known sections to find out if the level in this branch is known.
// 2.1 If yes.
// 2.11 The level is a map: for each key in the map, look for the same key in the known section.
// 2.111 If there is a matched key in known section.
// 2.1111 If its value type is Map/Slice, search in known sections with the path and start from step 2.
// 2.1112 If its value type is Invalid, do the same as above.
// 2.1113 If its value type is something else, convert data by using the type got.
// 2.12 The level is a slice: look for the value type in the section.
// 2.121 If its value type is Map/Slice, search in known sections with the path and start from step 2.
// 2.122 If its value type is something else, convert data by using the type got.
// 2.2 If no.
// 2.21 Find out the value type by reflect.
// 2.22 Update the sections.
// 2.23 Start from step 2.

// The whole process is under condition that there are no same keys in the entire map. In case there are, please rename one of them before go through the process.

// Sections is the var to store the state of the known JSON-turned map.
var Sections = StringKeyMapTypeSections{}
