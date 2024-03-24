package game

var DefaultPieces = map[rune]map[rune]BoardElement{
    'a': {
        '1': Rook{},
        '2': Pawn{},
        '7': Pawn{},
        '8': Rook{},
    },
    'b': {
        '1': Knight{},
        '2': Pawn{},
        '7': Pawn{},
        '8': Knight{},
    },
    'c': {
        '1': Bishop{},
        '2': Pawn{},
        '7': Pawn{},
        '8': Bishop{},
    },
    'd': {
        '1': Queen{},
        '2': Pawn{},
        '7': Pawn{},
        '8': Queen{},
    },
    'e': {
        '1': King{},
        '2': Pawn{},
        '7': Pawn{},
        '8': King{},
    },
    'f': {
        '1': Bishop{},
        '2': Pawn{},
        '7': Pawn{},
        '8': Bishop{},
    },
    'g': {
        '1': Knight{},
        '2': Pawn{},
        '7': Pawn{},
        '8': Knight{},
    },
    'h': {
        '1': Rook{},
        '2': Pawn{},
        '7': Pawn{},
        '8': Rook{},
    },
}

