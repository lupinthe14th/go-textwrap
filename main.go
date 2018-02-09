package main

import (
	"fmt"
	"strings"
)

func wrap(text string, lineWidth int) (wrapped string) {
	words := strings.Fields(text)
	if len(words) == 0 {
		return
	}
	wrapped = words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}
	return
}

var frog = `
In olden times when wishing still helped one, there lived a king
whose daughters were all beautiful, but the youngest was so beautiful
that the sun itself, which has seen so much, was astonished whenever
it shone in her face.  Close by the king's castle lay a great dark
forest, and under an old lime-tree in the forest was a well, and when
the day was very warm, the king's child went out into the forest and
sat down by the side of the cool fountain, and when she was bored she
took a golden ball, and threw it up on high and caught it, and this
ball was her favorite plaything.`

var japanese = `
祇園精舎の鐘の声、諸行無常の響きあり。沙羅双樹の花の色、盛者必衰の理をあらはす。おごれる人も久しからず、ただ春の夜の夢のごとし。たけき者も遂にはほろびぬ、ひとへに風の前の塵に同じ。 

遠くの異朝をとぶらへば、秦の趙高、漢の王莽、梁の朱忌、唐の禄山、これらは皆、旧主先皇の政にも従はず、楽しみを極め、諫めをも思ひ入れず、天下の乱れんことを悟らずして、民間の愁ふるところを知らざつしかば、久しからずして、亡じにし者どもなり。 

近く本朝をうかがふに、承平の将門、天慶の純友、康和の義親、平治の信頼、これらはおごれる心もたけきことも、皆とりどりにこそありしかども、間近くは六波羅の入道前太政大臣平朝臣清盛公と申しし人のありさま、伝え承るこそ、心も詞も及ばれね。 
`

func main() {
	fmt.Println("wrapped at 80:")
	fmt.Println(wrap(frog, 80))
	fmt.Println("wrapped at 72:")
	fmt.Println(wrap(frog, 72))
	fmt.Println("wrapped at 80:")
	fmt.Println(wrap(japanese, 80))
	fmt.Println("wrapped at 72:")
	fmt.Println(wrap(japanese, 72))
}
