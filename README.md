# reversi

## api
POST /api/v1/rooms
- create room
req val
owner_name :string

ret val
room_id :string

## scheme
- Player
name :string NonNull
player_id :string NonNUll

- Game
field :int[64]
state :int

- Room
player_name :string[2]
できないなら
balck_player_name :string
white_player_name :string


