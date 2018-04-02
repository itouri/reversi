import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { catchError } from 'rxjs/operators';

import { Room, ReqRoom } from './room';
import { StaticInjector } from '@angular/core/src/di/injector';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable()
export class RoomService {
  private roomUrl = 'http://localhost:12345/api/v1/rooms';

  constructor( private http: HttpClient ) { }

  getRooms(): Observable<Room[]> {
    return this.http.get<Room[]>(this.roomUrl)
    .pipe(
      catchError(this.handleError('getRooms', []))
    );
  }

  createRoom(player_id: string, player_name: string): Observable<string> {
    const body = new(ReqRoom);
    body.player_name = player_name;
    if (player_id !== undefined) {
      body.player_id = player_id;
    }
    return this.http.post<string>(this.roomUrl, body).pipe(
      catchError(this.handleError<string>('createRoom'))
    );
  }

  enterRoom(room_id: string, player_id: string, player_name: string): Observable<any> {
    const body = new(ReqRoom);
    body.room_id = room_id;
    body.player_name = player_name;
    if (player_id !== undefined) {
      body.player_id = player_id;
    }
    return this.http.put(this.roomUrl, body).pipe(
      catchError(this.handleError<any>('enterRoom'))
    );
  }

  /**
   * 失敗したHttp操作を処理します。
   * アプリを持続させます。
   * @param operation - 失敗した操作の名前
   * @param result - observableな結果として返す任意の値
   */
  private handleError<T> (operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: リモート上のロギング基盤にエラーを送信する
      console.error(error); // かわりにconsoleに出力

      // 空の結果を返して、アプリを持続可能にする
      return of(result as T);
    };
  }
}
