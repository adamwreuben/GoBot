package GoBot

type GoBotLifecycle struct {
	ActiveStory         string
	ActiceStoryType     string
	ActiceNextStoryType string
	ActiveChoiceValue   string
	ActiveChoice        GoBotChoice
	NextStory           string
	ActiveFormIds       []string
	ActiveForm          GoBotForm
	ActiveInput         GoBotInput
	ActiveInputValue    string
	ActiveCounter       int
	SavedResults        map[string]interface{}
}

func NewLifecycle() *GoBotLifecycle {
	return &GoBotLifecycle{
		SavedResults: make(map[string]interface{}),
	}
}

func (goBotLifecycle *GoBotLifecycle) SetState(state string, storyType string) {
	goBotLifecycle.ActiveStory = state
	goBotLifecycle.ActiceStoryType = storyType
}

func (goBotLifecycle *GoBotLifecycle) SetNextStory(nextStory string, storyType string) {
	goBotLifecycle.NextStory = nextStory
	goBotLifecycle.ActiceNextStoryType = storyType
}

func (goBotLifecycle *GoBotLifecycle) GetNextStory() (string, string) {
	if goBotLifecycle.NextStory != "" {
		return goBotLifecycle.NextStory, goBotLifecycle.ActiceNextStoryType
	} else {
		return "", ""
	}
}

func (goBotLifecycle *GoBotLifecycle) GetState() (string, string) {
	if goBotLifecycle.ActiveStory != "" {
		return goBotLifecycle.ActiveStory, goBotLifecycle.ActiceStoryType
	} else {
		return "", ""
	}
}
