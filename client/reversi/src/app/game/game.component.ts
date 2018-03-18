import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-game',
  templateUrl: './game.component.html',
  styleUrls: ['./game.component.css']
})
export class GameComponent implements OnInit {
  stone: { [key: number]: string; } = {};
  field: number[][] = new Array();
  turn: number;

  blackNum: number;
  whiteNum: number;

  constructor() { }

  ngOnInit() {
    this.init();
  }

  onClickCell(x: number, y: number) {
    if (this.turn === 0) { return; }
    if (this.putStone(x, y, true)) {
      this.field[y][x] = this.turn;
      this.turn *= -1;
    }
    // TODO 美しくない
    let b = 0, w = 0;
    this.field.forEach(row => {
      row.forEach(cell => {
        if (cell ===  1) { b++; }
        if (cell === -1) { w++; }
      });
    });
    this.blackNum = b;
    this.whiteNum = w;
    this.isFinish();
  }

  isFinish() {
    if ( this.isPass() && this.isPass() ) {
      this.turn = 0;
      return true;
    }
    return false;
  }

  isPass() {
    for ( let i = 0; i < 8; i++ ) {
      for ( let j = 0; j < 8; j++ ) {
        if ( this.putStone(i, j, false) ) {
          return false;
        }
      }
    }
    this.turn *= -1;
    return true;
  }

  putStone(x: number, y: number, turn: boolean) {
    // 置いてあるとこには置けない
    if ( this.field[y][x] !== 0 ) { return false; }
    let turnCell = new Array();
    for (let i = -1; i < 2; i++) {
      for (let j = -1; j < 2; j++) {
        let tx = x + j;
        let ty = y + i;
        let stockCell = new Array();
        while (true) {
          if ( ty < 0 || 7 < ty || tx < 0 || 7 < tx ) { break; }
          if ( this.field[ty][tx] === this.turn * -1 ) {
            stockCell.push([ty, tx]);
            tx += j;
            ty += i;
          } else if ( this.field[ty][tx] === this.turn ) {
            stockCell.forEach(cell => {
              turnCell.push(cell);
            });
            break;
          } else { // 0
            break;
          }
        }
      }
    }
    turnCell.forEach(cell => {
      if (turn) {
        this.field[cell[0]][cell[1]] *= -1;
      }
    });
    return (turnCell.length === 0) ? false : true;
  }

  init() {
    for ( let i = 0; i < 8; i++ ) {
      this.field[i] = new Array();
      for ( let j = 0; j < 8; j++ ) {
        this.field[i][j] = 0;
      }
    }
    this.field[3][3] = this.field[4][4] = -1;
    this.field[3][4] = this.field[4][3] = 1;

    // TODO 美しくない
    this.stone[-1] = '⚪';
    this.stone[ 0] = '';
    this.stone[ 1] = '⚫';

    this.blackNum = this.whiteNum = 2;

    this.turn = 1;
  }

}
