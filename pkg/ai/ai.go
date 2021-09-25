package ai

type AI interface {
	Placeholder()
	//Move
	//ChooseNextTarget
}

type AIHandler struct {
	ai map[int]AI
}

func NewAIHandler() *AIHandler {
	return &AIHandler{
		ai: make(map[int]AI),
	}
}

func (handler *AIHandler) AddAI(unitType int, AILogic AI) {
	handler.ai[unitType] = AILogic
}

func (handler *AIHandler) GetAI(unitType int) AI {
	return handler.ai[unitType]
}
