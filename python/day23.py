from functools import lru_cache

rooms = (("C", "D"), ("A", "D"), ("B", "B"), ("C", "A"))
roomsUnfolded = (("C", "D", "D", "D"), ("A", "C", "B", "D"), ("B", "B", "A", "B"), ("C", "A", "C", "A"))

ROOM_MAP = (2, 4, 6, 8)
HALL_SPOTS = (0, 1, 3, 5, 7, 9, 10)
DESTINATION = {"A": 0, "B": 1, "C": 2, "D": 3}
COSTS = {"A": 1, "B": 10, "C": 100, "D": 1000}


def solve(lines):


    room_size = len(lines[0])

    hallway_start = tuple(None for _ in range(len(ROOM_MAP) + len(HALL_SPOTS)))

    @lru_cache(maxsize=None)
    def helper(hallway, rooms):
        if rooms == (("A",) * room_size, ("B",) * room_size, ("C",) * room_size, ("D",) * room_size):
            return 0

        best_cost = float('inf')
        for i, square in enumerate(hallway):  # Move from the hallway into a room.
            if square is None:
                continue
            dest = DESTINATION[square]
            can_move = True
            for roommate in rooms[dest]:
                if roommate is not None and roommate != square:
                    # Foreigner in room: can't move there.
                    can_move = False
                    break
            if not can_move:
                continue
            offset = 1 if ROOM_MAP[dest] > i else -1
            for j in range(i + offset, ROOM_MAP[dest] + offset, offset):
                if hallway[j] is not None:
                    can_move = False
                    break
            if not can_move:
                continue
            none_count = sum(elem is None for elem in rooms[dest])
            new_room = (None,) * (none_count - 1) + (square,) * (room_size - none_count + 1)
            steps = none_count + abs(i - ROOM_MAP[dest])
            cost = steps * COSTS[square]
            helper_result = helper(hallway[:i] + (None,) + hallway[i + 1:], rooms[:dest] + (new_room,)
                                   + rooms[dest + 1:])
            new_cost = cost + helper_result
            if new_cost < best_cost:
                best_cost = new_cost
        for i, room in enumerate(rooms):  # Move from a room into the hallway.
            wants_to_move = False
            for elem in room:
                if elem is not None and DESTINATION[elem] != i:
                    wants_to_move = True
            if not wants_to_move:
                continue
            none_count = sum(elem is None for elem in room)
            steps = none_count + 1
            square = room[none_count]
            for hall_destination in HALL_SPOTS:
                destination_steps = steps + abs(hall_destination - ROOM_MAP[i])
                destination_cost = destination_steps * COSTS[square]
                blocked = False
                for j in range(min(hall_destination, ROOM_MAP[i]), max(hall_destination, ROOM_MAP[i])+1):
                    if hallway[j] is not None:
                        blocked = True
                        break
                if blocked:
                    continue
                new_room = (None,) * (none_count + 1) + room[none_count + 1:]
                helper_result = helper(
                    hallway[:hall_destination] + (square,) + hallway[hall_destination + 1:],
                    rooms[:i] + (new_room,) + rooms[i + 1:])
                new_cost = destination_cost + helper_result
                if new_cost < best_cost:
                    best_cost = new_cost

        return best_cost

    cost = helper(hallway_start, lines)
    return cost


print(solve(rooms))
print(solve(roomsUnfolded))
