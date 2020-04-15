package dicionario

//ErrorConst -
const (
	ErrNaoEncontrado      = ErrDicionario("Não foi possível encontrar a palavra procurada")
	ErrPalavraExistente   = ErrDicionario("Palavra já existente, não é possível adicioná-la")
	ErrPalavraInexistente = ErrDicionario("Palavra inexistente para ser atualizada")
)

//Dicionario -
type Dicionario map[string]string

//ErrDicionario -
type ErrDicionario string

//Busca -
func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, err := d[palavra]
	if !err {
		return "", ErrNaoEncontrado
	}

	return definicao, nil
}

//Add -
func (d Dicionario) Add(palavra, definicao string) error {
	_, err := d.Busca(palavra)

	switch err {
	case ErrNaoEncontrado:
		d[palavra] = definicao
	case nil:
		return ErrPalavraExistente
	default:
		return err

	}

	return nil
}

//Atualiza -
func (d Dicionario) Atualiza(palavra, novaDefinicao string) error {
	_, err := d.Busca(palavra)

	switch err {

	case ErrNaoEncontrado:
		return ErrPalavraInexistente
	case nil:
		d[palavra] = novaDefinicao
	default:
		return err
	}

	return nil
}

//Delete -
func (d Dicionario) Delete(palavra string) {
	delete(d, palavra)
}

func (e ErrDicionario) Error() string {
	return string(e)
}
