package main

import (
    "fmt"
    "regexp"
    "io/ioutil"
    "strconv"
    "math"
    "container/list"
)

type MarbleGame struct {
    numMarbles int
    numPlayers int
    playerScores []int
    data *list.List
    current *list.Element
}

func NewMarbleGame(numMarbles, numPlayers int) *MarbleGame {
    return &MarbleGame {
        numMarbles,
        numPlayers,
        make([]int, numPlayers),
        list.New(),
        nil,
    }
}

func (game *MarbleGame) initialize() {
    game.current = game.data.PushFront(0)
}

func (game *MarbleGame) play() int {
    game.initialize()
    for i := 1; i <= game.numMarbles; i++ {
        if i % 23 == 0 {
            copyCurrent := game.current
            j := 7
            for j != 0 {
                copyCurrent = copyCurrent.Prev()
                if copyCurrent == nil {
                    copyCurrent = game.data.Back()
                }
                j--
            }
            game.current = copyCurrent.Next()
            elementScore := game.data.Remove(copyCurrent)
            game.playerScores[i % game.numPlayers] += i + elementScore.(int)
        } else {
            if game.current.Next() == nil {
                front := game.data.Front()
                game.current = game.data.InsertAfter(i, front)
            } else {
                game.current = game.data.InsertAfter(i, game.current.Next())
            }
        }
    }
    return game.findWinningScore()
}

func (game *MarbleGame) findWinningScore() int {
    winningScore := 0
    for i := 0; i < game.numPlayers; i++ {
        winningScore = int(math.Max(float64(winningScore), float64(game.playerScores[i])))
    }
    return winningScore
}


func main() {
    content, _ := ioutil.ReadFile("input.txt")
    input := string(content[:len(content)])
    var players int
    var marbles int
    r := regexp.MustCompile("[0-9]+")
    matches := r.FindAllString(input, -1)
    players, _ = strconv.Atoi(matches[0])
    marbles, _ = strconv.Atoi(matches[1])
    marbleGame := NewMarbleGame(marbles * 100, players)
    fmt.Println(marbleGame.play())
}
