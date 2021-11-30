package managers

import (
	"github.com/thangtmc73/flappybird/objects"
)

type ScreenManager struct {
	Stack []objects.ScreenHandler
}

var instance *ScreenManager

func ScreenManagerInstance() *ScreenManager {
	if instance == nil {
		instance = newScreenManager()
	}
	return instance
}

func newScreenManager() *ScreenManager {
	manager := &ScreenManager{}
	return manager
}

func (s *ScreenManager) Back() (objects.ScreenHandler, int) {
	stackLen := len(s.Stack)
	if stackLen == 0 {
		return nil, -1
	}
	index := stackLen - 1
	return s.Stack[index], index
}

func (s *ScreenManager) IsEmpty() bool {
	return len(s.Stack) == 0
}

func (s *ScreenManager) AddScreen(screen objects.ScreenHandler) {
	s.Stack = append(s.Stack, screen)
	newScreen, _ := s.Back()
	if newScreen == nil {
		return
	}
	if !newScreen.Initialized() {
		newScreen.Init()
	}
}

func (s *ScreenManager) RemoveScreen() bool {
	if s.IsEmpty() {
		return false
	}

	currentScreen, index := s.Back()
	currentScreen.Destroy()
	s.Stack = s.Stack[:index]
	return true
}
