def read_input() -> list:
    fileContent = ""
    with open("./puzzleInput.txt") as file:
        fileContent = file.read()
    
    return fileContent.split("\n")

def main():
    puzzleInput = read_input()

    scoreKey = {
        "win": 6,
        "draw": 3,
        "loose": 0
    }

    moveKey = {
        "rock": 1,
        "paper": 2,
        "scissors": 3
    }

    



if __name__ == "__main__":
    main()

