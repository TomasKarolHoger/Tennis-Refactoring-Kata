package tenniskata

type tennisGame1 struct {
	player1Points int
	player2Points int
	player1Name   string
	player2Name   string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	game := &tennisGame1{
		player1Name: player1Name,
		player2Name: player2Name}

	return game
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == "player1" {
		game.player1Points += 1
	} else {
		game.player2Points += 1
	}
}

func (game *tennisGame1) GetScore() string {
	score := ""
	tempScore := 0
	if game.player1Points == game.player2Points {
		score = game.tieScore()
	} else if game.player1Points >= 4 || game.player2Points >= 4 {
		score = game.advantageOrWinScore()
	} else {
		for i := 1; i < 3; i++ {
			if i == 1 {
				tempScore = game.player1Points
			} else {
				score += "-"
				tempScore = game.player2Points
			}
			switch tempScore {
			case 0:
				score += "Love"
			case 1:
				score += "Fifteen"
			case 2:
				score += "Thirty"
			case 3:
				score += "Forty"
			}
		}
	}
	return score
}

func (game *tennisGame1) tieScore() string {
	switch game.player1Points {
	case 0:
		return "Love-All"
	case 1:
		return "Fifteen-All"
	case 2:
		return "Thirty-All"
	default:
		return "Deuce"
	}
}

func (game *tennisGame1) advantageOrWinScore() string {
	pointDifference := game.player1Points - game.player2Points
	if pointDifference == 1 {
		return "Advantage player1"
	} else if pointDifference == -1 {
		return "Advantage player2"
	} else if pointDifference >= 2 {
		return "Win for player1"
	}
	return "Win for player2"
}
