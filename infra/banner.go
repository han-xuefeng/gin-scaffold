package infra

import "fmt"

const banner = `
 _______ _________ _               _______  _______  _______  _______  _______  _______  _        ______  
(  ____ \\__   __/( (    /|       (  ____ \(  ____ \(  ___  )(  ____ \(  ____ \(  ___  )( \      (  __  \ 
| (    \/   ) (   |  \  ( |       | (    \/| (    \/| (   ) || (    \/| (    \/| (   ) || (      | (  \  )
| |         | |   |   \ | | _____ | (_____ | |      | (___) || (__    | (__    | |   | || |      | |   ) |
| | ____    | |   | (\ \) |(_____)(_____  )| |      |  ___  ||  __)   |  __)   | |   | || |      | |   | |
| | \_  )   | |   | | \   |             ) || |      | (   ) || (      | (      | |   | || |      | |   ) |
| (___) |___) (___| )  \  |       /\____) || (____/\| )   ( || )      | )      | (___) || (____/\| (__/  )
(_______)\_______/|/    )_)       \_______)(_______/|/     \||/       |/       (_______)(_______/(______/ 

`

func init() {
	fmt.Println(banner)
}