def read_input() -> list:
    fileContent = ""
    with open("./puzzleInput.txt") as file:
        fileContent = file.read()
    return fileContent.split("\n")

def main():
    totalCalorieList = []
    puzzleInput = read_input()

    currentTotal = 0
    for calories in puzzleInput:
        if not calories:
            totalCalorieList.append(currentTotal)
            currentTotal = 0
            continue
        calorieValue = int(calories)
        currentTotal = currentTotal + calorieValue

    topThreeTotal = 0
    totalCalorieList.sort()
    for total in totalCalorieList[-3:]:
        topThreeTotal = total + topThreeTotal
    
    print(topThreeTotal)


if __name__ == "__main__":
    main()





        