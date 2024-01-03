import requests
import getpass

def login_or_signup():
    userId = ""
    while True:
        print("Select an option:")
        print("1. Login")
        print("2. Signup")
        choice = input("Enter your choice (1 or 2): ")

        if choice == "1":
            username = input("Enter username: ")
            password = getpass.getpass(prompt="Enter password: ")
            response = requests.post("http://localhost:8080/user/login", json={"Username": username, "Password": password})
            
            # print(response.text)
            # print(response.status_code)

            if str(response.status_code)[:2] == '20':
                print("Login successful!")
                userId = response.json()['userID']
                return userId
            else:
                print("Invalid username or password. Please try again.")
        elif choice == "2":
            name = input("Enter your name: ")
            level = input("Enter your level: ")
            username = input("Enter username: ")
            password = getpass.getpass(prompt="Enter password: ")
            response = requests.post("http://localhost:8080/user/register", json={"Username": username, "Password": password, "Name": name, "Level": int(level)})
            # print(response.json())
            # print(response.status_code)
            if str(response.status_code)[:2] == '20':
                userId = response.json()['userID']
                print("Signup successful! Please login now.")
                return userId
            else:
                print("Error in signup. Please try again.")

def register_new_room(userId):
    room_name = input("Enter room name: ")
    capacity = int(input("Enter capacity: "))
    response = requests.post("http://localhost:8080/rooms/create", json={ "Capacity": int(capacity), "RoomName": int(room_name)})
    # print(response.json())
    # print(response.status_code)
    if str(response.status_code)[:2] == '20':
        print("Room registered successfully!")
    else:
        print("Error registering room.")

def book_room(userId):
    capacity = int(input("Enter required capacity: "))
    duration = int(input("Enter required duration (in minutes): "))

    response = requests.post("http://localhost:8080/rooms/book", json={"requiredCapacity": int(capacity), "userID": int(userId), "duration": int(duration)})
    # print(response.json())
    # print(response.status_code)
    if str(response.status_code)[:2] == '20':
        booked_room = response.json()['roomName']
        print(f"Room {booked_room} booked successfully!")
    else:
        print("No rooms available!")


def main():
    userId = login_or_signup()

    while True:
        print("\nOptions:")
        print("1. Register new room")
        print("2. Book a room")
        option = input("Enter your choice (1 or 2): ")

        if option == "1":
            register_new_room(userId)
            input("Press Enter to continue...")
        elif option == "2":
            book_room(userId)
            input("Press Enter to continue...")
        else:
            print("Invalid option. Please select again.")
        

if __name__ == "__main__":
    main()
