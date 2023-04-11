import fs from 'fs'

const path = __dirname + '/inputs/day2.txt'

const getMatches = (path: string): string[] => {
    try {
        const input = fs.readFileSync(path, 'utf-8')
        return input.split('\n')
    } catch (e) {
        console.log(e)
        return []
    }
}

export const matches = getMatches(path)

//A X = ROCK (1)
//B Y = PAPER (2)
//C Z = SCISSORS (3)

//WIN = 6
//TIE = 3
//LOSS = 0

//A X - LOSS (3)
//A Y - TIE (4)
//A Z - WIN (8)
//B X - LOSS (1)
//B Y - TIE (5)
//B Z - WIN (9)
//C X - LOSS (2)
//C Y - TIE (6)
//C Z - WIN (7)

const matchups = (match: string): number => {
    let score: number = 0
    switch (match) {
        case 'A X': 
            score = 3
            break
        case 'A Y': 
            score = 4
            break
        case 'A Z': 
            score = 8
            break
        case 'B X':
            score = 1
            break
        case 'B Y':
            score = 5
            break
        case 'B Z': 
            score = 9
            break
        case 'C X': 
            score = 2
            break
        case 'C Y':
            score = 6
            break
        case 'C Z':
            score = 7
            break
        default:
            console.log('invalid match')
            break
    }
    return score
}

export const scores: number[] = matches.map((match) => {
    return matchups(match)
})

export const total: number = scores.reduce((acc, curr) => acc + curr)

