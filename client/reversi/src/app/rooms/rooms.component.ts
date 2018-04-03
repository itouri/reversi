import { Component, OnInit } from '@angular/core';
import { RoomService } from '../room.service';
import { Room } from '../room';
import { Router } from '@angular/router';

@Component({
  selector: 'app-rooms',
  templateUrl: './rooms.component.html',
  styleUrls: ['./rooms.component.css']
})
export class RoomsComponent implements OnInit {
  rooms: Room[];
  player_name: string;

  constructor(
    private roomService: RoomService,
    private router: Router
  ) { }

  ngOnInit() {
    this.getRooms();
  }

  getRooms(): void {
    this.roomService.getRooms()
    // .subscribe(function(rooms) { this.rooms = rooms; console.log(this.rooms); } );
    .subscribe(rooms => this.rooms = rooms);
  }

  onClickCreate(): void {
    if (!this.player_name) { return; }
    this.roomService.createRoom(this.player_name)
    .subscribe((resJSON) => {
      console.log('resJSON', resJSON);
      this.router.navigateByUrl(`/game/${resJSON['RoomID']}/${resJSON['PlayerID']}/${this.player_name}`);
    });
  }

  onClickEnter(room_id: string): void {
    if (!this.player_name) { return; }
    this.roomService.enterRoom(room_id, this.player_name)
    .subscribe((player_id) => {
      this.router.navigateByUrl(`/game/${room_id}/${player_id}/${this.player_name}`);
    }); // TODO もっといいリダイレクトの方法
  }
}
