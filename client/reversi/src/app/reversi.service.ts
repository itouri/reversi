import {Injectable} from '@angular/core';
import {Subject} from 'rxjs/Subject';
import {WebSocketService} from './websocket.service';
import {ReversiMessage} from './reversi.message';

@Injectable()
export class ReversiService {

  private messages: Subject<ReversiMessage>;

  private roomUrl(room_id: string, player_id: string, player_name: string): string {
    console.log('roomUrl', room_id);
    return `ws://localhost:23456/ws?room_id=${room_id}&player_name=${player_name}&player_id=${player_id}`;
  }

  constructor(private webSocketService: WebSocketService) {
  }

  connect(room_id: string, player_id: string, player_name: string): Subject<ReversiMessage> {
  return this.messages = <Subject<ReversiMessage>>this.webSocketService
    .connect(this.roomUrl(room_id, player_id, player_name))
    .map((response: MessageEvent): ReversiMessage => {
      // TODO serviceに書かない
      const data = JSON.parse(response.data) as ReversiMessage;
      return data;
    });
  }

  send(funcName: string, body: string): void {
    this.messages.next(this.createMessage(funcName, body));
  }

  private createMessage(funcName: string, body: string): ReversiMessage {
    return new ReversiMessage(funcName, body, false);
  }
}
