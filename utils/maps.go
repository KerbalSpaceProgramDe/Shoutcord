/*
 Shoutcord
 Bridging Modern Shoutbox and Discord

 Licensed under the Terms of the MIT License
 Copyright (c) 2018 KerbalSpaceProgram.de, Dorian Stoll
*/

package utils

// Taken from: https://www.dotnetperls.com/duplicates-go
func RemoveDuplicatesUnordered(elements []string) []string {
    encountered := map[string]bool{}

    // Create a map of all unique elements.
    for v := range elements {
        encountered[elements[v]] = true
    }

    // Place all keys from the map into a slice.
    result := []string{}
    for key, _ := range encountered {
        result = append(result, key)
    }
    return result
}

/*
 Exchanges the keys with the values in a map
*/
func SwapMap(m map[string]string) map[string]string {
    newMap := map[string]string{}
    for item := range m {
        newMap[m[item]] = item
    }
    return newMap
}
