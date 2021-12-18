package datastructures

import "math"

type SnailFishNumber struct {
	Parent *SnailFishNumber
	Left   *SnailFishNumber
	Right  *SnailFishNumber

	Value int
}

func (s *SnailFishNumber) Add(sfn *SnailFishNumber) *SnailFishNumber {
	sum := &SnailFishNumber{
		Left:  s,
		Right: sfn,
		Value: 0,
	}
	sum.Left.Parent = sum
	sum.Right.Parent = sum
	madeChanges := false
	for {
		madeChanges = sum.Explode(sum, 0)
		if madeChanges {
			continue
		}
		madeChanges = sum.Split(sum)
		if !madeChanges {
			break
		}
	}
	return sum
}

func (s *SnailFishNumber) Explode(head *SnailFishNumber, depth int) bool {
	if depth > 4 {
		s.Parent.genExplode()
		return true
	}
	if s.Parent == nil && depth > 0 {
		return false
	}
	if s.Left == nil {
		return false
	}
	madeChanges := s.Left.Explode(head, depth+1)
	if madeChanges || s.Right == nil {
		return madeChanges
	}
	return s.Right.Explode(head, depth+1)
}

func (s *SnailFishNumber) Split(head *SnailFishNumber) bool {
	if s.Left == nil && s.Right == nil {
		if s.Value > 9 {
			s.genSplitTree()
			return true
		}
		return false
	}
	if s.Left == nil {
		return false
	}
	madeChanges := s.Left.Split(head)
	if s.Right == nil || madeChanges {
		return madeChanges
	}

	return s.Right.Split(head)
}

func (s *SnailFishNumber) genSplitTree() {
	s.Left = &SnailFishNumber{
		Left:   nil,
		Right:  nil,
		Parent: s,
		Value:  int(math.Floor(float64(s.Value) / 2)),
	}
	s.Right = &SnailFishNumber{
		Left:   nil,
		Right:  nil,
		Parent: s,
		Value:  int(math.Ceil(float64(s.Value) / 2)),
	}
	s.Value = 0
}

func (s *SnailFishNumber) genExplode() {
	left := s.findFirstLeftRegular()
	if left != nil {
		left.Value += s.Left.Value
	}
	right := s.findFirstRightRegular()
	if right != nil {
		right.Value += s.Right.Value
	}

	s.Left = nil
	s.Right = nil
	s.Value = 0
}

func (s SnailFishNumber) isRegular() bool {
	return s.Left == nil && s.Right == nil
}

func (s *SnailFishNumber) findFirstLeftRegular() *SnailFishNumber {
	prev := s
	for number := s.Parent; number != nil; {
		if number.isRegular() {
			return number
		}

		if number.Right == prev {
			prev = number
			number = number.Left
			continue
		}

		if number.Left == prev {
			if s.Parent == nil {
				break
			}
			prev = number
			number = number.Parent
			continue
		}

		number = number.Right
	}

	return nil
}

func (s *SnailFishNumber) findFirstRightRegular() *SnailFishNumber {
	prev := s
	for number := s.Parent; number != nil; {
		if number.isRegular() {
			return number
		}

		if number.Left == prev {
			prev = number
			number = number.Right
			continue
		}

		if number.Right == prev {
			if s.Parent == nil {
				break
			}
			prev = number
			number = number.Parent
			continue
		}

		number = number.Left
	}

	return nil
}

func (s *SnailFishNumber) Magnitude() int {
	sum := 0
	if s.Left != nil {
		sum += s.Left.Magnitude() * 3
	}
	if s.Right != nil {
		sum += s.Right.Magnitude() * 2
	}
	if s.isRegular() {
		return s.Value
	}

	return sum
}

func (s *SnailFishNumber) MakeCopy() *SnailFishNumber {
	clone := &SnailFishNumber{}
	if s.Left != nil {
		clone.Left = s.Left.MakeCopy()
		clone.Left.Parent = clone
	}
	if s.Right != nil {
		clone.Right = s.Right.MakeCopy()
		clone.Right.Parent = clone
	}
	clone.Value = s.Value

	return clone
}
