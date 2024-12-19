package api

type Gas = uint64
type GasMeter interface {
	GasConsumed() Gas
	GasLimit() Gas
	GasRemaining() Gas
	ConsumeGas(gas uint64, descriptor string)
}

type GasMeterNoOp struct{}

func (g *GasMeterNoOp) GasConsumed() Gas {
	return 0
}
func (g *GasMeterNoOp) GasLimit() Gas {
	return 0
}
func (g *GasMeterNoOp) GasRemaining() Gas {
	return 0
}
func (g *GasMeterNoOp) ConsumeGas(gas uint64, descriptor string) {}
