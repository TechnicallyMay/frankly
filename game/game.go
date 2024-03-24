package game

type Board struct {
    FileNames, RankNames []rune
    Files map[rune]map[rune]BoardElement 
    SelectedRank, SelectedFile rune
}

func (b Board) GetElementAtCoordinates(fileName rune, rankName rune) BoardElement {
    file, fileExists := b.Files[fileName] 
    boardElement, rankExists := file[rankName]

    if (!(fileExists && rankExists)) {
        panic("Tried to access a square that doesn't exist: " + string(fileName + rankName))
    }

    return boardElement
}

func (b Board) MovePiece(srcFile rune, srcRank rune, dstFile rune, dstRank rune) {
    srcElement := b.GetElementAtCoordinates(srcFile, srcRank)
    replaceWith := Square{}

    b.Files[dstFile][dstRank] = srcElement
    b.Files[srcFile][srcRank] = replaceWith
}

type BoardElement interface {
    GetSymbol() rune
}

type Square struct {}

func (s Square) GetSymbol() rune {
    return '~'
}

type Color int8

const (
    White Color = 0
    Black Color = 1
)

type GameState struct {
    PlayerColor Color
    Board Board
}


func GetDefaultBoard() Board {
    fileNames := [8]rune{ 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h' }
    rankNames := [8]rune{ '1', '2', '3', '4', '5', '6', '7', '8' }

    files := make(map[rune]map[rune]BoardElement, 8)

    for _, fileName := range fileNames {
        files[fileName] = make(map[rune]BoardElement, 8)
        for _, rankName := range rankNames {
            defaultPiece, foundPiece := DefaultPieces[fileName][rankName]

            if (foundPiece) {
                files[fileName][rankName] = defaultPiece
            } else {
                files[fileName][rankName] = Square{}
            }
        }
    }

    return Board{FileNames: fileNames[:], RankNames: rankNames[:], Files: files, SelectedFile: 'a', SelectedRank: '1'} 
}

