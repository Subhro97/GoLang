package iomanager

type IOManager interface {
	ReadLines() ([]float64, error)
	WriteJSON(data interface{}) error
}
