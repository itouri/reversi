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

  constructor(private roomService: RoomService) { }

  ngOnInit() {
    this.getRooms();
  }

  getRooms(): void {
    this.roomService.getRooms()
    // .subscribe(function(rooms) { this.rooms = rooms; console.log(this.rooms); } );
    .subscribe(rooms => this.rooms = rooms);
  }
}
