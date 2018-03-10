# reversi

## api
POST /api/v1/rooms
- create room
req val
owner_name :string

ret val
room_id :string

GET /api/v1/rooms/:id

## scheme
- Player
name :string NonNull
player_id :string NonNUll

- Game
field :int[64]
state :int

- Room
player_names :string[2]
    できないなら
    balck_player_name :string
    white_player_name :string
spectator_names :string[]


