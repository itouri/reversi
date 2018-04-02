import { Player } from './player';

export class Room {
  room_id: string;
  players: Player[];
  spactators: string[];
}

export class ReqRoom {
  room_id?: string;
  player_id?: string;
  player_name?: string;
}

// type Room struct {
// 	RoomID         string   `bson:"room_id" validate:"required"`
// 	PlayerNames    []string `bson:"player_names" validate:"required"`
// 	SpectatorNames []string `bson:"spector_names"`
// }
