/*
供一个下面这样的列表：
'a' 'b' 'a' 'a' 'a' 'c' 'd' 'e' 'f' 'g'
它将打印出没有后续重复的项目：
'a' 'b' 'a' 'c' 'd' 'e' 'f'
下面列出的 7.8 是 Perl 实现的算法。

Listing 7.8. uniq(1) 的 Perl 实现

#!/usr/bin/perl
my @a = qw/a b a a a c d e f g/;
print my $first = shift@a;
foreach (@a) {
	if ($first ne $_) { print ; $first = $_ ;}
}
*/
package main

import (
	"fmt"
)

func main() {
	list := []string{"a", "b", "a", "a", "c", "d", "e", "f"}
	first := list[0]

	fmt.Printf("%s ", first)
	for _, v := range list[1:] {
		if first != v {
			fmt.Printf("%s ", v)
			first = v
		}
	}
}
