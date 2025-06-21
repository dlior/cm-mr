import requests


def main():
    print("App 2: Main function executed!")
    joke = requests.get("https://api.chucknorris.io/jokes/random").json()
    print(f"joke: {joke['value']}")


if __name__ == "__main__":
    main()
