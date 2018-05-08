import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, Params, Router} from '@angular/router';
import {ReversiService} from '../reversi.service';

import {Reversi} from './reversi';

@Component({
  selector: 'app-game',
  templateUrl: './game.component.html',
  styleUrls: ['./game.component.css']
})
export class GameComponent implements OnInit {
  // TODO このクラスがやってること多すぎ
  player_id: string;
  player_name: string;
  room_id: string;

  rs: Reversi;

  opponent: string;

  constructor(
    private reversiService: ReversiService,
    private route: ActivatedRoute,
    private router: Router,
  ) { }

  ngOnInit() {
    // TODO 絶対コンポーネント間で値を持ち回れるはず
    // チュートリアルではこうやってる
    this.room_id = this.route.snapshot.paramMap.get('room_id');
    this.player_id = this.route.snapshot.paramMap.get('player_id');
    this.player_name = this.route.snapshot.paramMap.get('player_name');
    console.log(this.room_id);
    console.log(this.player_id);
    this.rs = new Reversi();
    this.initWebsocket();
  }

  send(funcName: string, body: string) {
    this.reversiService.send(funcName, body);
  }

  rematch() {
    console.log('rematch');
    this.rs.init();
  }

  stoneColor(myColor: string) {
    console.log('mycolor:' + myColor);
    this.rs.myColor = Number(myColor);
  }

  getStone(putPos: string) {
    const pos = putPos.split(',').map(Number);
    this.rs.putStone(pos[0], pos[1], true);
    console.log('getStone: ', pos[0], pos[1]);
    // 呼んだ方もこの関数を実行してしまう
  }

  join(name: string) {
    if (this.opponent === undefined) {
      this.send('join', this.player_name);
    }
    this.opponent = name;
  }

  exit() {
    this.opponent = undefined;
  }

  initWebsocket() {
    // TODO connectの完了をハンドルするように書き換える
    setTimeout(() => this.send('join', this.player_name), 50 );
    this.reversiService.connect(this.room_id, this.player_id, this.player_name)
    .subscribe(msg => {
      // TODO もっといい方法ないのか?
      console.log(msg);
      switch (msg.funcName) {
        case 'rematch':
          this.rematch();
          break;
        case 'stoneColor':
          this.stoneColor(msg.body);
          break;
        case 'getStone':
          this.getStone(msg.body);
          break;
        case 'join':
          this.join(msg.body);
          break;
        case 'exit':
          this.exit();
          break;
      }
    });
  }

  onClickExit() {
    this.router.navigateByUrl(`/rooms`);
    // TODO ここで websocket の接続を切る
  }

  onClickRematch() {
    this.rs.init();
    this.send('rematch', '');
  }

  onClickCell(x: number, y: number) {
    if (this.rs.turn === 0 || this.rs.turn !== this.rs.myColor) { return; }
    if (this.rs.putStone(x, y, true)) {
      this.send('getStone', `${x}, ${y}`);
    }
    this.rs.isFinish();
  }
}
