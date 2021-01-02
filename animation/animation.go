package animation

// Animation manages data to do animation
type Animation struct {
	start, end, index int
	timeAnim          int64
	timeLocal         int64
}

// GetIndex returns current index in list animation frames
func (anim Animation) GetIndex() int {
	return anim.index
}

// New initializes new Animation
func New(start, end, index int, timeAnim int64) *Animation {
	anim := &Animation{}
	return anim.new(start, end, index, timeAnim)
}

func (anim *Animation) new(start, end, index int, timeAnim int64) *Animation {
	anim.start = start
	anim.end = end
	anim.index = index
	anim.timeAnim = timeAnim
	return anim
}

// Update updates data in Animation when time flying
func (anim *Animation) Update(ellapseTime int64) {
	anim.timeLocal += ellapseTime
	if anim.timeLocal >= anim.timeAnim {
		anim.timeLocal = 0
		anim.next()
	}
}

func (anim *Animation) next() {
	anim.index++
	if anim.index > anim.end {
		anim.index = anim.start
	}
}
