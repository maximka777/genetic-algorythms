package errors

type GeneticError struct {
	Message string
}

func (err GeneticError) Error() string {
	return err.Message
}
