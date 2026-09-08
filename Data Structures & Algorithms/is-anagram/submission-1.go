func isAnagram(s string, t string) bool {
	
	
	
	
	
	
	


	if len(s) != len(t) {

		return false
	}
	lets := make([]int,26)
	for _, each := range s {
		lets[(each % 26)] += 1
	}

	for _, each := range t {
		lets[(each % 26)] -= 1
		if lets[(each % 26)] < 0 {
			return false
		}
	}

	return true


}
