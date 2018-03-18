export class Room {
  room_id: string;
  player_names: string[];
  spactator_names: string[];
}

// type Room struct {
// 	RoomID         string   `bson:"room_id" validate:"required"`
// 	PlayerNames    []string `bson:"player_names" validate:"required"`
// 	SpectatorNames []string `bson:"spector_names"`
// }
