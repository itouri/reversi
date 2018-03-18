import { Component, OnInit } from '@angular/core';
import { RoomService } from '../room.service';
import { Room } from '../room';

@Component({
  selector: 'app-rooms',
  templateUrl: './rooms.component.html',
  styleUrls: ['./rooms.component.css']
})
export class RoomsComponent implements OnInit {
  rooms: Room[];
  player_name: string;

  constructor(private roomService: RoomService) { }

  ngOnInit() {
    this.getRooms();
    this.player_name = 'angular';
  }

  getRooms(): void {
    this.roomService.getRooms()
    // .subscribe(function(rooms) { this.rooms = rooms; console.log(this.rooms); } );
    .subscribe(rooms => this.rooms = rooms);
  }

  onClick(room_id: string): void {
    this.roomService.enterRoom(room_id, this.player_name);
  }
}
