package main

// If we want to treat simple object and complex object "Uniformly"
// First we make a interface!
type NeuronInterface interface {
	Iter() []*Neuron
}

// Neuron is a single structure (simple object)
type Neuron struct {
	In, Out []*Neuron
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

// ConnectTo function is used to connect several neuron
func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	n.In = append(other.In, n)
}

// NeuronLayer is a bunch of Neuron interconnected to each other Neuron (complex object)
type NeuronLayer struct {
	Neurons []Neuron
}

// Connect function is used to connect both of Neuron and NeuronLayer "Uniformly"
// This is implementation of Composite Design Pattern
func Connect(left, right NeuronInterface) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}

/*
	Why this is works?
	-> Because both of Neuron and NeuronLayer connected by the same interface
	-> In this case is "NeuronInterface" that have Iter() function.
	-> That's why this approach is working!
*/

func (n *NeuronLayer) Iter() []*Neuron {
	neurons := make([]*Neuron, 0)
	for indeks := range n.Neurons {
		neurons = append(neurons, &n.Neurons[indeks])
	}
	return neurons
}

// Constructor function of NeuronLayer
func NewNeuronLayer(count int) *NeuronLayer {
	return &NeuronLayer{make([]Neuron, count)}
}

func main() {
	neuron1, neuron2 := &Neuron{}, &Neuron{}
	layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

	// We want to connect the simple and complex object "Uniformly"
	// Using this Connect function!
	Connect(neuron1, neuron2)
	Connect(neuron1, layer1)
	Connect(layer2, neuron1)
	Connect(layer1, layer2)
}
