package main

import (
	"bytes"
	"fmt"
)

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"
	x.Remove(144)
	x.Remove(9)
	fmt.Println(x.String()) // "{1 42}

	x.Clear()
	fmt.Println(x.String()) // "{}
	x.Add(3)
	fmt.Println(x.String()) // "{3}

	fmt.Println(y.Len()) //1

	copy := x.Copy()
	fmt.Println(copy.String()) // "{3}
	copy.Add(2)
	fmt.Println(copy.String()) // "{3,2}
	fmt.Println(x.String())    // "{3}
	x.AddAll(1, 2)
	fmt.Println(x.String()) // "{3,1,2}

	fmt.Println(x.Elem())

	var a, b IntSet
	a.Add(1)
	a.Add(2)
	a.Add(3)
	b.Add(2)
	b.Add(4)

	fmt.Println((a.IntersectWuith(&b)).String())    //2
	fmt.Println((a.DifferenceWuith(&b)).String())   //1 3
	fmt.Println(a.SymmetricDifference(&b).String()) //1 3 4
}

type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if !(word < len(s.words)) {
		return
	}
	s.words[word] ^= 1 << bit

}

func (s *IntSet) Clear() {
	s.words = make([]uint64, 0)

}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
	//s.words[word] = s.words[word] | (1<< bit);
}

func (s *IntSet) Copy() *IntSet {

	var copy IntSet
	var a int = len(s.words)
	copy.words = make([]uint64, a)
	for i, word := range s.words {
		copy.words[i] = word
	}
	return &copy
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}
func (s *IntSet) AddAll(values ...int) {
	for _, v := range values {
		word, bit := v/64, uint(v%64)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}

func (s *IntSet) Elem() []int {
	elements := make([]int, s.Len())
	var index int = 0
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				elements[index] = 64*i + j
				index++
			}
		}
	}
	return elements
}

func (s *IntSet) IntersectWuith(t *IntSet) *IntSet {
	var intersectSet IntSet
	elementsS := s.Elem()
	for _, num := range elementsS {
		if t.Has(num) {
			intersectSet.Add(num)
		}
	}
	return &intersectSet
}
func (s *IntSet) DifferenceWuith(t *IntSet) *IntSet {
	var diffSet IntSet
	elementsS := s.Elem()
	for _, num := range elementsS {
		if !t.Has(num) {
			diffSet.Add(num)
		}
	}
	return &diffSet
}
func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	var diffSet IntSet
	elementsS := s.Elem()
	elementsT := t.Elem()
	for _, num := range elementsS {
		if !t.Has(num) {
			diffSet.Add(num)
		}
	}
	for _, num := range elementsT {
		if !s.Has(num) {
			diffSet.Add(num)
		}
	}
	return &diffSet
}
