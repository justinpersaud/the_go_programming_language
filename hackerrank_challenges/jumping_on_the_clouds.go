/*
Emma is playing a new mobile game that starts with consecutively numbered clouds. Some of the clouds are thunderheads and others are cumulus. She can jump on any cumulus cloud having a number that is equal to the number of the current cloud plus 1 or 2. She must avoid the thunderheads. Determine the minimum number of jumps it will take Emma to jump from her starting postion to the last cloud. It is always possible to win the game.

For each game, Emma will get an array of clouds numbered 0 if they are safe or 1 if they must be avoided. For example, c = [0,1,0,0,0,1,0] indexed from 0...6. The number on each cloud is its index in the list so she must avoid the clouds at indexes 1 and 5. She could follow the following two paths: 0 -> 2 -> 4 -> 6  or 0 -> 2 -> 3 -> 4 -> 6. The first path takes 3 jumps while the second takes 4.

Function Description

Complete the jumpingOnClouds function in the editor below. It should return the minimum number of jumps required, as an integer.

jumpingOnClouds has the following parameter(s):

c: an array of binary integers

Sample Input 0

7
0 0 1 0 0 1 0

Sample Output 0

4

Sample Input 1

6
0 0 0 0 1 0
Sample Output 1

3
*/
package main

import "fmt"

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	var hops int32
	var lastcloud = len(c) - 1
	for hops, cloud := 0, 0; cloud < lastcloud; cloud++ {
		if cloud == lastcloud {
			break
		}
		if isCloudSafe(c[cloud+2]) {
			fmt.Printf("Hopping two clouds over to %d\n", cloud+2)
			cloud++
		} else {
			fmt.Printf("Hopping one cloud over to %d\n", cloud+1)
			continue
		}
		fmt.Printf("Hop count is %d\n", hops)
		hops++
	}
	return hops
}

func isCloudSafe(n int32) bool {
	if n == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Printf("%d", jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0}))
}
