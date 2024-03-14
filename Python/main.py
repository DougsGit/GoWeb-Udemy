class GameObject:

    def __init__(self, name, appearance, feel, smell):
        self.name = name
        self.appearance = appearance
        self.feel = feel
        self.smell = smell

    def look(self):
        return f"You look at the {self.name}. {self.appearance} \n"

    def touch(self):
        return f"You touch the {self.name}. {self.feel} \n"

    def sniff(self):
        return f"You sniff the {self.name}. {self.smell} \n"


class Room:
    escape_code = 0
    game_objects = []
    def __init__(self, escape_code, game_objects):
        self.escape_code = escape_code
        self.game_objects

    def check_code(self, code):
        return self.escape_code == code

    def get_game_object_name(self):
        names = []
        for object in self.game_objects:
            names.append(object.name)
        return names


class Game:
    def __init__(self):
        self.attempts = 0
        objects = self.create_objects
        self.room = Room(1111,[])
    
    def create_objects(self):
        return[
            GameObject(
                "Sweater",
                "It's a blue sweater that had the number 12 stitched on it.",
                "Some has unstitched the second number, leaving only the 1.",
                "The sweater smells of laudry detergent."
            ),
            GameObject(
                "Chair",
                "It's a chair with only 3 legs.",
                "Someone has deliberately sawed off one of the legs.",
                "It smells like old wood."
            ),
            GameObject(
                
            )
        ]
