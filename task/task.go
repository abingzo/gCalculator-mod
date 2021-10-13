package task

func NewCommandTask() error {
	return nil
}

func NewCalculationTask(r string) string {
	return Calculate(ToPostfixExp(r))
}
