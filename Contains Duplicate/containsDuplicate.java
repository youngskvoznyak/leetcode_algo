import java.util.HashSet;
import java.util.Set;

class contains_duplicate {
    public boolean containsDuplicate(int[] nums) {
        Set<Integer> numbers = new HashSet<Integer>();
        for (int num : nums) {
            if (numbers.contains(num)) return true;
            numbers.add(num);
        }

        return false;
    }
}