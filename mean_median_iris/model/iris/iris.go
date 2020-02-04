package iris

type Iris struct {
	SepalLength,
	SepalWidth,
	PetalLength,
	PetalWidth float64
	Variety string
}

func New(sepaLenght, sepalWidth, petalLenght, petalWidth float64, variety string) *Iris {

	return &Iris{
		SepalLength: sepaLenght,
		SepalWidth:  sepalWidth,
		PetalLength: petalLenght,
		PetalWidth:  petalWidth,
		Variety:     variety,
	}

}
