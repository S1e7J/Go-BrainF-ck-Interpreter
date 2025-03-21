package loopcount

type LoopCount struct {
	loops []int
}

func New() LoopCount {
  return LoopCount{}
}

func (lc *LoopCount) Push(val int) {
  lc.loops = append(lc.loops, val)
}

func (lc *LoopCount) Peek() (int, bool) {
    if len(lc.loops) == 0 {
        return 0, false
    }
    return lc.loops[len(lc.loops)-1], true
}

func (lc *LoopCount) Pop() (int, bool) {
    if len(lc.loops) == 0 {
        return 0, false
    }
    lastIndex := len(lc.loops) - 1
    item := lc.loops[lastIndex]
    lc.loops = lc.loops[:lastIndex]
    return item, true
}

func (lc *LoopCount) Contains(val int) bool {
  for _, element := range lc.loops {
    if (element == val) {
      return true
    }
  }
  return false
}
