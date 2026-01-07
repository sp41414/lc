package jumpGame

/*
logic:
[2, 3, 1, 1, 4]
            x, target
we start at the end of the slice, and also set our target there.

[2, 3, 1, 1, 4]
          x  target
then we check if our index (in this case is 3) + the value, which is 1, is greater than (or equal to) the target.
in this case, 3 + 1 = 4.
the target is the index, not the value.
the target here is 4, so the comparison is true.
we set target to i.

[2, 3, 1, 1, 4]
       x  target
then again, 2 + 1 = 3, which is equal to our target index.

[2, 3, 1, 1, 4]
    x target

here: 1 + 3 = 4.
our target is 2
4 > 2 so the target is now x.
[2, 3, 1, 1, 4]
 x target

here: 0 + 2 = 2
the target is 1
2 > 1 so the target is now at index 0.
the loop is finished,
so we check if the target is 0 or not. if it is, then we can jump to the end.
if it is not, then we cannot jump to the end.


another example of FALSE

[3,2,1,0,4]
         x,t

here: 4 + 4 = 8
8 > 4, so the target stays there.
[3,2,1,0,4]
       x t
here: 3 + 0 = 3
3 is NOT bigger than the target, so the target stays there.

[3,2,1,0,4]
     x   t

here: 2 + 1 = 3
3 is NOT bigger than the target, so target stays there.

this will repeat until the end of the loop, and target stays at the
end of the list.
since target is != 0, the function returns false
*/

func CanJump(nums []int) bool {
	target := len(nums) - 1

	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= target {
			target = i
		}
	}

	return target == 0
}
