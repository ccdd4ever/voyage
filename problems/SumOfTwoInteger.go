//Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.
//see @https://leetcode.com/problems/sum-of-two-integers/
package problems

//we can calculate like adder or summer in digit circut
// A	B	C	S
// 0	0	0	0
// 1	0	0	1
// 0	1	0	1
// 1	1	1	0
//c means carry,we can get c using "&"(and) operator;s means sum,we can get s using "^"(xor)
//our final result could be calculate by 2*carry+s,and we can use recursion to replace the "+" operator

func add(a int32, b int32) int32 {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	return add(a&b<<1, a^b)
}

//func lengthOfLongestSubstring(s string) int {
//    length:=1
//    for i:=0;i<len(s);i++ {
//        for p:=i+1 ; p < len(s);p ++ {
//            if s[i]==s[p] {
//                if (p-i)>length {
//                   length=p-i
//                }
//                break
//            }
//        }
// }
// return length
//}
