package game

type Pawn struct { }

func (p Pawn) GetSymbol() rune {
    return 'p'
}

type Bishop struct { }

func (b Bishop) GetSymbol() rune {
    return 'B'
}

type Knight struct { }

func (k Knight) GetSymbol() rune {
    return 'N'
}

type Rook struct { }

func (r Rook) GetSymbol() rune {
    return 'R'
}

type Queen struct { }

func (q Queen) GetSymbol() rune {
    return 'Q'
}

type King struct { }

func (k King) GetSymbol() rune {
    return 'K'
}

