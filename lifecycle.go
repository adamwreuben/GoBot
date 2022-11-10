package GoBot

type GoBotLifecycle struct {
	ActiveStory     string
	ActiceStoryType string
	ActiveFormIds   []string
	ActiveForm      GoBotForm
	ActiveCounter   int
}

func NewLifecycle() *GoBotLifecycle {
	return &GoBotLifecycle{}
}

func (goBotLifecycle *GoBotLifecycle) SetState(state string, storyType string) {
	goBotLifecycle.ActiveStory = state
	goBotLifecycle.ActiceStoryType = storyType
}

func (goBotLifecycle *GoBotLifecycle) GetState() (string, string) {
	if goBotLifecycle.ActiveStory != "" {
		return goBotLifecycle.ActiveStory, goBotLifecycle.ActiceStoryType
	} else {
		return "default", "default"
	}
}
