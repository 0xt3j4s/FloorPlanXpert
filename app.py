import requests
import getpass

class User:
    def __init__(self):
        self.user_id = None

    def login(self):
        while True:
            print("Select an option:")
            print("1. Login")
            print("2. Signup")
            choice = input("Enter your choice (1 or 2): ")

            if choice == "1":
                username = input("Enter username: ")
                password = getpass.getpass(prompt="Enter password: ")

                response = requests.post("http://localhost:8080/user/login", json={"Username": username, "Password": password})

                print(response)

                if str(response.status_code)[:2] == '20':
                    print("Login successful!")
                    # print(response.json())
                    self.user_id = response.json()['userID']
                    return self.user_id
                else:
                    print("Invalid username or password. Please try again.")
            elif choice == "2":
                name = input("Enter your name: ")
                level = input("Enter your level: ")
                username = input("Enter username: ")
                password = getpass.getpass(prompt="Enter password: ")
                response = requests.post("http://localhost:8080/user/register", json={"Username": username, "Password": password, "Name": name, "Level": int(level)})
                if str(response.status_code)[:2] == '20':
                    self.user_id = response.json()['userID']
                    print("Signup successful! Please login now.")
                    return self.user_id
                else:
                    print("Error in signup. Please try again.")

class Room:
    @staticmethod
    def register_new():
        room_name = input("Enter room name: ")
        capacity = int(input("Enter capacity: "))
        response = requests.post("http://localhost:8080/rooms/create", json={"Capacity": int(capacity), "RoomName": int(room_name)})
        if str(response.status_code)[:2] == '20':
            print("Room registered successfully!")
        else:
            print("Error registering room.")

    @staticmethod
    def book(user_id):
        capacity = int(input("Enter required capacity: "))
        start_time = input("Enter start time (in HH:MM format): ")
        duration = int(input("Enter required duration (in hrs): "))
        end_time = str(int(start_time[:2]) + duration) + start_time[2:]

        response = requests.post("http://localhost:8080/rooms/book", json={"requiredCapacity": int(capacity), "userID": int(user_id), "duration": int(duration)})
        if str(response.status_code)[:2] == '20':
            booked_room = response.json()['roomName']
            print(f"Room {booked_room} booked successfully!")
        else:
            print("No rooms available!")

def main():
    user = User()
    user_id = user.login()

    while True:
        print("\nOptions:")
        print("1. Register new room")
        print("2. Book a room")
        option = input("Enter your choice (1 or 2): ")

        if option == "1":
            Room.register_new()
            input("Press Enter to continue...")
        elif option == "2":
            Room.book(user_id)
            input("Press Enter to continue...")
        else:
            print("Invalid option. Please select again.")

if __name__ == "__main__":
    main()
